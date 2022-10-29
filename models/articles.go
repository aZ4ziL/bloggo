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
	CategoryID uint           `json:"category_id"`
	AuthorID   uint           `json:"author_id"`
	Title      string         `gorm:"size:100" json:"title"`
	Slug       string         `gorm:"size:100;unique;index" json:"slug"`
	Logo       string         `gorm:"size:255;null" json:"logo"`
	Desc       string         `gorm:"size:255" json:"desc"`
	Content    string         `gorm:"type:longtext" json:"content"`
	Status     string         `gorm:"default:DRAFTED" json:"status"`
	Views      int            `gorm:"default:0" json:"views"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime:nano" json:"updated_at"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// CreateNewArticle create new article
func CreateNewArticle(article *Article) error {
	err := db.Create(article).Error
	return err
}

// Being worked on
