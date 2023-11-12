package comment

import (
	"gorm.io/gorm"
)

type Comment struct {
	ID      uint64 `db:"id" gorm:"primaryKey"`
	PostID  uint64 `db:"post_id"`
	Content string `db:"content"`
	Author  string `db:"author"`
}

type Model interface {
	AddComment(comment Comment) error
}

type commentModel struct {
	db *gorm.DB
}

func NewCommentModel(db *gorm.DB) Model {
	return &commentModel{db: db}
}

func (m *commentModel) AddComment(comment Comment) error {
	tx := m.db.Create(comment)
	return tx.Error
}
