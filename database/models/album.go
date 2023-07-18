package models

type Album struct {
	ID     string  `gorm:"primaryKey" json:"id"`
	Title  string  `gorm:"size:255" json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
