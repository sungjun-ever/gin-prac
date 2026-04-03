package service

type BookService struct{}

func (s *BookService) GetMemoList() []string {
	return []string{"메모1", "메모2"}
}
