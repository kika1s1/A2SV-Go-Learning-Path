package data

import "example/web-service-gin/models"

var Albums = []models.Album{
	{
		ID:     "1",
		Title:  "Blue Train",
		Artist: "John Coltrane",
		Price:  56.99,
	},
	{
		ID:     "2",
		Title:  "Jeru",
		Artist: "Gerry Mulligan",
		Price:  17.99,
	},
	{
		ID:     "3",
		Title:  "Sarah Vaughan and Clifford Brown",
		Artist: "Sarah Vaughan",
		Price:  39.99,
	},
	{
		ID:     "4",
		Title:  "The Magic Flute",
		Artist: "Graham Norton",
		Price:  39.99,
	},
}
