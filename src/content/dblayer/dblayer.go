package dblayer

import "go-api/content/models"

type DBLayer interface {
	GetAllContents(int, int) ([]models.Content, error)
	GetContent(int) (models.Content, error)
	AddContent(models.Content) (models.Content, error)
	UpdateContent(int, models.Content) (models.Content, error)
	DeleteContent(int) (models.Content, error)
}
