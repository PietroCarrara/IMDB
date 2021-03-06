package model

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Pessoa struct {
	Nome          string
	Nascimento    mysql.NullTime `gorm:"type:date"`
	Imagens       []Imagem       `gorm:"many2many:pessoa_imagem"`
	Participacoes []Filme        `gorm:"many2many:pessoa_filme"`

	ID uint
}

// ProfilePic retorna a primeira
// imagem da pessoa
func (p Pessoa) ProfilePic() Imagem {
	return p.Imagens[0]
}

// Nasc formata a data de
// nascimento de uma pessoa
func (p Pessoa) Nasc() string {
	return p.Nascimento.Time.Format("02/01/2006 ")
}

// Load carrega a pessoa
// a partir do banco de dados
func (p *Pessoa) Load(db *gorm.DB) {
	db.Preload("Imagens").Preload("Participacoes").First(p)

	for i := 0; i < len(p.Participacoes); i++ {
		db.Preload("Imagens").First(&p.Participacoes[i])
	}
}
