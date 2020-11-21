package dao

import (
	"context"
	"go-gateway/ent"
)

type Dao struct {
	c string
	client *ent.Client
}
func New(c string) (d *Dao) {
	client,_:=ent.Open("mysql", c)
	d = &Dao{
		c: c,
		client:client,
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