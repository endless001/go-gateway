package dao

import (
	"github.com/gin-gonic/gin"
	"go-gateway/internal/ent"
	"go-gateway/internal/ent/accesscontrol"
)

func (d *Dao) FindAccessControl(c *gin.Context, serviceId int64) (*ent.AccessControl, error) {
	model, err := d.client.AccessControl.Query().
		Where(accesscontrol.ServiceIDEQ(serviceId)).
		Only(c)

	return model, err
}

func (d *Dao) SaveAccessControl(c *gin.Context, model ent.AccessControl) error {
	_, err := d.client.AccessControl.Create().
		SetOpenAuth(model.OpenAuth).
		Save(c)
	return err
}

func (d *Dao) GetAccessControl(c *gin.Context, serviceId int64) ([]*ent.AccessControl, int, error) {
	list, err := d.client.AccessControl.Query().
		Where(accesscontrol.ServiceIDEQ(serviceId)).
		All(c)
	if err != nil {
		return nil, 0, err
	}
	count := len(list)
	return list, count, nil
}
