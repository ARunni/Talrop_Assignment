package interfaces

import "search-api/internal/domain"

type UserRepository interface {
	SearchByName(name string, limit, offset int) ([]domain.User, error)
}
