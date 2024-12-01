package controllers

import (
	"desafiot/domain/pessoa"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type PessoaController struct {
	repository pessoa.PessoaRepository
}

func (c *PessoaController) GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := c.repository.Get()
		if err != nil {
			log.Fatal(err)
		}
		ctx.IndentedJSON(http.StatusOK, map[string]any{"users": users})
	}
}

func (c *PessoaController) GetUserById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		users, err := c.repository.GetById(id)

		if err != nil {
			ctx.Status(http.StatusNotFound)
			return
		}

		ctx.IndentedJSON(http.StatusOK, users)
	}
}

func (c *PessoaController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		if _, err := c.repository.GetById(id); err != nil {
			ctx.Status(http.StatusNotFound)
			return
		}

		if err := c.repository.Delete(id); err != nil {
			ctx.Status(http.StatusNotFound)
			return
		}

		ctx.Status(http.StatusOK)
	}
}

func (c *PessoaController) UpdateUser() gin.HandlerFunc {
	return gin.HandlerFunc(
		func(ctx *gin.Context) {
			id := ctx.Param("id")

			var userRequest pessoa.Pessoa

			user, err := c.repository.GetById(id)

			if err != nil {
				ctx.Status(http.StatusNotFound)
				return
			}

			if err := ctx.BindJSON(&userRequest); err != nil {
				ctx.Status(http.StatusUnprocessableEntity)
				return
			}

			if len(userRequest.Nome) != 0 && userRequest.Nome != "" {
				user.Nome = userRequest.Nome
			}

			if len(userRequest.Email) != 0 && userRequest.Email != "" {
				user.Email = userRequest.Email
			}

			if len(userRequest.Password) != 0 && userRequest.Password != "" {
				user.Password = userRequest.Password
			}

			if err = c.repository.Update(user); err != nil {
				ctx.Status(http.StatusBadRequest)
				return
			}

			ctx.IndentedJSON(http.StatusOK, user)
		},
	)
}

func (c *PessoaController) CreateUser() gin.HandlerFunc {
	return gin.HandlerFunc(
		func(ctx *gin.Context) {
			var pessoa pessoa.Pessoa

			if err := ctx.BindJSON(&pessoa); err != nil || !pessoa.Valid() {
				ctx.Status(http.StatusUnprocessableEntity)
				return
			}

			findpessoa, err := c.repository.GetByEmail(pessoa.Email)

			if err == nil && findpessoa.Id != "" {
				ctx.Status(http.StatusUnprocessableEntity)
				return
			}

			pessoa.Id = uuid.NewString()
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pessoa.Password), 10)

			if err != nil {
				ctx.Status(http.StatusBadRequest)
				return
			}
			pessoa.Password = string(hashedPassword)

			if err = c.repository.Create(pessoa); err != nil {
				log.Fatal(err)
			}

			if err != nil {
				ctx.Status(http.StatusInternalServerError)
				return
			}
			ctx.IndentedJSON(http.StatusCreated, pessoa)
		},
	)
}

func NewPessoaController(r pessoa.PessoaRepository) *PessoaController {
	return &PessoaController{
		repository: r,
	}
}
