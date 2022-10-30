package models

import "time"

type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	FullName  string    `gorm:"size:100" json:"full_name" validate:"required"`
	Email     string    `gorm:"size:100" json:"email" validate:"required,email"`
	ArticleID uint      `gorm:"comment:Foreign Key for Article ID" json:"article_id" validate:"required,number"`
	Content   string    `gorm:"type:longtext" json:"content" validate:"required"`
	Approved  bool      `gorm:"default:0" json:"approved"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:nano" json:"updated_at"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// CreateNewComment create new comment
func CreateNewComment(comment *Comment) error {
	err := db.Create(comment).Error
	return err
}

// GetAllCommentsByArticleID get all comments by article ID
func GetAllCommentsByArticleID(id uint) []Comment {
	var comments []Comment
	db.Model(&Comment{}).Where("article_id = ?", id).Find(&comments)
	return comments
}

// GetCommentByID get comment by passing the ID
func GetCommentByID(id uint) (Comment, error) {
	var comment Comment
	err := db.Model(&Comment{}).Where("id = ?", id).First(&comment).Error
	return comment, err
}
