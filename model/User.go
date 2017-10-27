package model

import (
	"database/sql"
	"log"

	"github.com/apexskier/httpauth"
)

type User struct {
	Nome      string
	senha     string
	IsAdmin   bool
	Watchlist []*Filme
	// Mapa que recebe o ID do filme e retorna a nota
	Avaliacoes  map[int]float32
	Comentarios []Comentario
	UserData    httpauth.UserData

	id int
}

// Comentario representa um comentário do site
// feito em algum filme.
type Comentario struct {
	Alvo     *Filme
	Usuario  *User
	Conteudo string
}

func LoadAllUsers(db *sql.DB) []User {

	res, err := db.Query("SELECT id, nome, senha, is_admin FROM usuario")
	defer res.Close()
	if err != nil {
		log.Print("Erro ao executar Query em 'LoadAll()': " + err.Error())
		return nil
	}

	var users []User

	for res.Next() {

		var user User

		res.Scan(&user.id, &user.Nome, &user.senha, &user.IsAdmin)

		user.load(db)

		users = append(users, user)
	}

	return users
}

// LoadUserByName loads the user with the username u
func LoadUserByName(db *sql.DB, u string) *User {
	ps, err := db.Prepare("SELECT id, nome, senha FROM usuario WHERE nome = ?")
	defer ps.Close()
	if err != nil {
		log.Print("Erro ao preparar o statement em 'LoadFilme(db, " + u + ")': " + err.Error())
		return nil
	}

	res, err := ps.Query(u)
	defer res.Close()
	if err != nil {
		log.Print("Erro ao executar o PS em 'LoadFilme(db, " + u + ")': " + err.Error())
		return nil
	}

	var user User

	if res.Next() {
		res.Scan(&user.id, &user.Nome, &user.senha)

		user.load(db)

		return &user
	}

	// Se chegou até aqui é porque não
	// achou ninguém
	return nil
}

func LoadUserByID(db *sql.DB, u int) *User {
	ps, err := db.Prepare("SELECT id, nome, senha FROM usuario WHERE id = ?")
	defer ps.Close()
	if err != nil {
		log.Print("Erro ao preparar o statement em 'LoadFilme(db, )': " + err.Error())
		return nil
	}

	res, err := ps.Query(u)
	defer res.Close()
	if err != nil {
		log.Print("Erro ao executar o PS em 'LoadFilme(db, )': " + err.Error())
		return nil
	}

	var user User

	if res.Next() {
		res.Scan(&user.id, &user.Nome, &user.senha)

		user.load(db)

		return &user
	}

	// Se chegou até aqui é porque não
	// achou ninguém
	return nil
}

func (self *User) load(db *sql.DB) {
	self.loadWatchlist(db)
	self.loadComentarios(db)
}

func (self *User) loadWatchlist(db *sql.DB) {

	ps, err := db.Prepare("SELECT id_filme FROM usuario_filme_lista WHERE id_usuario = ?")
	defer ps.Close()
	if err != nil {
		log.Printf("Erro ao preparar o PS em 'loadWatchlist(db)': %s", err.Error())
	}

	res, err := ps.Query(self.id)
	defer res.Close()
	if err != nil {
		log.Printf("Erro ao executar o PS em 'loadWatchlist(db)': %s", err.Error())
	}

	for res.Next() {
		var id int

		res.Scan(&id)

		self.Watchlist = append(self.Watchlist, LoadFilme(db, id))
	}
}

func (self *User) loadComentarios(db *sql.DB) {
	ps, err := db.Prepare("SELECT id_filme, texto FROM usuario_filme_review WHERE id_usuario = ?")
	defer ps.Close()
	if err != nil {
		log.Printf("Erro ao preparar o PS em 'loadComentarios(db, %v)': %s", self, err.Error())
		return
	}

	res, err := ps.Query(self.id)
	defer res.Close()
	if err != nil {
		log.Printf("Erro ao executar o PS em 'oadComentarios(db, %v)': %s", self, err)
		return
	}

	for res.Next() {

		var (
			c  Comentario
			id int
		)

		res.Scan(&id, &c.Conteudo)

		c.Alvo = LoadFilme(db, id)
		c.Usuario = self

		self.Comentarios = append(self.Comentarios, c)
	}
}
