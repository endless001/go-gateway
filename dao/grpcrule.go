package dao

import (
	"github.com/gin-gonic/gin"
	"go-gateway/ent"
	"go-gateway/ent/grpcrule"
)

func (d *Dao) FindGrpcRule(c *gin.Context, id int64) (*ent.GrpcRule, error) {
	model, err := d.client.GrpcRule.Query().
		Where(grpcrule.ID(id)).
		Only(c)
	return model, err
}

func (d *Dao) SaveGrpcRule(c *gin.Context, rule ent.GrpcRule) error {
	_, err := d.client.GrpcRule.Create().
		SetID(rule.ID).
		Save(c)
	return err
}

func (d *Dao) GetGrpcRule(c *gin.Context, serviceId int64) ([]*ent.GrpcRule, int, error) {
	list, err := d.client.GrpcRule.Query().
		Where(grpcrule.ServiceIDEQ(serviceId)).
		All(c)
	if err != nil {
		return nil, 0, err
	}
	count := len(list)
	return list, count, nil
}
