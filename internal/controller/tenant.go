package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-gateway/internal/ent"
	"go-gateway/internal/flow/handler"
	"go-gateway/internal/middleware"
	"go-gateway/internal/service"
	"go-gateway/internal/util"
	"time"
)

type TenantList struct {
	PageSize  int `json:"page_size" form:"page_size" comment:"页数" validate:"required,min=1,max=999"`
	PageIndex int `json:"page_index" form:"page_index" comment:"页码" validate:"required,min=1,max=999"`
}

type TenantListResult struct {
	List  []*ent.Tenant `json:"list" form:"list" comment:"租户列表"`
	Count int           `json:"count" form:"count" comment:"租户总数"`
}
type StatisticsResult struct {
	Today     []int64 `json:"today" form:"today" comment:"今日统计" validate:"required"`
	Yesterday []int64 `json:"yesterday" form:"yesterday" comment:"昨日统计" validate:"required"`
}
type TenantDetail struct {
	ID int64 `json:"id" form:"id" comment:"租户ID" validate:"required"`
}

type CreateTenant struct {
	AppID    string `json:"app_id" form:"app_id" gorm:"column:app_id" comment:"租户id" validate:""`
	Name     string `json:"name" form:"name" gorm:"column:name" comment:"租户名称" validate:"required"`
	Secret   string `json:"secret" form:"secret" gorm:"column:secret" comment:"密钥" validate:"required"`
	WhiteIPS string `json:"white_ips" form:"white_ips" gorm:"column:white_ips" comment:"ip白名单，支持前缀匹配		"`
	Qpd      int    `json:"qpd" form:"qpd" gorm:"column:qpd" comment:"日请求量限制"`
	Qps      int    `json:"qps" form:"qps" gorm:"column:qps" comment:"每秒请求量限制"`
}

type UpdateTenant struct {
	ID       int64  `json:"id" form:"id" gorm:"column:id" comment:"主键ID" validate:"required"`
	AppID    string `json:"app_id" form:"app_id" gorm:"column:app_id" comment:"租户id" validate:""`
	Name     string `json:"name" form:"name" gorm:"column:name" comment:"租户名称" validate:"required"`
	Secret   string `json:"secret" form:"secret" gorm:"column:secret" comment:"密钥" validate:"required"`
	WhiteIPS string `json:"white_ips" form:"white_ips" gorm:"column:white_ips" comment:"ip白名单，支持前缀匹配		"`
	Qpd      int    `json:"qpd" form:"qpd" gorm:"column:qpd" comment:"日请求量限制"`
	Qps      int    `json:"qps" form:"qps" gorm:"column:qps" comment:"每秒请求量限制"`
}

type TenantController struct {
	ts *service.TenantService
}

func TenantRegister(router *gin.RouterGroup) {
	tenant := TenantController{}

	router.GET("/list", tenant.TenantList)
	router.GET("/detail", tenant.TenantDetail)
	router.GET("/stat", tenant.TenantStatistics)
	router.GET("/delete", tenant.TenantDelete)
	router.POST("/add", tenant.TenantAdd)
	router.POST("/update", tenant.TenantUpdate)
}

func (t *TenantController) TenantList(c *gin.Context) {
	var tenanList TenantList

	if err := c.ShouldBind(&tenanList); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	list, count, err := t.ts.TenantList(c, tenanList.PageIndex, tenanList.PageSize)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	result := TenantListResult{
		List:  list,
		Count: count,
	}
	middleware.ResponseSuccess(c, result)
	return

}
func (t *TenantController) TenantDetail(c *gin.Context) {

	var tenantDetail TenantDetail

	if err := c.ShouldBind(&tenantDetail); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	tenant, err := t.ts.GetTenant(c, tenantDetail.ID)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	middleware.ResponseSuccess(c, tenant)
}

func (t *TenantController) TenantStatistics(c *gin.Context) {
	var tenantDetail TenantDetail

	if err := c.ShouldBind(&tenantDetail); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	tenant, err := t.ts.GetTenant(c, tenantDetail.ID)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	//今日流量全天小时级访问统计
	todayStat := []int64{}

	counter, err := handler.FlowCounterHandler.GetCounter(util.FlowAppPrefix + tenant.AppID)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		c.Abort()
		return
	}
	currentTime := time.Now()
	for i := 0; i <= time.Now().In(time.Local).Hour(); i++ {
		dateTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), i, 0, 0,
			0, time.Local)
		hourData, _ := counter.GetHourData(dateTime)
		todayStat = append(todayStat, hourData)
	}

	//昨日流量全天小时级访问统计
	yesterdayStat := []int64{}
	yesterTime := currentTime.Add(-1 * time.Duration(time.Hour*24))
	for i := 0; i <= 23; i++ {
		dateTime := time.Date(yesterTime.Year(), yesterTime.Month(), yesterTime.Day(), i, 0, 0,
			0, time.Local)
		hourData, _ := counter.GetHourData(dateTime)
		yesterdayStat = append(yesterdayStat, hourData)
	}
	stat := StatisticsResult{
		Today:     todayStat,
		Yesterday: yesterdayStat,
	}
	middleware.ResponseSuccess(c, stat)
	return

}

func (t *TenantController) TenantDelete(c *gin.Context) {

	var tenantDetail TenantDetail

	if err := c.ShouldBind(&tenantDetail); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	err := t.ts.DeleteTenant(c, tenantDetail.ID)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	middleware.ResponseSuccess(c, "")
	return
}

func (t *TenantController) TenantAdd(c *gin.Context) {

	var createTenant CreateTenant
	if err := c.ShouldBind(&createTenant); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	if _, err := t.ts.FindTenant(c, createTenant.AppID); err != nil {
		middleware.ResponseError(c, 2002, errors.New("租户ID被占用，请重新输入"))
		return
	}
	if createTenant.Secret == "" {
		createTenant.Secret = util.MD5(createTenant.AppID)
	}
	model := &ent.Tenant{
		AppID:    createTenant.AppID,
		Name:     createTenant.Name,
		Secret:   createTenant.Secret,
		WhiteIps: createTenant.WhiteIPS,
		QPS:      createTenant.Qps,
		Qpd:      createTenant.Qpd,
	}
	if err := t.ts.CreateTenant(c, model); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	middleware.ResponseSuccess(c, "")
	return
}

func (t TenantController) TenantUpdate(c *gin.Context) {
	var updateTenant UpdateTenant

	if err := c.ShouldBind(&updateTenant); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	tenant, err := t.ts.GetTenant(c, updateTenant.ID)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	if updateTenant.Secret == "" {
		updateTenant.Secret = util.MD5(updateTenant.AppID)
	}
	tenant.Name = updateTenant.Name
	tenant.Secret = updateTenant.Secret
	tenant.WhiteIps = updateTenant.WhiteIPS
	tenant.QPS = updateTenant.Qps
	tenant.Qpd = updateTenant.Qpd
	if err := t.ts.UpdateTenant(c, tenant); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	middleware.ResponseSuccess(c, "")
	return
}
