package interfaces

import "search-api/internal/domain"

type Helper interface {
	RankUsers(inputName string, users []domain.User) []domain.User
	IsAlphabetic(input string) bool
}
