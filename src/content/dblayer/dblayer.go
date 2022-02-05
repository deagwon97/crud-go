package dblayer

import "go-api/content/models"

type DBLayer interface {
	GetAllContents() ([]models.Content, error)
	GetContent(int) (models.Content, error)
}
