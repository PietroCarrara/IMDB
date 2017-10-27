package model

type Cargo int

const (
	Ator Cargo = iota
	Diretor
	Sonoplasta
	Dublador
)

func (self Cargo) String() string {
	switch self {
	case Ator:
		return "Ator"
	case Diretor:
		return "Diretor"
	case Sonoplasta:
		return "Sonoplasta"
	case Dublador:
		return "Dublador"
	default:
		return "Null"
	}
}

// GetCargos retorna que papeis foram
// desempenhados pela pessoa p no filme f
// e nil se não há participação alguma
func GetCargos(p Pessoa, f Filme) []Cargo {
	return nil
}
