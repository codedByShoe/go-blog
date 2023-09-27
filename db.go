package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./blog.db")
	if err != nil {
		panic(err)
	}
}

func getAllPosts() ([]Post, error) {
	rows, err := db.Query("SELECT id, title FROM posts")
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

func getPostByID(id int) (*Post, error) {
	// Fetch the post
	row := db.QueryRow("SELECT id, title, content FROM posts WHERE id = ?", id)
	var p Post
	err := row.Scan(&p.ID, &p.Title, &p.Content)
	if err != nil {
		return nil, err
	}

	// Fetch the comments
	rows, err := db.Query("SELECT id, post_id, content, author FROM comments WHERE post_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	p.Comments = []Comment{}
	for rows.Next() {
		var c Comment
		if err := rows.Scan(&c.ID, &c.PostID, &c.Content, &c.Author); err != nil {
			return nil, err
		}
		p.Comments = append(p.Comments, c)
	}

	return &p, nil
}

func createPost(title string, content string) error {
	_, err := db.Exec("INSERT INTO posts (title, content) VALUES (?, ?)", title, content)
	return err
}

func updatePost(id int, title string, content string) error {
	_, err := db.Exec("UPDATE posts SET title = ?, content = ? WHERE id = ?", title, content, id)
	return err
}

func deletePost(id int) error {
	_, err := db.Exec("DELETE FROM posts WHERE id = ?", id)
	return err
}

func addComment(postId int, content string, author string) error {
	_, err := db.Exec("INSERT INTO comments (post_id, content, author) VALUES (?, ?, ?)", postId, content, author)
	return err
}
