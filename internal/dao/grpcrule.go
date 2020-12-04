package dao

import (
	"github.com/gin-gonic/gin"
	"go-gateway/internal/database"
	"go-gateway/internal/ent"
)

type GrpcRuleDao struct {
}

func (d *GrpcRuleDao) GetGrpcRule(c *gin.Context, id int64) (*ent.GrpcRule, error) {
	model, err := database.Client.GrpcRule.
		Get(c, id)
	return model, err
}

func (d *GrpcRuleDao) CreateGrpcRule(c *gin.Context, rule *ent.GrpcRule) error {
	_, err := database.Client.GrpcRule.Create().
		SetID(rule.ID).
		Save(c)
	return err
}
