package router

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-gateway/internal/controller"
	"go-gateway/internal/middleware"
	"log"
)

func RegisterRoutes(middlewares ...gin.HandlerFunc) *gin.Engine {

	router := gin.Default()

	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	store, err := sessions.NewRedisStore(10, "tcp", "47.103.25.179:6379",
		"", []byte("secret"))

	if err != nil {
		log.Fatalf("sessions.NewRedisStore err:%v", err)
	}

	userRouter := router.Group("/user")

	userRouter.Use(sessions.Sessions("session", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.TranslationMiddleware())
	{
		controller.UserRegister(userRouter)
	}

	return router

}
