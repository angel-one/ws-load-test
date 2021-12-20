package middlewares

import (
	"github.com/angel-one/go-utils/handlers"
	"github.com/gin-gonic/gin"
)

// JWT is used to get the jwt middleware
func JWT(options handlers.JWTOptions) (gin.HandlerFunc, error) {
	j, err := handlers.NewJWTHandler(options)
	if err != nil {
		return nil, err
	}
	return j.Middleware(), nil
}
