package redis

import (
	"github.com/garyburd/redigo/redis"
	"math/rand"
	"time"
)

type Redis struct {
	ProxyList    []string
	Password     string
	Db           int
	ConnTimeout  int
	ReadTimeout  int
	WriteTimeout int
}

func RedisConnFactory(redisConf Redis) (redis.Conn, error) {

	randHost := redisConf.ProxyList[rand.Intn(len(redisConf.ProxyList))]

	if redisConf.ConnTimeout == 0 {
		redisConf.ConnTimeout = 50
	}
	if redisConf.ReadTimeout == 0 {
		redisConf.ReadTimeout = 100
	}
	if redisConf.WriteTimeout == 0 {
		redisConf.WriteTimeout = 100
	}

	c, err := redis.Dial(
		"tcp",
		randHost,
		redis.DialConnectTimeout(time.Duration(redisConf.ConnTimeout)*time.Millisecond),
		redis.DialReadTimeout(time.Duration(redisConf.ReadTimeout)*time.Millisecond),
		redis.DialWriteTimeout(time.Duration(redisConf.WriteTimeout)*time.Millisecond))

	if err != nil {
		return nil, err
	}

	if redisConf.Password != "" {
		if _, err := c.Do("AUTH", redisConf.Password); err != nil {
			c.Close()
			return nil, err
		}
	}
	if redisConf.Db != 0 {
		if _, err := c.Do("SELECT", redisConf.Db); err != nil {
			c.Close()
			return nil, err
		}
	}
	return c, nil

}
