package util

import (
	"github.com/garyburd/redigo/redis"
	libraryRedis "go-gateway/library/redis"
)

func RedisConfPipline(pip ...func(c redis.Conn)) error {

	redisConf := libraryRedis.RedisConfig{}
	c, err := libraryRedis.RedisConnFactory(redisConf)
	if err != nil {
		return err
	}
	defer c.Close()
	for _, f := range pip {
		f(c)
	}
	c.Flush()
	return nil
}

func RedisConfDo(commandName string, args ...interface{}) (interface{}, error) {
	redisConf := libraryRedis.RedisConfig{}
	c, err := libraryRedis.RedisConnFactory(redisConf)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	return c.Do(commandName, args...)

}
