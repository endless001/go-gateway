package controller

import (
	"github.com/gin-gonic/gin"
	"go-gateway/internal/service"
)

type ServiceInfoController struct {
	sv *service.ServiceInfoService
}

func ServiceRegister(group *gin.RouterGroup) {

}
