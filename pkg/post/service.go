package post

import "errors"

var ErrInvalidID = errors.New("invalid ID provided")

type Service interface {
	GetAllPosts() ([]Post, error)
	GetPostByID(id int) (*Post, error)
	CreatePost(title string, content string) error
	UpdatePost(id int, title string, content string) error
	DeletePost(id int) error
}

type postService struct {
	repo Repository
}

func NewPostService(r Repository) Service {
	return &postService{repo: r}
}

func (s *postService) GetAllPosts() ([]Post, error) {
	return s.repo.GetAllPosts()
}

func (s *postService) GetPostByID(id int) (*Post, error) {
	// example of handling business logic
	if id <= 0 {
		return nil, ErrInvalidID
	}
	return s.repo.GetPostByID(id)
}

func (s *postService) CreatePost(title string, content string) error {
	return s.repo.CreatePost(title, content)
}

func (s *postService) UpdatePost(id int, title string, content string) error {
	return s.repo.UpdatePost(id, title, content)
}

func (s *postService) DeletePost(id int) error {
	return s.repo.DeletePost(id)
}
