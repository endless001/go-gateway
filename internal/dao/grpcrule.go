package dao

import (
	"github.com/gin-gonic/gin"
	"go-gateway/internal/database"
	"go-gateway/internal/ent"
	"go-gateway/internal/ent/grpcrule"
)

type GrpcRuleDao struct {
}

func (d *GrpcRuleDao) GetGrpcRule(c *gin.Context, id int64) (*ent.GrpcRule, error) {
	model, err := database.Client.GrpcRule.
		Get(c, id)
	return model, err
}

func (d *GrpcRuleDao) SaveGrpcRule(c *gin.Context, rule ent.GrpcRule) error {
	_, err := database.Client.GrpcRule.Create().
		SetID(rule.ID).
		Save(c)
	return err
}

func (d *GrpcRuleDao) FindGrpcRule(c *gin.Context, serviceId int64) ([]*ent.GrpcRule, int, error) {
	list, err := database.Client.GrpcRule.Query().
		Where(grpcrule.ServiceIDEQ(serviceId)).
		All(c)
	if err != nil {
		return nil, 0, err
	}
	count := len(list)
	return list, count, nil
}
