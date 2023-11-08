package comment

type Service interface {
	AddCommentToPost(postId int, author string, content string) error
}

type commentService struct {
	repo Repository
}

func NewCommentService(r Repository) Service {
	return &commentService{repo: r}
}

func (s *commentService) AddCommentToPost(postId int, author string, content string) error {
	return s.repo.AddComment(postId, author, content)
}
