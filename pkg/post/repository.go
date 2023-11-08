package post

import (
	"database/sql"

	"github.com/codedbyshoe/go-blog/pkg/comment"
)

type Repository interface {
	// Define Methods
	GetAllPosts() ([]Post, error)
	GetPostByID(id int) (*Post, error)
	CreatePost(title string, content string) error
	UpdatePost(id int, title string, content string) error
	DeletePost(id int) error
}

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) Repository {
	return &postRepository{db: db}
}

func (r *postRepository) GetAllPosts() ([]Post, error) {
	rows, err := r.db.Query("SELECT id, title FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.ID, &p.Title); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func (r *postRepository) GetPostByID(id int) (*Post, error) {
	// Fetch the post
	row := r.db.QueryRow("SELECT id, title, content FROM posts WHERE id = ?", id)
	var p Post
	err := row.Scan(&p.ID, &p.Title, &p.Content)
	if err != nil {
		return nil, err
	}

	// Fetch the comments
	rows, err := r.db.Query("SELECT id, post_id, content, author FROM comments WHERE post_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	p.Comments = []comment.Comment{}
	for rows.Next() {
		var c comment.Comment
		if err := rows.Scan(&c.ID, &c.PostID, &c.Content, &c.Author); err != nil {
			return nil, err
		}
		p.Comments = append(p.Comments, c)
	}

	return &p, nil
}

func (r *postRepository) CreatePost(title string, content string) error {
	_, err := r.db.Exec("INSERT INTO posts (title, content) VALUES (?, ?)", title, content)
	return err
}

func (r *postRepository) UpdatePost(id int, title string, content string) error {
	_, err := r.db.Exec("UPDATE posts SET title = ?, content = ? WHERE id = ?", title, content, id)
	return err
}

func (r *postRepository) DeletePost(id int) error {
	_, err := r.db.Exec("DELETE FROM posts WHERE id = ?", id)
	return err
}
