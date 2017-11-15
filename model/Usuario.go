package model

import (
	"github.com/apexskier/httpauth"
	"github.com/jinzhu/gorm"
)

// Usuario representa um usuário
// do sistema
type Usuario struct {
	Nome        string
	Senha       string
	IsAdmin     bool
	Watchlist   []Filme `gorm:"many2many:usuario_filme_watchlist"`
	Comentarios []Comentario
	Avaliacoes  []Avaliacao

	ID uint
}

// Comentario representa um texto
// escrito por um usuário em um filme
type Comentario struct {
	Filme    Filme
	Usuario  Usuario
	Conteudo string

	UsuarioID uint
	FilmeID   uint
	ID        uint
}

// Avaliacao representa uma nota
// que um usuário dá a um filme
type Avaliacao struct {
	Filme Filme
	Nota  float32

	UsuarioID uint
	FilmeID   uint
	ID        uint
}

// Load carrega o usuário
// a partir do banco de dados
func (user *Usuario) Load(db *gorm.DB) {
	db.Preload("Watchlist").Preload("Comentarios").Preload("Avaliacoes").Where(&Usuario{ID: user.ID}).First(user)

	for i := 0; i < len(user.Watchlist); i++ {
		user.Watchlist[i].Load(db)
	}
	for i := 0; i < len(user.Comentarios); i++ {
		user.Comentarios[i].Load(db)
	}
	for i := 0; i < len(user.Avaliacoes); i++ {
		user.Avaliacoes[i].Load(db)
	}
}

// Load carrega o comentário
// a partir do banco de dados
func (c *Comentario) Load(db *gorm.DB) {

	db.Where(&Usuario{ID: c.UsuarioID}).First(&c.Usuario)

}

// Load carrega a avaliacao
// a partir do banco de dados
func (a *Avaliacao) Load(db *gorm.DB) {
	db.Where(&Usuario{ID: a.FilmeID}).First(&a.Filme)
}

// ComentarioAdd adiciona um comentário
// ao usuário
func (user *Usuario) ComentarioAdd(c Comentario) {
	user.Comentarios = append(user.Comentarios, c)
}

// AvaliacaoAdd adiciona uma avaliação
// ao usuário
func (user *Usuario) AvaliacaoAdd(a Avaliacao) {
	user.Avaliacoes = append(user.Avaliacoes, a)
}

// UserData generates login data
// for the user struct
func (user *Usuario) UserData() httpauth.UserData {

	role := "user"
	if user.IsAdmin {
		role = "admin"
	}

	return httpauth.UserData{Username: user.Nome, Role: role}
}
