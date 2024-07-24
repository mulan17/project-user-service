package User

type storage interface {
	Create(u User)
	GetUsers() []User
}

type Service struct {
	s storage
}

func NewService(s storage) *Service {
	return &Service{
		s: s,
	}
}

func (s *Service) SignUp(email, password string) {
	// complex logic of gathering user data
	user := New(email, password)

	//user notifictaions: emails, sms etc
	s.s.Create(user)

	//sending internal events to notify other services that user was created
	//etc
}

func (s *Service) GetUsers() []User {
	return s.s.GetUsers()
}
