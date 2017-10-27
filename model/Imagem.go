package model

import (
	"database/sql"
	"log"
)

type Imagem struct {
	Caminho string
	Filme   *Filme

	Id int
	db *sql.DB
}

func LoadAllImagens(db *sql.DB) []Imagem {
	res, err := db.Query("SELECT id, caminho, id_filme FROM imagem")
	defer res.Close()
	if err != nil {
		log.Printf("Erro ao executar o PS em 'LoadAllImagens(db)': %s", err.Error())
	}

	var imagens []Imagem

	for res.Next() {
		var i Imagem
		var idFilme int

		i.db = db

		res.Scan(&i.Id, &i.Caminho, &idFilme)

		i.Filme = LoadFilme(db, idFilme)

		imagens = append(imagens, i)
	}

	return imagens
}

func LoadImagem(db *sql.DB, id int) *Imagem {
	ps, err := db.Prepare("SELECT caminho, id_filme FROM imagem WHERE id = ?")
	defer ps.Close()
	if err != nil {
		log.Printf("Erro ao preparar o PS em 'LoadImagem(db, %d)': %s", id, err.Error())
	}

	res, err := ps.Query(id)
	defer res.Close()
	if err != nil {
		log.Printf("Erro ao executar o PS em 'LoadImagem(db, %d)': %s", id, err.Error())
	}

	for res.Next() {
		var i Imagem
		var idFilme int

		i.db = db
		i.Id = id

		res.Scan(&i.Caminho, &idFilme)

		i.Filme = LoadFilme(db, idFilme)

		return &i
	}

	return nil
}

func (self *Imagem) Pessoas() []Pessoa {
	pessoas := LoadAllPessoas(self.db)

	var res []Pessoa

	for _, p := range pessoas {
		for _, i := range p.Imagens {
			if i.Id == self.Id {
				res = append(res, p)
				break
			}
		}
	}

	return res
}
