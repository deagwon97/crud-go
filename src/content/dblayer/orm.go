package dblayer

import (
	"database/sql"

	"go-api/content/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"time"
)

type DBORM struct {
	*gorm.DB
}

// DBORM의 생성자
func NewORM(dbengine string, dsn string) (*DBORM, error) {
	sqlDB, err := sql.Open(dbengine, dsn)
	// gorm.Open은 *gorm.DB 타입을 초기화한다.
	gormDB, err := gorm.Open(
		mysql.New(mysql.Config{Conn: sqlDB}),
		&gorm.Config{},
	)
	return &DBORM{
		DB: gormDB,
	}, err
}

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (db *DBORM) GetAllContents(page int, pageSize int) (contents []models.Content, err error) {
	return contents, db.Scopes(Paginate(page, pageSize)).Find(&contents).Error
}

func (db *DBORM) GetContent(id int) (content models.Content, err error) {
	return content, db.First(&content, id).Error
}

func (db *DBORM) AddContent(content models.Content) (models.Content, error) {
	return content, db.Create(&content).Error
}

func (db *DBORM) UpdateContent(id int, content models.Content) (models.Content, error) {
	loc, _ := time.LoadLocation("Asia/Seoul")
	kst := time.Now().In(loc)
	content.UpdatedAt = kst.String()

	var new_content models.Content
	db.Where("content_id = ?", id).First(&new_content)
	err := db.Model(&new_content).Updates(content).Error

	return new_content, err
}

func (db *DBORM) DeleteContent(id int) (models.Content, error) {
	var content models.Content
	db.Where("content_id = ?", id).First(&content)
	return content, db.Delete(&content).Error
}
