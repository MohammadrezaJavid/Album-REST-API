package models

type Album struct {
	ID     int64  `gorm:"primaryKey;autoIncrement"`
	Title  string `gorm:"size:255"`
	Artist string
	Price  float64
}
