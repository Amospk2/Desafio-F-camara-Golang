package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	User string `json:"user"`
	Exp  string `json:"exp"`
}

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		header := strings.Split(ctx.GetHeader("Authorization"), " ")

		if len(header) == 0 || !strings.Contains(header[0], "Bearer") {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		accessToken := header[1]
		if accessToken == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		parsedAccessToken, _ := jwt.ParseWithClaims(
			accessToken, jwt.MapClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("SECRET")), nil
			},
		)

		if !parsedAccessToken.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("user", parsedAccessToken.Claims)

		ctx.Next()
	}
}
