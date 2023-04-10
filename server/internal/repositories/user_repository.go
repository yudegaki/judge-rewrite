package repositories

import (
	"gorm.io/gorm"
	"yudegaki.github.com/rewrite-judge/internal/entities"
)

type UserRepository struct {
	Conn *gorm.DB
}

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50);not null"`
	Password string `gorm:"type:varchar(100);not null"`
}

func NewUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{Conn: conn}
}

func convertRepositoryUserToEntityUser(user *User) *entities.User {
	var entityUser entities.User
	entityUser.SetID(user.ID)
	entityUser.SetName(user.Name)

	return &entityUser
}

func convertRepositoryUsersToEntityUsers(users []*User) []*entities.User {
	var entityUsers []*entities.User
	for _, user := range users {
		entityUsers = append(entityUsers, convertRepositoryUserToEntityUser(user))
	}
	return entityUsers
}

func (r *UserRepository) Get() ([]*entities.User, error) {
	var dbUsers []*User
	if err := r.Conn.Find(&dbUsers).Error; err != nil {
		return nil, err
	}
	return convertRepositoryUsersToEntityUsers(dbUsers), nil
}

func (r *UserRepository) GetDetail(id uint) (*entities.User, error) {
	var dbUser User
	if err := r.Conn.First(&dbUser, id).Error; err != nil {
		return nil, err
	}
	return convertRepositoryUserToEntityUser(&dbUser), nil
}
