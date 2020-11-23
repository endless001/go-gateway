package dao

import (
	"github.com/gin-gonic/gin"
	"go-gateway/ent"
	"go-gateway/ent/app"
)

func (d *Dao) FindApp(c *gin.Context, name string) (*ent.App, error) {
	model, err := d.client.App.Query().
		Where(app.NameContains(name)).
		Only(c)
	return model, err
}

func (d *Dao) SaveApp(c *gin.Context, model *ent.App) error {
	_, err := d.client.App.Create().
		SetAppID(model.AppID).
		Save(c)
	return err
}
func (d *Dao) AppList(c *gin.Context, pageIndex int, pageSize int) ([]*ent.App, int64, error) {
	offset := (pageIndex - 1) * pageSize

	appList, err := d.client.App.Query().
		Limit(pageSize).
		Offset(offset).
		All(c)
	return appList, 1, err
}
