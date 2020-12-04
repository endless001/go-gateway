package dao

import (
	"github.com/gin-gonic/gin"
	"go-gateway/internal/database"
	"go-gateway/internal/ent"
)

type HttpRuleDao struct {
}

func (d *HttpRuleDao) GetHttpRule(c *gin.Context, id int64) (*ent.HttpRule, error) {
	model, err := database.Client.HttpRule.
		Get(c, id)
	return model, err
}

func (d *HttpRuleDao) CreateHttpRule(c *gin.Context, rule *ent.HttpRule) error {
	_, err := database.Client.HttpRule.Create().
		SetID(rule.ID).
		Save(c)
	return err
}
