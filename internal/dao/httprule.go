package dao

import (
	"github.com/gin-gonic/gin"
	"go-gateway/internal/database"
	"go-gateway/internal/ent"
	"go-gateway/internal/ent/httprule"
)

type HttpRuleDao struct {
}

func (d *HttpRuleDao) FindHttpRule(c *gin.Context, id int64) (*ent.HttpRule, error) {
	model, err := database.Client.HttpRule.Query().
		Where(httprule.ID(id)).
		Only(c)
	return model, err
}

func (d *HttpRuleDao) SaveHttpRule(c *gin.Context, rule ent.HttpRule) error {
	_, err := database.Client.HttpRule.Create().
		SetID(rule.ID).
		Save(c)
	return err
}

func (d *HttpRuleDao) GetHttpRule(c *gin.Context, serviceId int64) ([]*ent.HttpRule, int, error) {
	list, err := database.Client.HttpRule.Query().
		Where(httprule.ServiceIDEQ(serviceId)).
		All(c)
	if err != nil {
		return nil, 0, err
	}
	count := len(list)
	return list, count, nil
}
