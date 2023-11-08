package post

import (
	"github.com/codedbyshoe/go-blog/pkg/comment"
)

type Post struct {
	ID       int
	Title    string
	Content  string
	Comments []comment.Comment
}
