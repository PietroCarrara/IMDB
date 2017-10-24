package model

import (
	"time"
)

type Pessoa struct {
	Filmes     []Filme
	Fotos      []Imagem
	Nome       string
	Nascimento time.Time
	id         int
}
