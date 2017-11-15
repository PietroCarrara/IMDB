package model

import "github.com/jinzhu/gorm"

// Tag representa uma categoria
// à qual algum filme pertence
// (Terror, ação, etc...)
type Tag struct {
	Titulo string
	Filmes []Filme `gorm:"many2many:filme_tag"`

	ID uint
}

// Load carrega a tag
// a partir do banco de dados
func (t *Tag) Load(db *gorm.DB) {
	db.Preload("Filmes").Where(&Tag{ID: t.ID}).First(t)
}
