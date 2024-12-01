package routes

import (
	"desafiot/infra/controllers"
	"desafiot/infra/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	controller *controllers.PessoaController
}

func (p *UserRouter) Load(gin *gin.Engine) {
	user := gin.Group("/users/")

	user.GET("/", p.controller.GetUsers())
	user.GET("/:id", p.controller.GetUserById())
	user.PUT("/:id", p.controller.UpdateUser())
	user.DELETE("/:id", p.controller.Delete(), middleware.AuthenticationMiddleware())
	user.POST("/", p.controller.CreateUser())
}

func NewUserRouter(
	controller *controllers.PessoaController,
) *UserRouter {
	return &UserRouter{
		controller: controller,
	}
}
