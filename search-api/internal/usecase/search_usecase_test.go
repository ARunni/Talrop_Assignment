package usecase_test

import (
	"errors"
	"search-api/internal/domain"
	"search-api/internal/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of the UserRepository interface
type MockUserRepository struct {
	mock.Mock
}

// SearchByName mocks the SearchByName method from the UserRepository interface
func (m *MockUserRepository) SearchByName(name string, limit, offset int) ([]domain.User, error) {
	args := m.Called(name, limit, offset)
	return args.Get(0).([]domain.User), args.Error(1)
}

// MockHelper is a mock implementation of the Helper interface
type MockHelper struct {
	mock.Mock
}

// RankUsers mocks the RankUsers method from the Helper interface
func (m *MockHelper) RankUsers(inputName string, users []domain.User) []domain.User {
	args := m.Called(inputName, users)
	return args.Get(0).([]domain.User)
}

// IsAlphabetic mocks the IsAlphabetic method from the Helper interface
func (m *MockHelper) IsAlphabetic(name string) bool {
	args := m.Called(name)
	return args.Bool(0)
}

func TestSearchUsersByName(t *testing.T) {
	// Test data
	inputName := "John"
	limit := 10
	offset := 0
	dbUsers := []domain.User{
		{ID: 1, Name: "John", PhoneNumber: "+1234567890", Country: "USA"},
		{ID: 2, Name: "Johnny", PhoneNumber: "+9876543210", Country: "Canada"},
	}

	// Mock the UserRepository to return the test users
	mockRepo := new(MockUserRepository)
	mockRepo.On("SearchByName", inputName, limit, offset).Return(dbUsers, nil)

	// Mock the Helper (ranking) to return the same users for simplicity
	mockHelper := new(MockHelper)
	mockHelper.On("IsAlphabetic", inputName).Return(true) 
	mockHelper.On("RankUsers", inputName, dbUsers).Return(dbUsers)

	// Initialize the use case with mocked dependencies
	useCase := usecase.NewSearchUseCase(mockRepo, mockHelper)

	// Call the method under test
	result, err := useCase.SearchUsersByName(inputName, limit, offset)

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, result, 2)  // Expect 2 users in the result
	assert.Equal(t, dbUsers, result) // The result should match the mock users

	// Verify that the mock methods were called with the expected arguments
	mockRepo.AssertExpectations(t)
	mockHelper.AssertExpectations(t)
}

func TestSearchUsersByName_Error(t *testing.T) {
	// Test data
	inputName := "John"
	limit := 10
	offset := 0

	// Mock the UserRepository to simulate an error
	mockRepo := new(MockUserRepository)
	mockRepo.On("SearchByName", inputName, limit, offset).Return([]domain.User{}, errors.New("database error"))

	// Mock the Helper (ranking) - we don't expect this to be called due to the error
	mockHelper := new(MockHelper)

	// Initialize the use case with mocked dependencies
	useCase := usecase.NewSearchUseCase(mockRepo, mockHelper)

	mockHelper.On("IsAlphabetic", inputName).Return(true) 
	
	// Call the method under test
	result, err := useCase.SearchUsersByName(inputName, limit, offset)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, []domain.User{}, result) // Expect an empty slice instead of nil
	mockRepo.AssertExpectations(t)
	mockHelper.AssertExpectations(t)
}

func TestSearchUsersByName_InvalidName(t *testing.T) {
	// Test data
	inputName := "12345" // Invalid name (non-alphabetic)
	limit := 10
	offset := 0

	// Mock the Helper to simulate the invalid name check
	mockHelper := new(MockHelper)
	mockHelper.On("IsAlphabetic", inputName).Return(false)

	// Initialize the use case with mocked dependencies
	useCase := usecase.NewSearchUseCase(nil, mockHelper)

	// Call the method under test
	result, err := useCase.SearchUsersByName(inputName, limit, offset)

	// Assertions
	assert.Error(t, err)                // Expecting an error due to invalid name
	assert.Equal(t, []domain.User{}, result) // Expecting an empty slice (not nil)
	mockHelper.AssertExpectations(t)
}

