package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/cbroglie/mustache"
	"gitlab.com/rosso_pietro/IMDB/model"

	_ "github.com/go-sql-driver/mysql"
)

var (
	/*
		backend     httpauth.LeveldbAuthBackend
		aaa         httpauth.Authorizer
		port        = 8009
		backendfile = "auth.leveldb"
	*/

	db *sql.DB
)

func main() {

	var err error

	db, err = sql.Open("mysql", "IMDB_USER:3T3Dp1uaNXAxbxWv@/IMDB")
	defer db.Close()
	if err != nil {
		log.Print("Erro ao abrir a conexão com o banco em 'main()': " + err.Error())
	}

	o := model.LoadFilme(db, 2)

	sla, err := mustache.RenderFile("app/view/movie.html", o)
	if err != nil {
		// fmt.Printf("%s", err.Error())
	}

	fmt.Println(sla)

	/*
	   	var err error
	   	os.Mkdir(backendfile, 0755)
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

	   	// set up routers and route handlers
	   	r := mux.NewRouter()

	   	r.HandleFunc("/", getRoot).Methods("GET")

	   	http.Handle("/", r)
	   	fmt.Printf("Server running on port %d\n", port)
	   	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	   }

	   func getRoot(w http.ResponseWriter, r *http.Request) {
	   	//*/

}
