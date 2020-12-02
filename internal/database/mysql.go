package database

import (
	_ "github.com/go-sql-driver/mysql"
	"go-gateway/internal/conf"
	"go-gateway/internal/ent"
)

var Client *ent.Client

func New(config *conf.Config) {

	Client, _ = ent.Open(config.DataBase.DriverName, config.DataBase.DataSourceName)

}
