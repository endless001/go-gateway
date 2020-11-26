package dao

import (
	"context"
	"go-gateway/conf"
	"go-gateway/ent"
)

type Dao struct {
	c      *conf.Config
	client *ent.Client
}

func New(c *conf.Config) (d *Dao) {
	client, _ := ent.Open(c.DataBase.DriverName, c.DataBase.DataSourceName)
	d = &Dao{
		c:      c,
		client: client,
	}
	return
}

func (d *Dao) Close() {
	if d.client != nil {
		d.client.Close()
	}
}

// Ping dao ping
func (d *Dao) Ping(c context.Context) (err error) {

	return
}
