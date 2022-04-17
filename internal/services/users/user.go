package users

type UserServiceInterface interface {
	Create()
}

func (s *UserService) Create() {
	s.repo.Create()
}
