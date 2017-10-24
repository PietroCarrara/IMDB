package model

type Cargo int

const (
	Ator Cargo = iota
	Diretor
	Sonoplasta
	Dublador
)

// GetCargos retorna que papeis foram
// desempenhados pela pessoa p no filme f
// e nil se não há participação alguma
func GetCargos(p Pessoa, f Filme) []Cargo {
	return nil
}
