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

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetAllPosts() ([]Post, error) {
	return s.repo.GetAllPosts()
}

func (s *service) GetPostByID(id int) (*Post, error) {
	// example of handling business logic
	if id <= 0 {
		return nil, ErrInvalidID
	}
	return s.repo.GetPostByID(id)
}

func (s *service) CreatePost(title string, content string) error {
	return s.repo.CreatePost(title, content)
}

func (s *service) UpdatePost(id int, title string, content string) error {
	return s.repo.UpdatePost(id, title, content)
}

func (s *service) DeletePost(id int) error {
	return s.repo.DeletePost(id)
}
