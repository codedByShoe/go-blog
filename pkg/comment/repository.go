package comment

import "database/sql"

type Repository interface {
	// Define Methods
	AddComment(postId int, content string, author string) error
}

type commentRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &commentRepository{db: db}
}

func (r *commentRepository) AddComment(postId int, content string, author string) error {
	_, err := r.db.Exec("INSERT INTO comments (post_id, content, author) VALUES (?, ?, ?)", postId, content, author)
	return err
}
