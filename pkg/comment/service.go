package comment

type Service interface {
	AddCommentToPost(postId int, author string, content string) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) AddCommentToPost(postId int, author string, content string) error {
	return s.repo.AddComment(postId, author, content)
}
