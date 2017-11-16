package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/cbroglie/mustache"
	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/apexskier/httpauth"
	"github.com/jinzhu/gorm"
	"gitlab.com/rosso_pietro/IMDB/model"
)

var (
	db *gorm.DB

	backend     httpauth.LeveldbAuthBackend
	aaa         httpauth.Authorizer
	roles       map[string]httpauth.Role
	backendfile = "auth.leveldb"
	port        = 8000
)

func main() {
	var err error

	db, err = gorm.Open("mysql", "IMDB:X8NJumITHkty0LnP@/IMDB_TABLE")
	// db = db.Debug()
	defer db.Close()
	if err != nil {
		log.Printf("Erro ao abrir a conexão com o banco em 'main()': %s", err.Error())
	}

	db.AutoMigrate(&model.Filme{})
	db.AutoMigrate(&model.Pessoa{})
	db.AutoMigrate(&model.Cargo{})
	db.AutoMigrate(&model.Tag{})
	db.AutoMigrate(&model.Imagem{})
	db.AutoMigrate(&model.Usuario{})
	db.AutoMigrate(&model.Comentario{})
	db.AutoMigrate(&model.Avaliacao{})

	setupAuth()

	r := setupRouter()

	http.ListenAndServe(":8000", r)
}

func setupAuth() {

	var err error

	os.Mkdir(backendfile, 0755)
	defer os.Remove(backendfile)

	// create the backend
	backend, err = httpauth.NewLeveldbAuthBackend(backendfile)
	if err != nil {
		panic(err)
	}

	// create some default roles
	roles = make(map[string]httpauth.Role)
	roles["user"] = 30
	roles["admin"] = 80
	aaa, err = httpauth.NewAuthorizer(backend, []byte("cookie-encryption-key"), "user", roles)
}

func setupRouter() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/movie/{id}", filme)
	r.HandleFunc("/movie/{id}/nota", getNota)
	r.HandleFunc("/movie/{id}/rate", avaliar)
	r.HandleFunc("/movie/{id}/comment", comentar)
	r.HandleFunc("/movie/{id}/tag-add", addTag)
	r.HandleFunc("/busca", busca)
	r.HandleFunc("/admin/insert/movie", insFilmePage).Methods("GET")
	r.HandleFunc("/admin/insert/movie", insFilme).Methods("POST")
	r.HandleFunc("/admin/insert/person", insPessoaPage).Methods("GET")
	r.HandleFunc("/admin/insert/person", insPessoa).Methods("POST")
	r.HandleFunc("/admin/toggle/{id}", toggleAdmin)
	r.HandleFunc("/user/{nome}", usuario)
	r.HandleFunc("/pessoa/{id}", pessoa)
	r.HandleFunc("/tags/{nome}", tag)
	r.HandleFunc("/login", loginPage)
	r.HandleFunc("/auth", login)
	r.HandleFunc("/register", register)
	r.HandleFunc("/logout", logout)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	return r
}

func home(w http.ResponseWriter, r *http.Request) {
	filmes := []model.Filme{}
	db.Find(&filmes, &model.Filme{})
	model.LoadFilmeSlice(filmes, db)

	var user *model.Usuario

	logged := map[string]bool{}

	user = currentUser(w, r)
	if user != nil {
		logged["logged"] = true
	}

	str, err := mustache.RenderFile("./templates/index.html", filmes, logged, user)
	if err != nil {
		log.Println(err.Error())
		return
	}

	w.Header().Set("Content-type", "text/html")
	w.Write([]byte(str))
}

func filme(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	options := map[string]interface{}{}

	tags := []model.Tag{}
	db.Find(&tags)
	options["alltags"] = tags

	id, err := strconv.ParseUint(vars["id"], 10, 0)
	if err != nil {
		return
	}

	var user *model.Usuario

	user = currentUser(w, r)
	if user != nil {
		options["logged"] = true
	}

	var filme model.Filme
	db.Where(model.Filme{ID: uint(id)}).First(&filme)
	filme.Load(db)

	sla, err := mustache.RenderFile("./templates/movie.html", filme, user, options)
	if err != nil {
		log.Print(err)
		return
	}

	w.Header().Set("Content-type", "text/html")
	w.Write([]byte(sla))
}

func usuario(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	var user *model.Usuario

	options := map[string]interface{}{}

	user = currentUser(w, r)
	if user != nil {
		user.Load(db)
		options["logged"] = true
	}

	options["user"] = user

	var target model.Usuario
	db.Where(&model.Usuario{Nome: vars["nome"]}).First(&target)
	target.Load(db)

	options["target"] = target

	sla, err := mustache.RenderFile("./templates/user.html", options)
	if err != nil {
		log.Print(err)
		return
	}

	w.Header().Set("Content-type", "text/html")
	w.Write([]byte(sla))
}

func busca(w http.ResponseWriter, r *http.Request) {

	q := r.PostFormValue("q")

	filmes := []model.Filme{}
	db.Find(&filmes)

	filmes = model.SearchFilmes(filmes, q)
	model.LoadFilmeSlice(filmes, db)

	var user *model.Usuario

	logged := map[string]bool{}

	user = currentUser(w, r)
	if user != nil {
		logged["logged"] = true
	}

	str, err := mustache.RenderFile("./templates/index.html", filmes, logged, user)
	if err != nil {
		log.Println(err.Error())
		return
	}

	w.Header().Set("Content-type", "text/html")
	w.Write([]byte(str))
}

func pessoa(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 0)

	pessoa := model.Pessoa{ID: uint(id)}
	db.First(&pessoa)
}

func tag(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	var tag model.Tag
	db.Where(model.Tag{Titulo: vars["nome"]}).First(&tag)
	tag.Load(db)

	model.LoadFilmeSlice(tag.Filmes, db)

	var user *model.Usuario

	logged := map[string]bool{}

	user = currentUser(w, r)
	if user != nil {
		logged["logged"] = true
	}

	str, err := mustache.RenderFile("./templates/index.html", tag.Filmes, logged, user)
	if err != nil {
		log.Println(err.Error())
		return
	}

	w.Header().Set("Content-type", "text/html")
	w.Write([]byte(str))
}

func insPessoaPage(w http.ResponseWriter, r *http.Request) {
	options := map[string]interface{}{}

	usr := currentUser(w, r)
	if usr == nil || !usr.IsAdmin {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	options["logged"] = true
	options["user"] = usr

	str, _ := mustache.RenderFile("./templates/person.html", options)

	w.Write([]byte(str))
}

func insFilmePage(w http.ResponseWriter, r *http.Request) {
	options := map[string]interface{}{}

	usr := currentUser(w, r)
	if usr == nil || !usr.IsAdmin {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	options["logged"] = true
	options["user"] = usr

	str, _ := mustache.RenderFile("./templates/movieInsert.html", options)

	w.Write([]byte(str))
}

func insPessoa(w http.ResponseWriter, r *http.Request) {
	nome := r.PostFormValue("nome")

	dia, _ := strconv.Atoi(r.PostFormValue("dia"))
	mes, _ := strconv.Atoi(r.PostFormValue("mes"))
	ano, _ := strconv.Atoi(r.PostFormValue("ano"))

	m := time.Month(mes)

	t := time.Date(ano, m, dia, 0, 0, 0, 0, time.Local)

	nasc := mysql.NullTime{Time: t, Valid: true}

	pessoa := model.Pessoa{Nome: nome, Nascimento: nasc}

	db.Save(&pessoa)

	pic, _, err := r.FormFile("pic")
	if err != nil {
		log.Println(err.Error())
	}
	foto := model.Imagem{}
	foto.Pessoas = append(foto.Pessoas, pessoa)

	db.Save(&foto)

	name := fmt.Sprintf("/uploads/upload%d", foto.ID)

	file, err := os.Create("./static" + name)
	if err != nil {
		log.Println(err.Error())
	}

	_, err = io.Copy(file, pic)

	foto.Caminho = name

	db.Save(&foto)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func insFilme(w http.ResponseWriter, r *http.Request) {

	titulo := r.PostFormValue("titulo")
	sinopse := r.PostFormValue("sinopse")

	filme := model.Filme{Titulo: titulo, Sinopse: sinopse}

	db.Save(&filme)

	pic, _, err := r.FormFile("pic")
	if err != nil {
		log.Println(err.Error())
	}
	foto := model.Imagem{FilmeID: filme.ID}

	db.Save(&foto)

	name := fmt.Sprintf("/uploads/upload%d", foto.ID)

	file, err := os.Create("./static" + name)
	if err != nil {
		log.Println(err.Error())
	}

	_, err = io.Copy(file, pic)

	foto.Caminho = name

	db.Save(&foto)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func addTag(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	idStr := vars["id"]
	id, _ := strconv.ParseUint(idStr, 10, 0)

	bts, _ := ioutil.ReadAll(r.Body)

	tagName := string(bts)

	var tag model.Tag
	db.Where(&model.Tag{Titulo: tagName}).First(&tag)

	var filme model.Filme
	db.Where(&model.Filme{ID: uint(id)}).First(&filme)
	filme.Load(db)

	if filme.TagAdd(tag) {
		db.Save(&filme)
		w.Write([]byte(tag.Titulo))
	}
}

func loginPage(w http.ResponseWriter, r *http.Request) {

	sla, err := mustache.RenderFile("./templates/login.html", nil)
	if err != nil {
		return
	}

	w.Header().Set("Content-type", "text/html")
	w.Write([]byte(sla))
}

func comentar(w http.ResponseWriter, r *http.Request) {

	user := currentUser(w, r)
	if user == nil {
		log.Println("Comentar necessita de login")
		return
	}

	bytes, _ := ioutil.ReadAll(r.Body)
	text := string(bytes)

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 0)
	if err != nil {
		log.Println(err.Error())
		return
	}

	var filme model.Filme
	db.Where(&model.Filme{ID: uint(id)}).First(&filme)

	c := model.Comentario{Filme: filme, Conteudo: text}

	user.ComentarioAdd(c)

	db.Save(&user)

	w.Write([]byte(c.Conteudo))
}

func getNota(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 0)
	if err != nil {
		log.Println(err.Error())
		return
	}

	var filme model.Filme
	db.Where(&model.Filme{ID: uint(id)}).First(&filme)
	filme.Load(db)

	fmt.Fprintf(w, "%g", filme.Nota())
}

func avaliar(w http.ResponseWriter, r *http.Request) {

	user := currentUser(w, r)
	if user == nil {
		log.Println("Comentar necessita de login")
		return
	}

	bts, _ := ioutil.ReadAll(r.Body)
	n := string(bts)
	nota, _ := strconv.ParseFloat(n, 32)

	if nota > 5 {
		nota = 5
	} else if nota < 1 {
		nota = 1
	}

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 0)
	if err != nil {
		log.Println(err.Error())
		return
	}

	aval := model.Avaliacao{FilmeID: uint(id), UsuarioID: user.ID}

	db.Where(&aval).FirstOrCreate(&aval)

	aval.Nota = float32(nota)

	db.Save(&aval)

	w.Write([]byte(n))
}

func toggleAdmin(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, _ := strconv.ParseUint(vars["id"], 10, 0)

	usr := model.Usuario{ID: uint(id)}

	db.First(&usr)

	usr.IsAdmin = !usr.IsAdmin

	db.Save(&usr)

	fmt.Fprintf(w, "%t", usr.IsAdmin)
}

func login(w http.ResponseWriter, r *http.Request) {

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	if usr := currentUser(w, r); usr != nil {
		log.Printf("O usuário %s está tentando se logar novamente como %s!", usr.Nome, username)
		// getLoginWithErros(w, r, "Usuário não existe!")
		return
	}

	var user model.Usuario
	err := db.Where(&model.Usuario{Nome: username}).First(&user).Error

	if err != nil {
		log.Printf("O usuário %s não existe!", username)
		// getLoginWithErros(w, r, "Usuário não existe!")
		return
	}
	if user.Senha != password {
		log.Printf("O usuário %s errou a senha!", username)
		// getLoginWithErros(w, r, "Senha inválida!")
		return
	}

	backend.SaveUser(user.UserData())
	aaa.Update(w, r, username, password, "NONE")

	err = aaa.Login(w, r, username, password, "/")
	if err != nil {
		log.Printf("Erro no login do usuário '%s'", username)
		log.Println(err)
		backend.DeleteUser(username)
		// getLoginWithErros(w, r, "Erro no login!")
	} else {
		log.Printf("%s se logou", username)
		http.Redirect(w, r, "/", http.StatusAccepted)
	}
}

func register(w http.ResponseWriter, r *http.Request) {

	usr := r.PostFormValue("username")
	pwd := r.PostFormValue("password")

	u := model.Usuario{Nome: usr, Senha: pwd}

	db.Save(&u)

	login(w, r)
}

func logout(w http.ResponseWriter, r *http.Request) {

	if err := aaa.Logout(w, r); err != nil {
		log.Println(err)
		// this shouldn't happen
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func currentUser(w http.ResponseWriter, r *http.Request) *model.Usuario {

	data, err := aaa.CurrentUser(w, r)
	if err == nil {
		var res model.Usuario
		db.Where(&model.Usuario{Nome: data.Username}).First(&res)
		res.Load(db)

		return &res
	}

	return nil
}
