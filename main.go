package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/apexskier/httpauth"
	"github.com/cbroglie/mustache"
	"github.com/gorilla/mux"
	"gitlab.com/rosso_pietro/IMDB/model"

	_ "github.com/go-sql-driver/mysql"
)

var (
	backend     httpauth.LeveldbAuthBackend
	aaa         httpauth.Authorizer
	roles       map[string]httpauth.Role
	port        = 8009
	backendfile = "auth.leveldb"

	db *sql.DB
)

func main() {

	var err error

	db, err = sql.Open("mysql", "IMDB_USER:3T3Dp1uaNXAxbxWv@/IMDB")
	defer db.Close()
	if err != nil {
		log.Print("Erro ao abrir a conexão com o banco em 'main()': " + err.Error())
	}

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

	// create a default user
	username := "root"
	defaultUser := httpauth.UserData{Username: username, Role: "admin"}
	err = backend.SaveUser(defaultUser)
	if err != nil {
		panic(err)
	}

	// Update user with a password and email address
	aaa.Update(nil, nil, username, "adminadmin", "admin@localhost.com")

	for _, user := range model.LoadAllUsers(db) {
		err = backend.SaveUser(user.UserData)
		if err != nil {
			log.Fatal(err)
		}

		user.UpdateAuth(&aaa)
	}

	// set up routers and route handlers
	r := mux.NewRouter()

	r.HandleFunc("/", getRoot).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/user/{user}", getUser).Methods("GET")
	r.HandleFunc("/pessoa/{id}", getPessoa).Methods("GET")
	r.HandleFunc("/login", getLogin).Methods("GET")
	r.HandleFunc("/login", postLogin).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./app/view/")))

	http.Handle("/", r)
	fmt.Printf("Server running on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func getRoot(w http.ResponseWriter, r *http.Request) {

	obj := model.LoadAllFilmes(db)

	var user *model.User

	logged := map[string]bool{}

	data, err := aaa.CurrentUser(w, r)
	if err != nil {
		logged["logged"] = false
	} else {
		user = model.LoadUserByName(db, data.Username)
		logged["logged"] = true
	}

	str, err := mustache.RenderFile("./app/view/index.html", obj, user, logged)
	if err != nil {
		return
	}

	w.Header().Set("Content-type", "text/html")
	w.Write([]byte(str))
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return
	}

	filme := model.LoadFilme(db, id)

	sla, err := mustache.RenderFile("app/view/movie.html", filme)
	if err != nil {
		return
	}

	w.Header().Set("Content-type", "text/html")
	w.Write([]byte(sla))
}

func getUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	nome := vars["user"]

	user := model.LoadUserByName(db, nome)

	sla, err := mustache.RenderFile("app/view/user.html", user)
	if err != nil {
		return
	}

	w.Header().Set("Content-type", "text/html")
	w.Write([]byte(sla))
}

func getLogin(w http.ResponseWriter, r *http.Request) {

	sla, err := mustache.RenderFile("app/view/login.html", nil)
	if err != nil {
		return
	}

	w.Header().Set("Content-type", "text/html")
	w.Write([]byte(sla))
}

func getPessoa(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	i, _ := strconv.Atoi(vars["id"])

	p := model.LoadPessoaByID(db, i)

	fmt.Print(p.Filmes)

	page, err := mustache.RenderFile("app/view/pessoa.html", p)
	if err != nil {
		print(err)
		return
	}

	w.Header().Set("Content-type", "text/html")
	w.Write([]byte(page))
}

func getLoginWithErros(w http.ResponseWriter, r *http.Request, erros ...string) {

	sla, err := mustache.RenderFile("app/view/login.html", erros)
	if err != nil {
		return
	}

	w.Header().Set("Content-type", "text/html")
	w.Write([]byte(sla))
}

func postLogin(rw http.ResponseWriter, req *http.Request) {
	username := req.PostFormValue("username")
	password := req.PostFormValue("password")
	if err := aaa.Login(rw, req, username, password, "/"); err == nil || (err != nil && strings.Contains(err.Error(), "already authenticated")) {
		getLoginWithErros(rw, req, "Você já se autenticou!")
	} else if err != nil {
		fmt.Println(err)
		getLoginWithErros(rw, req, "Usuário/Senha incorretos!")
	}
}
