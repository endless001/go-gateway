package controller

import (
	"github.com/gin-gonic/gin"
	"go-gateway/internal/service"
)

type DashboardController struct {
	sv *service.ServiceInfoService
}

func DashboardRegister(group *gin.RouterGroup) {
	dashboard := &DashboardController{}

	group.GET("/panel/groupdata", dashboard.PanelGroupData)
	group.GET("/flowstat", dashboard.FlowStat)
	group.GET("/servicestat", dashboard.ServiceStat)
}

func (d *DashboardController) PanelGroupData(c *gin.Context) {
}

func (d *DashboardController) ServiceStat(c *gin.Context) {

}
func (service *DashboardController) FlowStat(c *gin.Context) {

}
