package model

type Filme struct {
	id          int
	Traducoes   map[Local]string
	Tags        []Tag
	Comentarios []Comentario
	Pessoas     []Pessoa
}
