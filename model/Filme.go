package model

import (
	"database/sql"
	"log"
)

type Filme struct {
	Traducoes []Traducao
	Tags      []Tag

	db *sql.DB
	id int
}

// Comentarios retorna os comentarios deste filme
func (self *Filme) Comentarios() []*Comentario {

	users := LoadAllUsers(self.db)

	var comentarios []*Comentario

	for _, u := range users {
		for _, c := range u.Comentarios {
			if c.Alvo.id == self.id {
				comentarios = append(comentarios, c)
			}
		}
	}

	return comentarios
}

func (self *Filme) Imagens() []Imagem {
	imgs := LoadAllImagens(self.db)

	var res []Imagem

	for _, i := range imgs {
		if i.Filme.id == self.id {
			res = append(res, i)
		}
	}

	return res
}

func (self *Filme) Banner() Imagem {
	imgs := self.Imagens()

	var i int

	menor := imgs[0].Id

	for index, val := range imgs[1:] {
		if val.Id < menor {
			i = index
			menor = val.Id
		}
	}

	return imgs[i]
}

func (self *Filme) Name() string {
	return self.Traducoes[Brasil].Titulo
}

func (self *Filme) Pessoas() []Pessoa {
	return nil
}

func (self *Filme) Nota() float32 {

	return 0
}

// LoadFilme carrega o filme de id informado
func LoadFilme(db *sql.DB, id int) *Filme {
	ps, err := db.Prepare("SELECT id FROM filme WHERE id = ?")
	defer ps.Close()
	if err != nil {
		log.Printf("Erro ao preparar o PS em 'LoadFilme(db, %d)': %s", id, err.Error())
		return nil
	}

	res, err := ps.Query(id)
	defer res.Close()
	if err != nil {
		log.Print("Erro ao executar o PS em 'LoadFilme(db, %d)': %s", id, err.Error())
		return nil
	}

	if res.Next() {

		filme := Filme{id: id}

		filme.db = db

		filme.Traducoes = LoadTraducoesByFilme(db, filme)
		filme.Tags = LoadTagsByFilme(db, filme)

		return &filme
	}

	return nil
}
