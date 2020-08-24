package model

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name  string
	Price float64
	Number int
}

