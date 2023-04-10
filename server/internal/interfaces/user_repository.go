package interfaces

import "yudegaki.github.com/rewrite-judge/internal/entities"

type UserRepository interface {
	Get() ([]*entities.User, error)
	GetDetail(id uint) (*entities.User, error)
}
