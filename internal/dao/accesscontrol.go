package dao

import (
	"github.com/gin-gonic/gin"
	"go-gateway/internal/database"
	"go-gateway/internal/ent"
	"go-gateway/internal/ent/accesscontrol"
)

type AccessControlDao struct {
}

func (d *AccessControlDao) GetAccessControl(c *gin.Context, id int64) (*ent.AccessControl, error) {
	model, err := database.Client.AccessControl.
		Get(c, id)

	return model, err
}

func (d *AccessControlDao) SaveAccessControl(c *gin.Context, model ent.AccessControl) error {
	_, err := database.Client.AccessControl.Create().
		SetOpenAuth(model.OpenAuth).
		Save(c)
	return err
}

func (d *AccessControlDao) FindAccessControl(c *gin.Context, serviceId int64) ([]*ent.AccessControl, int, error) {
	list, err := database.Client.AccessControl.Query().
		Where(accesscontrol.ServiceIDEQ(serviceId)).
		All(c)
	if err != nil {
		return nil, 0, err
	}
	count := len(list)
	return list, count, nil
}
