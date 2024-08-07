package usecase

import (
	"backend/entities"
	"backend/repository"
)

type ApiGameUsecase struct {
}

func (u *ApiGameUsecase) QueryByGame(game string) ([]*entities.Translate, error) {
	return repository.QueryAllTranslateByGame(game)
}

func (u *ApiGameUsecase) InsertOrUpdateTranslate(game, rawWord, translate string) error {
	return repository.InsertOrUpdateTranslate(game, rawWord, translate)
}

func (u *ApiGameUsecase) DeleteTranslate(game, rawWord string) error {
	return repository.DeleteTranslate(game, rawWord)
}
