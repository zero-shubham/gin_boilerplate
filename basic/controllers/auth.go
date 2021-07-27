package controllers

import (
	"basic/core/models"
	"basic/core/schemas"
	basicjwt "basic/libs/basic_jwt"
	"basic/libs/hashing"
	"basic/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(loginCreds *schemas.AuthPost) (*schemas.AuthPostResp, *gin.Error) {
	// * get db
	db, err := services.GetDB()
	if err != nil {
		return &schemas.AuthPostResp{}, &gin.Error{
			Err:  err,
			Type: http.StatusInternalServerError,
			Meta: "error getting DB conn",
		}
	}

	user, err := models.GetUserByUsername(db, loginCreds.Username)
	if err != nil {
		return &schemas.AuthPostResp{}, &gin.Error{
			Err:  err,
			Type: http.StatusInternalServerError,
			Meta: "error retrieving User record",
		}
	}
	if user == nil {
		return &schemas.AuthPostResp{}, &gin.Error{
			Err:  err,
			Type: http.StatusUnauthorized,
			Meta: "Credentials combination does not match",
		}
	}

	match := hashing.CheckPasswordHash(loginCreds.Password, user.Password)
	if !match {
		return &schemas.AuthPostResp{}, &gin.Error{
			Err:  err,
			Type: http.StatusUnauthorized,
			Meta: "Credentials combination does not match",
		}
	}

	token, err := basicjwt.CreateToken(
		user.ID,
		60,
		"auth",
	)
	if err != nil {
		return &schemas.AuthPostResp{}, &gin.Error{
			Err:  err,
			Type: http.StatusInternalServerError,
			Meta: "error retrieving token",
		}
	}

	return &schemas.AuthPostResp{
		Token: token,
	}, nil

}
