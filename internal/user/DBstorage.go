package User

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(connStr string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}

func (s *PostgresStorage) Create(u User) {
	_, err := s.db.Exec("INSERT INTO users (id, email, password, role) VALUES ($1, $2, $3, $4)", u.ID, u.Email, u.Password, u.Role)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to insert user")
	}
}

func (s *PostgresStorage) GetUsers() []User {
	rows, err := s.db.Query("SELECT id, email, password, role FROM users")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to fetch users")
		return nil
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Email, &u.Password, &u.Role); err != nil {
			log.Fatal().Err(err).Msg("Failed to scan user")
			continue
		}
		users = append(users, u)
	}

	return users
}

func (s *PostgresStorage) Exists(email string) bool {
	var exists bool
	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", email).Scan(&exists)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to check if user exists")
	}
	return exists
}

func (s *PostgresStorage) GetUserById(id string) (User, bool) {
	var user User
	err := s.db.QueryRow("SELECT id, email, password, role FROM users WHERE id=$1", id).Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.Name, &user.Lastname)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to fetch user ID")
		return User{}, false
	}
	return user, true // TODO do we need pinter here
}

func (s *PostgresStorage) UpdateUser(user User, id string) bool {
	_, err := s.db.Exec("UPDATE users SET email=$1, password=$2, role=$3, name=$4, lastname=$5 WHERE id=$6", user.Email, user.Password, user.Role, user.Name, user.Lastname, user.ID)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to update user")
		return false
	}
	return true
}
