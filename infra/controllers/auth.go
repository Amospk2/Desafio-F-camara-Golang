package controllers

import (
	"desafiot/domain/pessoa"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthController struct {
	repository pessoa.PessoaRepository
}

func (c *AuthController) Login() gin.HandlerFunc {
	return func(context *gin.Context) {
		var request AuthRequest

		if err := context.BindJSON(&request); err != nil {
			return
		}

		user, err := c.repository.GetByEmail(request.Email)

		if err != nil {
			log.Fatal(err)
			context.Status(http.StatusBadRequest)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

		if err != nil {
			context.Status(http.StatusBadRequest)
			return
		}

		accessToken, err := jwt.NewWithClaims(
			jwt.SigningMethodHS256,
			jwt.MapClaims{
				"user": user.Id,
				"exp":  time.Now().Add(time.Duration(time.Hour * 1)).Unix(),
			},
		).SignedString([]byte(os.Getenv("SECRET")))

		if err != nil {
			context.Status(http.StatusBadRequest)
			return
		}

		context.JSON(http.StatusOK, accessToken)
	}
}

func NewAuthController(r pessoa.PessoaRepository) *AuthController {
	return &AuthController{
		repository: r,
	}
}
