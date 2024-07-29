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
		log.Error().Err(err).Msg("Failed to insert user")
	}
}

func (s *PostgresStorage) GetUsers() []User {
	rows, err := s.db.Query("SELECT id, email, password, role FROM users")
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch users")
		return nil
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Email, &u.Password, &u.Role); err != nil {
			log.Error().Err(err).Msg("Failed to scan user")
			continue
		}
		users = append(users, u)
	}

	return users
}
