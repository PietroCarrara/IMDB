package model

// Imagem representa uma tela
// ou poster de um filme
type Imagem struct {
	Caminho string
	Pessoas []Pessoa `gorm:"many2many:pessoa_imagem"`

	FilmeID uint
	ID      uint
}
