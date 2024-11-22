package usecase

import (
	"errors"
	"search-api/internal/domain"
	"search-api/internal/repository/interfaces"
	usecase "search-api/internal/usecase/interfaces"
	utils "search-api/internal/utils/interfaces"
)

type SearchUseCase struct {
	repo   interfaces.UserRepository
	ranker utils.Helper
}

func NewSearchUseCase(repo interfaces.UserRepository, ranker utils.Helper) usecase.Usecase {
	return &SearchUseCase{repo: repo, ranker: ranker}
}

func (uc *SearchUseCase) SearchUsersByName(name string, limit, offset int) ([]domain.User, error) {
	ok := uc.ranker.IsAlphabetic(name)
	if !ok {
		return []domain.User{}, errors.New("name is not valid")
	}
	dbData, err := uc.repo.SearchByName(name, limit, offset)
	if err != nil {
		return []domain.User{}, err
	}

	result := uc.ranker.RankUsers(name, dbData)

	return result, nil
}
