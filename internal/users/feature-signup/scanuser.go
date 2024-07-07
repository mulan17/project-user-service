package Signup

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// Якщо потрібно отримувати дані про юзера з вводу користувача з консолі

func ScanUser() ([]byte, error) {
	var user User

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter user information:")

	fmt.Print("Firstname: ")
	user.PersonalInfo.Firstname, _ = reader.ReadString('\n')
	user.PersonalInfo.Firstname = strings.TrimSpace(user.PersonalInfo.Firstname)

	fmt.Print("Lastname: ")
	user.PersonalInfo.Lastname, _ = reader.ReadString('\n')
	user.PersonalInfo.Lastname = strings.TrimSpace(user.PersonalInfo.Lastname)

	fmt.Print("Email: ")
	user.Email, _ = reader.ReadString('\n')
	user.Email = strings.TrimSpace(user.Email)

	fmt.Print("Password: ")
	user.Password, _ = reader.ReadString('\n')
	user.Password = strings.TrimSpace(user.Password)

	fmt.Print("Role (user or admin): ")
	user.Role, _ = reader.ReadString('\n')
	user.Role = strings.TrimSpace(user.Role)

	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Fatal("Error encoding JSON:", err)
	}

	fmt.Println("\nJSON representation:")
	fmt.Println(string(userJSON))
	return userJSON, nil
}
