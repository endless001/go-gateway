package dao

import (
	"github.com/gin-gonic/gin"
	"go-gateway/internal/ent"
	"go-gateway/internal/ent/httprule"
)

func (d *Dao) FindHttpRule(c *gin.Context, id int64) (*ent.HttpRule, error) {
	model, err := d.client.HttpRule.Query().
		Where(httprule.ID(id)).
		Only(c)
	return model, err
}

func (d *Dao) SaveHttpRule(c *gin.Context, rule ent.HttpRule) error {
	_, err := d.client.HttpRule.Create().
		SetID(rule.ID).
		Save(c)
	return err
}

func (d *Dao) GetHttpRule(c *gin.Context, serviceId int64) ([]*ent.HttpRule, int, error) {
	list, err := d.client.HttpRule.Query().
		Where(httprule.ServiceIDEQ(serviceId)).
		All(c)
	if err != nil {
		return nil, 0, err
	}
	count := len(list)
	return list, count, nil
}
