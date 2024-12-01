package main

import (
	"github.com/gin-gonic/gin"
	"desafiot/infra/routes"
)

func main() {
	r := gin.Default()
	routes.CreateRoute(r)
	r.Run(":8000")
}
