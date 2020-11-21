package router

import "github.com/gin-gonic/gin"

func  InitializeRoter(middlewares ...gin.HandlerFunc) *gin.Engine {

	route := gin.Default()
	return  route
}
