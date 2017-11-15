package model

import "github.com/jinzhu/gorm"

// Participante representa o cargo
// de uma pessoa em um filme
type Participante struct {
	Pessoa Pessoa
	Cargos []Cargo `gorm:"many2many:participante_cargo"`
	Filme  Filme

	ID       uint
	FilmeID  uint
	PessoaID uint
}

// Load carrega a participação
// a partir do banco de dados
func (p *Participante) Load(db *gorm.DB) {
	db.Preload("Cargos").First(p)
}
