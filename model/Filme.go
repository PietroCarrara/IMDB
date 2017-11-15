package model

import (
	"database/sql"
	"log"
)

type Filme struct {
	Id        int64
	Traducoes []Traducao
	Tags      []Tag
	Sinopse   string

	db *sql.DB
}

// Comentarios retorna os comentarios deste filme
func (self *Filme) Comentarios() []*Comentario {

	users := LoadAllUsers(self.db)

	var comentarios []*Comentario

	for _, u := range users {
		for _, c := range u.Comentarios {
			if c.Alvo.Id == self.Id {
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
		if i.Filme.Id == self.Id {
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
	users := LoadAllUsers(self.db)

	var soma, total float32

	for _, user := range users {
		for _, aval := range user.Avaliacoes {
			if aval.Filme.Id == self.Id {
				soma += aval.Nota
				total++
				break
			}
		}
	}

	// Evitar divisão por zero
	if total < 1 {
		total = 1
	}

	return soma / total
}

// LoadFilme carrega o filme de Id informado
func LoadFilme(db *sql.DB, Id int) *Filme {
	ps, err := db.Prepare("SELECT Id, sinopse FROM filme WHERE Id = ?")
	defer ps.Close()
	if err != nil {
		log.Printf("Erro ao preparar o PS em 'LoadFilme(db, %d)': %s", Id, err.Error())
		return nil
	}

	res, err := ps.Query(Id)
	defer res.Close()
	if err != nil {
		log.Print("Erro ao executar o PS em 'LoadFilme(db, %d)': %s", Id, err.Error())
		return nil
	}

	if res.Next() {

		var filme Filme

		res.Scan(&filme.Id, &filme.Sinopse)

		filme.db = db

		filme.Traducoes = LoadTraducoesByFilme(db, filme)
		filme.Tags = LoadTagsByFilme(db, filme)

		return &filme
	}

	return nil
}

// LoadFilme carrega o filme de Id informado
func LoadAllFilmes(db *sql.DB) []*Filme {

	res, err := db.Query("SELECT Id, sinopse FROM filme")
	defer res.Close()
	if err != nil {
		log.Print("Erro ao executar o PS em 'LoadFilme(db)': %s", err.Error())
		return nil
	}

	var filmes []*Filme

	for res.Next() {

		var filme Filme

		res.Scan(&filme.Id, &filme.Sinopse)

		filme.db = db

		filme.Traducoes = LoadTraducoesByFilme(db, filme)
		filme.Tags = LoadTagsByFilme(db, filme)

		filmes = append(filmes, &filme)
	}

	return filmes
}
