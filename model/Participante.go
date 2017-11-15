package model

// Participante representa o cargo
// de uma pessoa em um filme
type Participante struct {
	Pessoa Pessoa
	Cargos []Cargo `gorm:"many2many:participante_cargo"`
	Filme  Filme

	ID       uint
	FilmeID  uint
	PessoaID uint
}
