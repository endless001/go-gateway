package util

import (
	"github.com/garyburd/redigo/redis"
)

func RedisConfPipline(pip ...func(c redis.Conn)) error {

	//c,err:=lib.RedisConnFactory("default")
	return nil
}
