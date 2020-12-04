package controller

import (
	"github.com/gin-gonic/gin"
	"go-gateway/internal/service"
)

type ServiceInfoController struct {
	sv *service.ServiceInfoService
}

func ServiceRegister(group *gin.RouterGroup) {
	service := &ServiceInfoController{}
	group.GET("/list", service.ServiceList)
	group.GET("/delete", service.ServiceDelete)
	group.GET("/detail", service.ServiceDetail)
	group.GET("/stat", service.ServiceStat)
	group.POST("/add/http", service.ServiceAddHttp)
	group.POST("/update/http", service.ServiceUpdateHttp)

	group.POST("/add/tcp", service.ServiceAddTcp)
	group.POST("/update/tcp", service.ServiceUpdateTcp)
	group.POST("/add/grpc", service.ServiceAddGrpc)
	group.POST("/update/grpc", service.ServiceUpdateGrpc)
}

func (s *ServiceInfoController) ServiceList(c *gin.Context) {
}

func (s *ServiceInfoController) ServiceDelete(c *gin.Context) {
}

func (s *ServiceInfoController) ServiceDetail(c *gin.Context) {
}

func (s *ServiceInfoController) ServiceStat(c *gin.Context) {
}

func (s *ServiceInfoController) ServiceAddHttp(c *gin.Context) {
}

func (s *ServiceInfoController) ServiceUpdateHttp(c *gin.Context) {
}

func (s *ServiceInfoController) ServiceAddTcp(c *gin.Context) {
}

func (s *ServiceInfoController) ServiceUpdateTcp(c *gin.Context) {
}

func (s *ServiceInfoController) ServiceAddGrpc(c *gin.Context) {
}

func (s *ServiceInfoController) ServiceUpdateGrpc(c *gin.Context) {
}
