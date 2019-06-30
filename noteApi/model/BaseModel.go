package model

import "github.com/jinzhu/gorm"

type BaseNote struct {
	gorm.Model
	Title     string `json:"title"`
}
