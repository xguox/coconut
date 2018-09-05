package serializer

import (
	"coconut/model"
	"coconut/util"

	"github.com/gin-gonic/gin"
)

type UserSerializer struct {
	C *gin.Context
}

type UserResponse struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Bio      string  `json:"bio"`
	Image    *string `json:"image"`
	Token    string  `json:"token"`
}

func (us *UserSerializer) Response() UserResponse {
	currentUser := us.C.MustGet("current_user").(model.User)

	user := UserResponse{
		Username: currentUser.Username,
		Email:    currentUser.Email,
		Token:    util.GenerateToken(currentUser.ID),
	}
	return user
}
