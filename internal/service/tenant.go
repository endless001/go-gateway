package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-gateway/internal/dao"
	"go-gateway/internal/ent"
)

type TenantService struct {
	td *dao.TenantDao
}

func (t TenantService) TenantList(c *gin.Context, pageIndex int, pageSize int) ([]*ent.Tenant, int, error) {
	list, count, err := t.td.TenantList(c, pageIndex, pageSize)

	if err != nil {
		return nil, 0, errors.New("查询失败!")
	}

	return list, count, nil
}

func (t *TenantService) FindTenant(c *gin.Context, appId string) (*ent.Tenant, error) {
	model, err := t.td.FindTenant(c, appId)
	if err != nil {
		return nil, errors.New("查询失败!")
	}
	return model, err
}

func (t *TenantService) GetTenant(c *gin.Context, id int64) (*ent.Tenant, error) {
	model, err := t.td.GetTenant(c, id)
	if err != nil {
		return nil, errors.New("查询失败!")
	}
	return model, err
}

func (t *TenantService) DeleteTenant(c *gin.Context, id int64) error {
	err := t.td.DeleteTenant(c, id)
	if err != nil {
		return errors.New("删除失败!")
	}
	return nil
}

func (t *TenantService) CreateTenant(c *gin.Context, model *ent.Tenant) error {
	err := t.td.CreateTenant(c, model)
	if err != nil {
		return errors.New("添加失败!")
	}
	return err

}

func (t *TenantService) UpdateTenant(c *gin.Context, model *ent.Tenant) error {
	err := t.td.UpdateTenant(c, model)
	if err != nil {
		return errors.New("修改失败!")
	}
	return err

}
