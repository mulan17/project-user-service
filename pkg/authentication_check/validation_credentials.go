package authentication_check

// import (
// 	"errors"

// 	"github.com/mulan17/project-user-service/internal/user"
// 	"github.com/mulan17/project-user-service/pkg/hashing"
// )


// func  ValidateCredentials(u *user.User) error {
// 	query := "SELECT id, password FROM users WHERE email = ?"
// 	row := db.DB.QueryRow(query, u.Email)

// 	var retrievedPassword string
// 	err := row.Scan(&u.ID, &retrievedPassword)

// 	if err != nil {
// 		return errors.New("Credentials invalid")
// 	}

// 	passwordIsValid := hashing.CheckPasswordHash(u.Password, retrievedPassword)

// 	if !passwordIsValid {
// 		return errors.New("Credentials invalid")
// 	}

// 	return nil
// }

