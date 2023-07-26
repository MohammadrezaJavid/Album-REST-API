package unittest

import (
	"album/database/models"
)

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
}

var Albums = []*models.Album{
	{
		ID:     "1",
		Title:  "mina photos",
		Artist: "Mina",
		Price:  99.9,
	},
	{
		ID:     "2",
		Title:  "ali photos",
		Artist: "Ali",
		Price:  100,
	},
	{
		ID:     "3",
		Title:  "akbar photos",
		Artist: "Akbar",
		Price:  60.3,
	},
	{
		ID:     "4",
		Title:  "meshkat photos",
		Artist: "Meshkat",
		Price:  98.9,
	},
}
