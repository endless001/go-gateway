package dao

import (
	"github.com/gin-gonic/gin"
	"go-gateway/internal/database"
	"go-gateway/internal/ent"
	"go-gateway/internal/ent/tenant"
)

type TenantDao struct {
}

func (d *TenantDao) FindTenant(c *gin.Context, appId string) (*ent.Tenant, error) {
	model, err := database.Client.Tenant.
		Query().
		Where(tenant.AppIDEQ(appId)).
		Only(c)
	return model, err
}

func (d *TenantDao) GetTenant(c *gin.Context, id int64) (*ent.Tenant, error) {
	model, err := database.Client.Tenant.
		Get(c, id)
	return model, err
}

func (d *TenantDao) CreateTenant(c *gin.Context, model *ent.Tenant) error {
	_, err := database.Client.Tenant.Create().
		SetAppID(model.AppID).
		Save(c)
	return err

}

func (d *TenantDao) UpdateTenant(c *gin.Context, model *ent.Tenant) error {
	_, err := database.Client.Tenant.
		UpdateOne(model).
		Save(c)
	return err

}
func (d *TenantDao) TenantList(c *gin.Context, pageIndex int, pageSize int) ([]*ent.Tenant, int, error) {
	offset := (pageIndex - 1) * pageSize

	list, err := database.Client.Tenant.Query().
		Limit(pageSize).
		Offset(offset).
		All(c)

	count, _ := database.Client.Tenant.
		Query().
		Count(c)

	return list, count, err
}

func (d *TenantDao) DeleteTenant(c *gin.Context, id int64) error {
	err := database.Client.Tenant.
		DeleteOneID(id).Exec(c)
	return err
}
