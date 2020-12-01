package dao

import (
	"github.com/gin-gonic/gin"
	"go-gateway/internal/database"
	"go-gateway/internal/ent"
	"go-gateway/internal/ent/tenant"
)

type TenantDao struct {
}

func (d *TenantDao) FindTenant(c *gin.Context, name string) (*ent.Tenant, error) {
	model, err := database.Client.Tenant.Query().
		Where(tenant.NameContains(name)).
		Only(c)
	return model, err
}

func (d *TenantDao) SaveTenant(c *gin.Context, model *ent.Tenant) error {
	_, err := database.Client.Tenant.Create().
		SetAppID(model.AppID).
		Save(c)
	return err
}
func (d *TenantDao) TenantList(c *gin.Context, pageIndex int, pageSize int) ([]*ent.Tenant, int64, error) {
	offset := (pageIndex - 1) * pageSize

	appList, err := database.Client.Tenant.Query().
		Limit(pageSize).
		Offset(offset).
		All(c)
	return appList, 1, err
}

func (d *TenantDao) GetTenant(c *gin.Context, appId string, secret string) (*ent.Tenant, error) {
	model, err := database.Client.Tenant.Query().
		Where(tenant.AppIDEQ(appId), tenant.Secret(secret)).
		Only(c)
	return model, err
}
