package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/apexskier/httpauth"
	"github.com/cbroglie/mustache"
	"github.com/gorilla/mux"
	"gitlab.com/rosso_pietro/IMDB/model"

	_ "github.com/go-sql-driver/mysql"
)

var (
	backend     httpauth.LeveldbAuthBackend
	aaa         httpauth.Authorizer
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

	/*os.Mkdir(backendfile, 0755)
	defer os.Remove(backendfile)

	// create the backend
	backend, err = httpauth.NewLeveldbAuthBackend(backendfile)
	if err != nil {
		panic(err)
	}

	// create a default user
	username := "admin"
	defaultUser := httpauth.UserData{Username: username}
	err = backend.SaveUser(defaultUser)
	if err != nil {
		panic(err)
	}
	// Update user with a password and email address
	err = aaa.Update(nil, nil, username, "adminadmin", "admin@localhost.com")
	if err != nil {
		panic(err)
	}
	*/

	// set up routers and route handlers
	r := mux.NewRouter()

	r.HandleFunc("/", getRoot).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/test", test).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./app/view/")))

	http.Handle("/", r)
	fmt.Printf("Server running on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadFile("app/view/index.html")
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(bytes)
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

	w.Write([]byte(sla))
}

func test(w http.ResponseWriter, r *http.Request) {

	obj := model.LoadUserByName(db, "pietro")

	sla, err := mustache.RenderFile("app/view/test.html", obj)
	if err != nil {
		return
	}

	w.Write([]byte(sla))
}
