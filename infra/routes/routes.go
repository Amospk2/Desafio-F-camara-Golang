package routes

import (
	"desafiot/infra/controllers"
	"desafiot/infra/database"
	"github.com/gin-gonic/gin"
)

func CreateRoute(gin *gin.Engine) {
	pool := database.NewConnect("postgres://postgres:1234567890@172.17.0.1:5432/postgres")
	NewUserRouter(controllers.NewPessoaController(database.NewPessoaRepositoryImp(pool))).Load(gin)
	NewAuthRouter(controllers.NewAuthController(database.NewPessoaRepositoryImp(pool))).Load(gin)
}
