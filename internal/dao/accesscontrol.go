package dao

import (
	"github.com/gin-gonic/gin"
	"go-gateway/internal/database"
	"go-gateway/internal/ent"
)

type AccessControlDao struct {
}

func (d *AccessControlDao) GetAccessControl(c *gin.Context, id int64) (*ent.AccessControl, error) {
	model, err := database.Client.AccessControl.
		Get(c, id)

	return model, err
}

func (d *AccessControlDao) CreateAccessControl(c *gin.Context, model *ent.AccessControl) error {
	_, err := database.Client.AccessControl.Create().
		SetOpenAuth(model.OpenAuth).
		Save(c)
	return err
}
