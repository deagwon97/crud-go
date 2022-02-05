package rest

import (
	"go-api/content/dblayer"
	"go-api/database"

	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	db dblayer.DBLayer
}

func (h *Handler) GetContents(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "dsn 오류"})
		return
	}
	contents, err := h.db.GetAllContents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contents)
}

func (h *Handler) GetContent(c *gin.Context) {
	content, err := h.db.GetContent(0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, content)
}

type HandlerInterface interface {
	GetContents(c *gin.Context)
}

func NewHandler() (HandlerInterface, error) {
	dsn := database.DataSource
	db, err := dblayer.NewORM("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &Handler{
		db: db,
	}, nil
}
