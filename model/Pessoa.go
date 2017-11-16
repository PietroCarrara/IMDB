package model

import "github.com/go-sql-driver/mysql"

type Pessoa struct {
	Nome          string
	Nascimento    mysql.NullTime
	Imagens       []Imagem `gorm:"many2many:pessoa_imagem"`
	Participacoes []Filme  `gorm:"many2many:pessoa_filme"`

	ID uint
}
