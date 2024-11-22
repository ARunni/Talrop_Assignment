package interfaces

import "search-api/internal/domain"

type Usecase interface {
	SearchUsersByName(name string, limit, offset int) ([]domain.User, error)
}
