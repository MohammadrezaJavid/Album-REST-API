package repositories

import "database/sql"

type Album struct {
	ID     int64
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func SelectAlbums(db *sql.DB) (*[]Album, error) {
	return nil, nil
}

func SelectAlbumByID(ID string, db *sql.DB) (*Album, error) {
	return nil, nil
}

func InsertAlbum(album *Album, db *sql.DB) error {
	return nil
}

func DeleteAlbumByID(ID string, db *sql.DB) (int64, error) {
	return 0, nil
}
