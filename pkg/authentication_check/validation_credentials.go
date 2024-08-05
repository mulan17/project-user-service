package authentication_check

import (
	"errors"
	
	"github.com/mulan17/project-user-service/pkg/hashing"
)

func  ValidateCredentials(credentialsPassword, userDBPassword string) error {
	passwordIsValid := hashing.CheckPasswordHash(credentialsPassword, userDBPassword)

	if !passwordIsValid {
		return errors.New("credentials password invalid")
	}

	return nil
}

