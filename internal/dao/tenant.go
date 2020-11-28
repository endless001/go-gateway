package dao

import (
	"github.com/gin-gonic/gin"
	"go-gateway/internal/ent"
	"go-gateway/internal/ent/tenant"
)

func (d *Dao) FindTenant(c *gin.Context, name string) (*ent.Tenant, error) {
	model, err := d.client.Tenant.Query().
		Where(tenant.NameContains(name)).
		Only(c)
	return model, err
}

func (d *Dao) SaveTenant(c *gin.Context, model *ent.Tenant) error {
	_, err := d.client.Tenant.Create().
		SetAppID(model.AppID).
		Save(c)
	return err
}
func (d *Dao) TenantList(c *gin.Context, pageIndex int, pageSize int) ([]*ent.Tenant, int64, error) {
	offset := (pageIndex - 1) * pageSize

	appList, err := d.client.Tenant.Query().
		Limit(pageSize).
		Offset(offset).
		All(c)
	return appList, 1, err
}

func (d *Dao) GetTenant(c *gin.Context, appId string, secret string) (*ent.Tenant, error) {
	model, err := d.client.Tenant.Query().
		Where(tenant.AppIDEQ(appId), tenant.Secret(secret)).
		Only(c)
	return model, err
}
