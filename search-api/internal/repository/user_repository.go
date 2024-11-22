package repository

import (
	"database/sql"
	"search-api/internal/domain"
)

type UserRepositoryPostgres struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryPostgres {
	return &UserRepositoryPostgres{db: db}
}

func (r *UserRepositoryPostgres) SearchByName(name string, limit, offset int) ([]domain.User, error) {

	query := `
		SELECT id, name, phone_number, country
		FROM users
		WHERE name ILIKE $1
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.Query(query, name+"%", limit, offset)
	if err != nil {
		return []domain.User{}, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Country); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
