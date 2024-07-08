package User

import (
	"sync"
)

type Storage struct {
	usersM sync.Mutex
	users  []User
}

func (s *Storage) Create(u User) {
	s.usersM.Lock()
	defer s.usersM.Unlock()

	s.users = append(s.users, u)
}
