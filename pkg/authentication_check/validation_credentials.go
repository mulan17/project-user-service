package authentication_check

import (
	"errors"
	"log"

	"github.com/mulan17/project-user-service/internal/user"
	"github.com/mulan17/project-user-service/pkg/hashing"
)


func  ValidateCredentials(u *user.User, s *user.PostgresStorage) error {
	query := "SELECT id, password FROM users WHERE email = $1"
	row := s.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("credentials invalid")
	}

	log.Printf("id from database %v, id from user input %v", retrievedPassword, u.Password)
	
	passwordIsValid := hashing.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials password invalid")
	}

	return nil
}

