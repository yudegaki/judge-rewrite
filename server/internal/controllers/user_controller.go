package controllers

import (
	"github.com/gin-gonic/gin"
	"yudegaki.github.com/rewrite-judge/internal/db"
	"yudegaki.github.com/rewrite-judge/internal/entities"
	"yudegaki.github.com/rewrite-judge/internal/repositories"
	"yudegaki.github.com/rewrite-judge/internal/usecases"
)

type UserResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func convertEntityUserToControllerUser(entityUser *entities.User) *UserResponse {
	return &UserResponse{
		ID:   entityUser.GetID(),
		Name: entityUser.GetName(),
	}
}

func convertEntityUsersToControllerUsers(entityUsers []*entities.User) []*UserResponse {
	var controllerUsers []*UserResponse
	for _, entityUser := range entityUsers {
		controllerUsers = append(controllerUsers, convertEntityUserToControllerUser(entityUser))
	}
	return controllerUsers
}

func GetAllUsers(c *gin.Context) {
	repository := repositories.NewUserRepository(db.DB)
	println("ok, repo")
	usecase := usecases.NewGetUsersUsecase(repository)
	println("ok, usecase")

	result, err := usecase.Execute()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	resp := convertEntityUsersToControllerUsers(result)
	println("ok, resp")

	c.JSON(200, resp)
}
