package model

import (
	"database/sql"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Traducao struct {
	Titulo   string
	Lugar    Local
	DataLanc time.Time

	id int
}

func LoadTraducoesByFilme(db *sql.DB, f Filme) []Traducao {

	ps, err := db.Prepare("SELECT id, id_lugar, data_lanc, titulo FROM traducao WHERE id_filme = ?")
	defer ps.Close()
	if err != nil {
		log.Printf("Erro ao preparar o PS em 'LoadTraducoesByFilme(db, %v)': %s", f, err.Error())
		return nil
	}

	res, err := ps.Query(f.id)
	defer res.Close()
	if err != nil {
		log.Printf("Erro ao executar o PS em 'LoadTraducoesByFilme(db, %v)': %s", f, err.Error())
		return nil
	}

	var traducoes []Traducao

	for res.Next() {
		var tr Traducao
		var data mysql.NullTime

		res.Scan(&tr.id, &tr.Lugar, &data, &tr.Titulo)

		if data.Valid {
			tr.DataLanc = data.Time
		}

		traducoes = append(traducoes, tr)
	}

	return traducoes
}
