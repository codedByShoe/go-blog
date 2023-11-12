package post

import (
	"github.com/codedbyshoe/go-blog/internal/comment"
	"gorm.io/gorm"
)

type Post struct {
	ID       int               `db:"id" gorm:"primaryKey"`
	Title    string            `db:"title"`
	Content  string            `db:"content"`
	Comments []comment.Comment `db:"comments"`
}

type Model interface {
	// Define Methods
	GetAllPosts() ([]Post, error)
	GetPostByID(id uint64) (Post, error)
	CreatePost(post Post) error
	UpdatePost(post Post) error
	DeletePost(id uint64) error
}

type postModel struct {
	db *gorm.DB
}

func NewPostModel(db *gorm.DB) Model {
	return &postModel{db: db}
}

func (m *postModel) GetAllPosts() ([]Post, error) {
	var posts []Post

	tx := m.db.Find(&posts)
	if tx.Error != nil {
		return []Post{}, tx.Error
	}

	return posts, nil
}

func (m *postModel) GetPostByID(id uint64) (Post, error) {
	var post Post
	tx := m.db.Where("id = ?", id).First(&post)
	if tx.Error != nil {
		return Post{}, tx.Error
	}
	return post, nil
}

func (m *postModel) CreatePost(post Post) error {
	tx := m.db.Create(&post)
	return tx.Error
}

func (m *postModel) UpdatePost(post Post) error {
	tx := m.db.Save(&post)
	return tx.Error
}

func (m *postModel) DeletePost(id uint64) error {
	tx := m.db.Unscoped().Delete(&Post{}, id)

	return tx.Error
}
