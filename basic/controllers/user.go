package controllers

import (
	"basic/core/models"
	"basic/core/schemas"
	"basic/libs/hashing"
	"basic/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateUser(objIn *schemas.CreateUser, roles []string) (*schemas.User, *gin.Error) {
	// * get db
	db, err := services.GetDB()
	if err != nil {
		return &schemas.User{}, &gin.Error{
			Err:  err,
			Type: http.StatusInternalServerError,
			Meta: "error getting DB conn",
		}
	}

	// * check if User exists
	user, err := models.GetUserByUsername(db, objIn.Username)
	if err != nil {
		return &schemas.User{}, &gin.Error{
			Err:  err,
			Type: http.StatusInternalServerError,
			Meta: "error retrieving User record",
		}
	}
	if user != nil {
		return &schemas.User{}, &gin.Error{
			Err:  err,
			Type: http.StatusConflict,
			Meta: "User record exists with specified username",
		}
	}

	// *hash password
	hashedPwd, err := hashing.HashPassword(objIn.Password)
	if err != nil {
		return &schemas.User{}, &gin.Error{
			Err:  err,
			Type: http.StatusInternalServerError,
			Meta: "error while hashing",
		}
	}
	// * create user
	user, err = models.CreateUser(db, &models.User{
		ID:        uuid.New(),
		Username:  objIn.Username,
		Password:  hashedPwd,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return &schemas.User{}, &gin.Error{
			Err:  err,
			Type: http.StatusInternalServerError,
			Meta: "error creating User record",
		}
	}

	// * add role to user
	enfcr, err := services.GetEnforcer()
	if err != nil {
		return &schemas.User{}, &gin.Error{
			Err:  err,
			Type: http.StatusInternalServerError,
			Meta: "Something went wrong internally.",
		}
	}
	enfcr.AddRolesForUser(user.ID.String(), roles)

	return &schemas.User{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
