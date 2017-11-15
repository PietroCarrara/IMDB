package model

// Participante representa o cargo
// de uma pessoa em um filme
type Participante struct {
	Pessoa Pessoa
	Cargos []Cargo
	Filme  Filme

	ID       uint
	FilmeID  uint
	PessoaID uint
}
