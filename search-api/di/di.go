package di

import (
	"database/sql"
	handler "search-api/internal/handler"
	"search-api/internal/repository"
	"search-api/internal/usecase"
	"search-api/internal/utils"
)

// Initialize repository, use case, and handler
func Initialize(db *sql.DB) *handler.SearchHandler {
	userRepo := repository.NewUserRepository(db)
	helper := utils.NewHelper()
	searchUseCase := usecase.NewSearchUseCase(userRepo, helper)
	searchHandler := handler.NewSearchHandler(searchUseCase)
	return searchHandler
}
