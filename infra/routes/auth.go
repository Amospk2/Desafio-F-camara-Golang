package routes

import (
	"desafiot/infra/controllers"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	controller *controllers.AuthController
}

func (p *AuthRouter) Load(r *gin.Engine) {
	auth := r.Group("/auth/")
	auth.POST("/login", p.controller.Login())
}

func NewAuthRouter(
	controller *controllers.AuthController,
) *AuthRouter {
	return &AuthRouter{
		controller: controller,
	}
}
