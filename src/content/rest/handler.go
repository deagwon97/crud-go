package rest

import (
	"go-api/content/dblayer"
	"go-api/content/models"
	"go-api/database"
	"net/http"

	"github.com/gin-gonic/gin"

	"strconv"
)

type Handler struct {
	db dblayer.DBLayer
}

func (h *Handler) GetContents(c *gin.Context) {
	page, _ := strconv.Atoi(c.Param("id"))
	pageSize, _ := strconv.Atoi(c.Param("id"))

	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "dsn 오류"})
		return
	}
	contents, err := h.db.GetAllContents(page, pageSize)
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

func (h *Handler) UpdateContent(c *gin.Context) {

	p := c.Param("id")
	id, err := strconv.Atoi(p)
	var content_data models.Content

	err = c.ShouldBindJSON(&content_data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	content, err := h.db.UpdateContent(id, content_data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, content)
}

func (h *Handler) DeleteContent(c *gin.Context) {

	p := c.Param("id")
	id, err := strconv.Atoi(p)

	content, err := h.db.DeleteContent(id)
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
	UpdateContent(c *gin.Context)
	DeleteContent(c *gin.Context)
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
