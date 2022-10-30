package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	DRAFTED   = "DRAFTED"
	PUBLISHED = "PUBLISHED"
)

type Article struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	CategoryID uint           `gorm:"comment:Foreign Key for category" json:"category_id" validate:"required"`
	AuthorID   uint           `gorm:"comment:Foreign Key for user" json:"author_id" validate:"required"`
	Title      string         `gorm:"size:100" json:"title" validate:"required"`
	Slug       string         `gorm:"size:100;unique;index" json:"slug" validate:"required"`
	Logo       string         `gorm:"size:255;null" json:"logo"`
	Desc       string         `gorm:"size:255" json:"desc" validate:"required"`
	Content    string         `gorm:"type:longtext" json:"content" validate:"required"`
	Status     string         `gorm:"default:DRAFTED" json:"status" validate:"required"`
	Views      int            `gorm:"default:0" json:"views"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime:nano" json:"updated_at"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Comments   []Comment      `gorm:"foreignKey:ArticleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
}

// CreateNewArticle create new article
func CreateNewArticle(article *Article) error {
	err := db.Create(article).Error
	return err
}

// GetArticleByID get article by passing the ID
func GetArticleByID(id uint) (Article, error) {
	var article Article
	err := db.Model(&Article{}).Where("id = ?", id).Preload("Comments").First(&article).Error
	return article, err
}

// GetArticleBySlug get article by passing the SLUG
func GetArticleBySlug(slug string) (Article, error) {
	var article Article
	err := db.Model(&Article{}).Where("slug = ?", slug).Preload("Comments").First(&article).Error
	return article, err
}

// GetAllArticles get all articles
func GetAllArticles() []Article {
	var articles []Article
	db.Model(&Article{}).Find(&articles)
	return articles
}
