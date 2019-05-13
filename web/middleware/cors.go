package middleware

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/context"
)
func CorsMiddleware()context.Handler  {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},   //允许通过的主机名称
		AllowCredentials: true,
		AllowedHeaders:[]string{"*"},
	})
}

