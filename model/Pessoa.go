package model

import (
	"time"
)

// Pessoa é alguém que participa em filmes
// (atores, diretores, etc...)
type Pessoa struct {
	Nascimento    time.Time
	Imagens       []Imagem `gorm:"many2many:pessoa_imagem"`
	Participacoes []Participante

	ID uint
}
