package dao

import (
	"github.com/gin-gonic/gin"
	"go-gateway/ent"
	"go-gateway/ent/accesscontrol"
)

func (d *Dao) FindAccessControl(c *gin.Context, serviceId int64) (*ent.AccessControl, error) {
	model, err := d.client.AccessControl.Query().
		Where(accesscontrol.ServiceIDEQ(serviceId)).
		Only(c)

	return model, err
}
