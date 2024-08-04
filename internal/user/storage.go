package user

import (
	"database/sql"
	"fmt"
	

	"strings"
	_ "github.com/lib/pq"

	"github.com/rs/zerolog/log"
	"github.com/mulan17/project-user-service/pkg/authentication_check"
)

type PostgresStorage struct {
	DB *sql.DB
}

func NewPostgresStorage(connStr string) (*PostgresStorage, error) {
	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("opening database: %v", err)
	}

	return &PostgresStorage{DB: DB}, nil
}

func (s *PostgresStorage) Create(u User) error {
	_, err := s.DB.Exec("INSERT INTO users (email, password, role, name, lastname, status) VALUES ($1, $2, $3, $4, $5, $6)", u.Email, u.Password, u.Role, u.Name, u.Lastname, u.Status)
	if err != nil {
		// log.Fatal().Err(err).Msg("Failed to insert user")
		return fmt.Errorf("inserting user: %v", err)
	}
	return nil
}

func (s *PostgresStorage) GetUsers() ([]User, error) {
	rows, err := s.DB.Query("SELECT id, email, password, role, name, lastname, status FROM users")
	if err != nil {
		// log.Fatal().Err(err).Msg("Failed to fetch users")
		return nil, fmt.Errorf("querying users: %v", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Email, &u.Password, &u.Role, &u.Name, &u.Lastname, &u.Status); err != nil {
			log.Fatal().Err(err).Msg("Failed to scan user")
			continue
		}
		users = append(users, u)
	}

	return users, nil
}

func (s *PostgresStorage) Login(email, password string) (User, error) {
	var user User
	err := s.DB.QueryRow("SELECT id, email, password, role, name, lastname, status FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.Name, &user.Lastname, &user.Status)
	if err != nil {
		return User{}, fmt.Errorf("Can't find user %v", err)
	} 
	log.Printf("(66)User from DB: %v", user)
	err = authentication_check.ValidateCredentials(password, user.Password)
	if err != nil {
		return User{}, fmt.Errorf("Wrong password %v", err)
	}
	return user, nil
}

func (s *PostgresStorage) Exists(email string) (bool, error) {
	var exists bool
	err := s.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", email).Scan(&exists)
	if err != nil {
		// log.Fatal().Err(err).Msg("Failed to check if user exists")
		return false, fmt.Errorf("checking if users exist: %v", err)
	}
	return exists, nil
}

func (s *PostgresStorage) GetUserById(id string) (User, error) {
	var user User
	err := s.DB.QueryRow("SELECT id, email, password, role, name, lastname, status FROM users WHERE id=$1", id).Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.Name, &user.Lastname, &user.Status)

	if err != nil {
		// log.Fatal().Err(err).Msg("ДЕВОЧКИ ПРИВІТ")
		return User{}, fmt.Errorf("querying user: %v", err)
	}
	return user, nil // TODO do we need pointer here
}

// func (s *PostgresStorage) UpdateUser(user User, id string) error {
// 	_, err := s.DB.Exec("UPDATE users SET email=$1, password=$2, role=$3, name=$4, lastname=$5, status=$6 WHERE id=$7", user.Email, user.Password, user.Role, user.Name, user.Lastname, user.Status, id)
// 	if err != nil {
// 		// log.Fatal().Err(err).Msg("Failed to update user")
// 		return fmt.Errorf("updating user: %v", err)
// 	}
// 	return nil
// }

func (s *PostgresStorage) UpdateUser(user User, id string) error {
    query := "UPDATE users SET"
    var updates []string
    var args []interface{}
    var argIdx int = 1

    if user.Email != "" {
        updates = append(updates, fmt.Sprintf(" email=$%d", argIdx))
        args = append(args, user.Email)
        argIdx++
    }
    if user.Password != "" {
        updates = append(updates, fmt.Sprintf(" password=$%d", argIdx))
        args = append(args, user.Password)
        argIdx++
    }
    if user.Role != "" {
        updates = append(updates, fmt.Sprintf(" role=$%d", argIdx))
        args = append(args, user.Role)
        argIdx++
    }
    if user.Name != "" {
        updates = append(updates, fmt.Sprintf(" name=$%d", argIdx))
        args = append(args, user.Name)
        argIdx++
    }
    if user.Lastname != "" {
        updates = append(updates, fmt.Sprintf(" lastname=$%d", argIdx))
        args = append(args, user.Lastname)
        argIdx++
    }
    if user.Status != "" {
        updates = append(updates, fmt.Sprintf(" status=$%d", argIdx))
        args = append(args, user.Status)
        argIdx++
    }

    // If no updates are provided, return an error
    if len(updates) == 0 {
        return fmt.Errorf("no fields to update")
    }

    // Join the updates and add the WHERE clause
    query += strings.Join(updates, ",") + fmt.Sprintf(" WHERE id=$%d", argIdx)
    args = append(args, id)

    // Execute the query
    _, err := s.DB.Exec(query, args...)
    if err != nil {
        return fmt.Errorf("updating user: %v", err)
    }

    return nil
}


func (s *PostgresStorage) BlockUser(id string) error {

	status := "blocked"

	_, err := s.DB.Exec("UPDATE users SET status=$1 WHERE id=$2", status, id)
	if err != nil {
		// log.Fatal().Err(err).Msg("Failed to block user")
		return fmt.Errorf("blocking user: %v", err)
	}
	return nil
}

func (s *PostgresStorage) LimitUser(id string) error {

	status := "limited"

	_, err := s.DB.Exec("UPDATE users SET status=$1 WHERE id=$2", status, id)
	if err != nil {
		// log.Fatal().Err(err).Msg("Failed to block user")
		return fmt.Errorf("limiting user: %v", err)
	}
	return nil
}
