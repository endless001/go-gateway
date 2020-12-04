package dao

import (
	"github.com/gin-gonic/gin"
	"go-gateway/internal/database"
	"go-gateway/internal/ent"
)

type ServiceInfoDao struct {
}

func (d *ServiceInfoDao) GetServiceInfo(c *gin.Context, id int64) (*ent.ServiceInfo, error) {
	model, err := database.Client.ServiceInfo.
		Get(c, id)
	return model, err
}

func (d *ServiceInfoDao) CreateServiceInfo(c *gin.Context, model *ent.ServiceInfo) error {
	_, err := database.Client.ServiceInfo.Create().
		SetID(model.ID).
		Save(c)
	return err
}

func (d *ServiceInfoDao) UpdateServiceInfo(c *gin.Context, model *ent.ServiceInfo) error {
	err := database.Client.ServiceInfo.
		UpdateOne(model).
		Exec(c)

	return err
}

func (d *ServiceInfoDao) DeleteServiceInfo(c *gin.Context, id int64) error {
	err := database.Client.ServiceInfo.
		DeleteOneID(id).Exec(c)
	return err
}
