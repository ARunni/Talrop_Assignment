package repository_test

import (
	"regexp"
	"search-api/internal/domain"
	"search-api/internal/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSearchByName(t *testing.T) {
	// Mock the database and sqlmock
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Initialize the repository
	repo := repository.NewUserRepository(db)

	// Define test inputs and expected outputs
	inputName := "John"
	limit := 10
	offset := 0
	expectedUsers := []domain.User{
		{ID: 1, Name: "John", PhoneNumber: "+1234567890", Country: "USA"},
		{ID: 2, Name: "Johnny", PhoneNumber: "+9876543210", Country: "Canada"},
	}

	// Mock the SQL query
	rows := sqlmock.NewRows([]string{"id", "name", "phone_number", "country"}).
		AddRow(expectedUsers[0].ID, expectedUsers[0].Name, expectedUsers[0].PhoneNumber, expectedUsers[0].Country).
		AddRow(expectedUsers[1].ID, expectedUsers[1].Name, expectedUsers[1].PhoneNumber, expectedUsers[1].Country)

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT id, name, phone_number, country
		FROM users
		WHERE name ILIKE $1
		LIMIT $2 OFFSET $3
	`)).
		WithArgs(inputName+"%", limit, offset).
		WillReturnRows(rows)

	// Call the repository function
	result, err := repo.SearchByName(inputName, limit, offset)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, result)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}
