package Signup

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
)

// Функція створює мок юзерів в форматі JSON
func EmulateUser() ([]byte, error) {
	names := []string{
		"James", "John", "Robert", "Michael", "William",
		"David", "Richard", "Joseph", "Thomas", "Charles",
		"Daniel", "Matthew", "Anthony", "Donald", "Mark",
		"Paul", "Steven", "Andrew", "Kenneth", "Joshua",
	}

	surnames := []string{
		"Smith", "Johnson", "Williams", "Brown", "Jones",
		"Garcia", "Miller", "Davis", "Rodriguez", "Martinez",
		"Hernandez", "Lopez", "Gonzalez", "Wilson", "Anderson",
		"Thomas", "Taylor", "Moore", "Jackson", "Martin",
	}

	alphabet := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

	domains := []string{"gmail.com", "yahoo.com", "outlook.com", "mail.ru", "hotmail.com", "yandex.ru", "icloud.com"}

	roles := []string{"user", "admin"}

	var user User

	user.PersonalInfo.Firstname = names[rand.Intn(len(names))]
	user.PersonalInfo.Lastname = surnames[rand.Intn(len(surnames))]
	user.ID = strconv.Itoa(rand.Intn(100000000))
	user.Email = user.PersonalInfo.Firstname + "-" + user.PersonalInfo.Lastname + "@" + domains[rand.Intn(len(domains))]
	for i := 0; i < 16; i++ {
		user.Password += string(alphabet[rand.Intn(len(alphabet))])
	}
	user.Role = roles[rand.Intn(len(roles))]

	userJSON, err := json.Marshal(user)

	if err != nil {
		return nil, fmt.Errorf("error encoding user JSON", err)
	}

	return userJSON, nil
}
