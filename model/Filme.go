package model

import "github.com/jinzhu/gorm"

// Filme representa um filme
// dentro do site
type Filme struct {
	Sinopse string `gorm:"size:1024"`
	Titulo  string

	Participantes []Pessoa `gorm:"many2many:pessoa_filme"`
	Tags          []Tag    `gorm:"many2many:filme_tag"`
	Imagens       []Imagem
	Comentarios   []Comentario
	Avaliacoes    []Avaliacao

	ID uint
}

// Load carrega o filme
// a partir do banco de dados
func (filme *Filme) Load(db *gorm.DB) {
	db.Preload("Participantes").Preload("Imagens").Preload("Tags").Preload("Comentarios").Preload("Avaliacoes").First(filme)

	for i := 0; i < len(filme.Comentarios); i++ {
		filme.Comentarios[i].Load(db)
	}

	for i := 0; i < len(filme.Participantes); i++ {
		filme.Participantes[i].Load(db)
	}
}

// TagAdd adiciona uma tag a um filme
// se a tag já não existe nele
func (filme *Filme) TagAdd(t Tag) bool {
	for _, tag := range filme.Tags {
		if t.ID == tag.ID {
			return false
		}
	}

	filme.Tags = append(filme.Tags, t)
	return true
}

// Banner retorna a primeira imagem
func (filme Filme) Banner() Imagem {
	return filme.Imagens[0]
}

// Nota faz a média de todas as
// avaliações do filme
func (filme Filme) Nota() float32 {

	var nota float32

	for _, aval := range filme.Avaliacoes {
		nota += aval.Nota
	}

	len := len(filme.Avaliacoes)
	if len < 1 {
		len = 1
	}

	return nota / float32(len)
}

// LoadFilmeSlice carrega um vetor de
// filmes a partir do banco de dados
func LoadFilmeSlice(f []Filme, db *gorm.DB) {
	for i := 0; i < len(f); i++ {
		f[i].Load(db)
	}
}
