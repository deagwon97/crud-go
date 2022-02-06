package rest

import (
	"go-api/content/dblayer"
	"go-api/content/models"
	"go-api/database"
	"strconv"

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

	p := c.Param("id")
	id, err := strconv.Atoi(p)

	content, err := h.db.GetContent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, content)
}

func (h *Handler) AddContent(c *gin.Context) {

	// p := c.Param("id")
	// id, err := strconv.Atoi(p)
	var content_data models.Content

	err := c.ShouldBindJSON(&content_data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	content, err := h.db.AddContent(content_data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, content)
}

type HandlerInterface interface {
	GetContents(c *gin.Context)
	GetContent(c *gin.Context)
	AddContent(c *gin.Context)
}

// HandlerInterface의 생성자
func NewHandler() (HandlerInterface, error) {
	dsn := database.DataSource
	// DBORM 초기화
	db, err := dblayer.NewORM("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &Handler{
		db: db,
	}, nil
}
