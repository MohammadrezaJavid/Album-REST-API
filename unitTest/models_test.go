package unittest

import (
	"album/database/models"
)

var PutAlbum *models.Album = &models.Album{
	Title:  "photo's akbar",
	Artist: "akbari",
	Price:  9.89,
}

var Users = []*models.User{
	{
		Name:     "Mina",
		Username: "mina",
		Email:    "mina@gmail.com",
		Password: "mina123",
	},
	{
		Name:     "Mohammadreza",
		Username: "mohammadreza",
		Email:    "mohammadreza@gmail.com",
		Password: "mohammadreza123",
	},
	{
		Name:     "Akbar",
		Username: "akbar",
		Email:    "akbar@gmail.com",
		Password: "akbar123",
	},
	{
		Name:     "Meshkat",
		Username: "meshkat",
		Email:    "meshkat@gmail.com",
		Password: "meshkat123",
	},
}

var Albums = []*models.Album{
	{
		Title:  "mina photos",
		Artist: "Mina",
		Price:  99.9,
	},
	{
		Title:  "ali photos",
		Artist: "Ali",
		Price:  100,
	},
	{
		Title:  "akbar photos",
		Artist: "Akbar",
		Price:  60.3,
	},
	{
		Title:  "meshkat photos",
		Artist: "Meshkat",
		Price:  98.9,
	},
}
