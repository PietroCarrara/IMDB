package model

import "database/sql"
import "log"

type Tag int

const (
	Terror Tag = iota
	SciFi
	Acao
	Animacao
)

func (self Tag) String() string {
	switch self {
	case Terror:
		return "Terror"
	case SciFi:
		return "Ficção Científica"
	case Acao:
		return "Acao"
	case Animacao:
		return "Animacao"
	default:
		return "Null"
	}
}

func LoadTagsByFilme(db *sql.DB, f Filme) []Tag {

	ps, err := db.Prepare("SELECT id_tag FROM filme_tag WHERE id_filme = ?")
	defer ps.Close()
	if err != nil {
		log.Printf("Erro ao preparar o PS em 'LoadTagsByFilme(db, %v)': %s", f, err.Error())
	}

	res, err := ps.Query(f.id)
	defer res.Close()
	if err != nil {
		log.Printf("Erro ao executar o PS em 'LoadTagsByFilme(db, %v)': %s", f, err.Error())
	}

	var tags []Tag

	for res.Next() {

		var tag Tag

		res.Scan(&tag)

		tags = append(tags, tag)
	}

	return tags
}
