package model

import "github.com/jinzhu/gorm"

// Loader representa uma struct
// que pode ser carregada a partir
// do banco de dados
type Loader interface {
	Load(*gorm.DB)
}
