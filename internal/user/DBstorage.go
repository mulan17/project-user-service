package user

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
	_, err := s.db.Exec("INSERT INTO users (id, email, password, role, name, lastname, blocked) VALUES ($1, $2, $3, $4, $5, $6, $7)", u.ID, u.Email, u.Password, u.Role, u.Name, u.Lastname, u.Blocked)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to insert user")
	}
}

func (s *PostgresStorage) GetUsers() []User {
	rows, err := s.db.Query("SELECT id, email, password, role, name, lastname, blocked FROM users")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to fetch users")
		return nil
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Email, &u.Password, &u.Role, &u.Name, &u.Lastname, &u.Blocked); err != nil {
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
	// err := s.db.QueryRow("SELECT id, email, password, role, name, lastname, blocked FROM users id=&1 email=$2, password=$3, role=$4, name=$5, lastname=$6, blocked=$7 WHERE id=$8", id).Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.Name, &user.Lastname, &user.Blocked, id)
	// err := s.db.QueryRow("SELECT * FROM users WHERE id=$1", id).Scan(&user)
	err := s.db.QueryRow("SELECT id, email, password, role, name, lastname, blocked FROM users WHERE id=$1", id).Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.Name, &user.Lastname, &user.Blocked)

	if err != nil {
		log.Fatal().Err(err).Msg("ДЕВОЧКИ ПРИВІТ")
		return User{}, false
	}
	return user, true // TODO do we need pointer here
}

func (s *PostgresStorage) UpdateUser(user User, id string) bool {
	_, err := s.db.Exec("UPDATE users SET email=$1, password=$2, role=$3, name=$4, lastname=$5, blocked=$6 WHERE id=$7", user.Email, user.Password, user.Role, user.Name, user.Lastname,user.Blocked, id)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to update user")
		return false
	}
	return true
}

func (s *PostgresStorage) BlockUser (id string) bool {
	
	block := true
	
	_, err := s.db.Exec("UPDATE users SET blocked=$1 WHERE id=$2", block, id)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to block user")
		return false
	}
	return true
}