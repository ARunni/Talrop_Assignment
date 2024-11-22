package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"search-api/internal/domain"
	"search-api/internal/handler"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock Usecase to simulate the SearchUsersByName method
type MockUsecase struct {
	mock.Mock
}

func (m *MockUsecase) SearchUsersByName(name string, limit, offset int) ([]domain.User, error) {
	args := m.Called(name, limit, offset)
	return args.Get(0).([]domain.User), args.Error(1)
}

func TestSearchHandler_SearchUsers(t *testing.T) {
	// Test data
	name := "John"
	page := 1
	limit := 100
	offset := (page - 1) * limit

	users := []domain.User{
		{ID: 1, Name: "John", PhoneNumber: "+1234567890", Country: "USA"},
		{ID: 2, Name: "Johnny", PhoneNumber: "+9876543210", Country: "Canada"},
	}

	// Mock Usecase
	mockUsecase := new(MockUsecase)
	mockUsecase.On("SearchUsersByName", name, limit, offset).Return(users, nil)

	// Create handler with the mock Usecase
	handler := handler.NewSearchHandler(mockUsecase)

	// Create a mock HTTP request
	req, err := http.NewRequest("GET", "/search?name=John&page=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the HTTP response
	rr := httptest.NewRecorder()
	handler.SearchUsers(rr, req)

	// Assertions
	assert.Equal(t, http.StatusOK, rr.Code) // Check if status code is 200 OK

	var response struct {
		Total int           `json:"total"`
		Users []domain.User `json:"users"`
	}

	// Decode the JSON response
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	// Assert the total number of users
	assert.Equal(t, len(users), response.Total)

	// Assert the users' details
	assert.Equal(t, users[0].Name, response.Users[0].Name)
	assert.Equal(t, users[1].Name, response.Users[1].Name)

	// Verify that the mock Usecase was called with the expected arguments
	mockUsecase.AssertExpectations(t)
}

func TestSearchHandler_SearchUsers_Error(t *testing.T) {
	// Test data
	name := "John"
	page := 1
	limit := 100
	offset := (page - 1) * limit

	// Mock Usecase with an error, but return an empty slice to avoid panic
	mockUsecase := new(MockUsecase)
	mockUsecase.On("SearchUsersByName", name, limit, offset).Return([]domain.User{}, errors.New("error retrieving users"))

	// Create handler with the mock Usecase
	handler := handler.NewSearchHandler(mockUsecase)

	// Create a mock HTTP request
	req, err := http.NewRequest("GET", "/search?name=John&page=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the HTTP response
	rr := httptest.NewRecorder()
	handler.SearchUsers(rr, req)

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, rr.Code) // Check if status code is 500 Internal Server Error
	assert.Contains(t, rr.Body.String(), "Error retrieving users") // Check error message in body

	// Verify that the mock Usecase was called with the expected arguments
	mockUsecase.AssertExpectations(t)
}


func TestSearchHandler_MissingNameParam(t *testing.T) {
	// Create a mock HTTP request without the 'name' query parameter
	req, err := http.NewRequest("GET", "/search?page=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the HTTP response
	rr := httptest.NewRecorder()
	handler := handler.NewSearchHandler(nil) // Use nil since we expect an error
	handler.SearchUsers(rr, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, rr.Code) // Check if status code is 400 Bad Request
	assert.Contains(t, rr.Body.String(), "Missing 'name' query parameter") // Check error message
}

func TestSearchHandler_MissingPageParam(t *testing.T) {
	// Create a mock HTTP request without the 'page' query parameter
	req, err := http.NewRequest("GET", "/search?name=John", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the HTTP response
	rr := httptest.NewRecorder()
	handler := handler.NewSearchHandler(nil) // Use nil since we expect an error
	handler.SearchUsers(rr, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, rr.Code) // Check if status code is 400 Bad Request
	assert.Contains(t, rr.Body.String(), "Missing 'page' query parameter") // Check error message
}
