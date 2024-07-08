package User

// import "encoding/json"

// type storage interface {
// 	CreateOneUser(userJSON []byte) error
// }

// type Service struct {
// 	s InMemStorage
// }

// func (s *Service) UserSignUp(email, password string) {
// 	user := New(email, password)
// 	userJSON, err := json.Marshal(user)
// 	if err != nil {
// 		return
// 	}

// 	s.s.CreateOneUser(userJSON);
