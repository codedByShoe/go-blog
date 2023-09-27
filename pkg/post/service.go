package post

type Service interface {
	GetAllPosts() ([]Post, error)
	GetPostByID(id int) (*Post, error)
	CreatePost(p *Post) error
	UpdatePost(p *Post) error
	DeletePost(id int) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetAllPosts() ([]Post, error) {
	// Here you could have some business logic before calling repo.
	return s.repo.GetAllPosts()
}

func (s *service) GetPostByID(id int) (*Post, error) {
	// Example of some business logic: Validating ID before calling repo.
	if id <= 0 {
		return nil, ErrInvalidID
	}
	return s.repo.GetPostByID(id)
}

func (s *service) CreatePost(p *Post) error {
	// You can have some business logic here, like validating the post details.
	return s.repo.CreatePost(p)
}

func (s *service) UpdatePost(p *Post) error {
	// You can have some business logic here, like validating the post details.
	return s.repo.UpdatePost(p)
}

func (s *service) DeletePost(id int) error {
	// You can have some business logic here, like validating the ID.
	return s.repo.DeletePost(id)
}
