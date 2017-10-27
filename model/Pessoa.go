package model

import (
	"database/sql"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Pessoa struct {
	Filmes     []Filme
	Imagens    []Imagem
	Nome       string
	Nascimento time.Time
	id         int
}

func LoadAllPessoas(db *sql.DB) []Pessoa {
	res, err := db.Query("SELECT id, nome, data_nasc FROM pessoa")
	defer res.Close()
	if err != nil {
		log.Printf("Erro ao executar o PS em 'LoadAllPessoas(db)': %s", err.Error())
	}

	var pessoas []Pessoa

	for res.Next() {
		var (
			p Pessoa
			d mysql.NullTime
		)

		res.Scan(&p.id, &p.Nome, &d)

		if d.Valid {
			p.Nascimento = d.Time
		}

		p.load(db)

		pessoas = append(pessoas, p)
	}

	return pessoas
}

func LoadPessoaByID(db *sql.DB, id int) *Pessoa {
	ps, err := db.Prepare("SELECT id, nome, data_nasc FROM pessoa WHERE id = ?")
	defer ps.Close()
	if err != nil {
		log.Printf("Erro ao preparar o PS em 'LoadPessoaByID(db, %d)': %s", id, err.Error())
	}

	res, err := ps.Query(id)
	defer res.Close()
	if err != nil {
		log.Printf("Erro ao executar o PS em 'LoadPessoaByID(db, %d)': %s", id, err.Error())
	}

	if res.Next() {
		var (
			p Pessoa
			d mysql.NullTime
		)

		res.Scan(&p.id, &p.Nome, &d)

		if d.Valid {
			p.Nascimento = d.Time
		}

		p.load(db)

		return &p
	}

	return nil
}

func (self *Pessoa) load(db *sql.DB) {
	self.loadFilmes(db)
	self.loadImagens(db)
}

func (self *Pessoa) loadFilmes(db *sql.DB) {
	ps, err := db.Prepare("SELECT id_filme FROM pessoa_filme WHERE id_pessoa = ?")
	defer ps.Close()
	if err != nil {
		log.Printf("Erro ao preparar o PS em 'loadFilmes(db)': %s", err.Error())
	}

	res, err := ps.Query(self.id)
	defer res.Close()
	if err != nil {
		log.Printf("Erro ao executar o PS em 'loadFilmes(db)': %s", err.Error())
	}

	for res.Next() {
		var id int

		res.Scan(&id)

		f := *LoadFilme(db, id)

		self.Filmes = append(self.Filmes, f)
	}
}

func (self *Pessoa) loadImagens(db *sql.DB) {
	ps, err := db.Prepare("SELECT id_imagem FROM pessoa_imagem WHERE id_pessoa = ?")
	defer ps.Close()
	if err != nil {
		log.Printf("Erro ao preparar o PS em 'loadImagens(db)': %s", err.Error())
	}

	res, err := ps.Query(self.id)
	defer res.Close()
	if err != nil {
		log.Printf("Erro ao executar o PS em 'loadImagens(db)': %s", err.Error())
	}

	for res.Next() {
		var id int

		res.Scan(&id)

		f := *LoadImagem(db, id)

		self.Imagens = append(self.Imagens, f)
	}
}
