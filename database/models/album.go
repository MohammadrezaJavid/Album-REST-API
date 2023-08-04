package models

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Title  string  `gorm:"size:255" json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type SwagAlbum struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
