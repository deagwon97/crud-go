package dblayer

import (
	"database/sql"

	"go-api/content/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBORM struct {
	*gorm.DB
}

func NewORM(dbengine string, dsn string) (*DBORM, error) {
	sqlDB, err := sql.Open(dbengine, dsn)
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	return &DBORM{
		DB: gormDB,
	}, err

}

func (db *DBORM) GetAllContents() (contents []models.Content, err error) {
	return contents, db.Find(&contents).Error
}

func (db *DBORM) GetContent(id int) (content models.Content, err error) {
	return content, db.First(&content, id).Error
}
