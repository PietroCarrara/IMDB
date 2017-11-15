package model

// Imagem representa uma tela
// ou poster de um filme
type Imagem struct {
	Caminho string
	FilmeID uint
	Pessoas []Pessoa `gorm:"many2many:pessoa_imagem"`

	ID uint
}
