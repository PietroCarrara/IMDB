package model

import (
	"github.com/apexskier/httpauth"
)

type User struct {
	id        int
	Nome      string
	senha     string
	Watchlist []Filme
	// Mapa que recebe o ID do filme e retorna a nota
	Avaliacoes  map[int]float32
	Comentarios []Comentario
	UserData    httpauth.UserData
}
