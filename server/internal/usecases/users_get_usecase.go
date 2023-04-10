package usecases

import (
	"yudegaki.github.com/rewrite-judge/internal/entities"
	"yudegaki.github.com/rewrite-judge/internal/interfaces"
)

type GetUsersUsecase struct {
	repository interfaces.UserRepository
}

func NewGetUsersUsecase(r interfaces.UserRepository) *GetUsersUsecase {
	return &GetUsersUsecase{
		repository: r,
	}
}

func (u *GetUsersUsecase) Execute() ([]*entities.User, error) {
	return u.repository.Get()
}
