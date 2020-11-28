package redis

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

type RedisConfig struct {
	ProxyList    []string
	Password     string
	Db           int
	ConnTimeout  int
	ReadTimeout  int
	WriteTimeout int
}

func RedisConnFactory(redisConf RedisConfig) (redis.Conn, error) {

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

func RedisLogDo(c redis.Conn, commandName string, args ...interface{}) (interface{}, error) {

	startExecTime := time.Now()
	reply, err := c.Do(commandName, args...)
	endExecTime := time.Now()

	if err != nil {
		log.WithFields(log.Fields{
			"method":    commandName,
			"err":       err,
			"bind":      args,
			"proc_time": fmt.Sprintf("%fs", endExecTime.Sub(startExecTime).Seconds()),
		}).Error("_com_redis_failure")
	} else {
		replyStr, _ := redis.String(reply, nil)
		log.WithFields(log.Fields{
			"method":    commandName,
			"bind":      args,
			"reply":     replyStr,
			"proc_time": fmt.Sprintf("%fs", endExecTime.Sub(startExecTime).Seconds()),
		}).Info("_com_redis_success")
	}
	return reply, err
}
func RedisConfDo(redisConf RedisConfig, commandName string, args ...interface{}) (interface{}, error) {

	c, err := RedisConnFactory(redisConf)
	if err != nil {
		log.WithFields(log.Fields{
			"method": commandName,
			"err":    errors.New("RedisConnFactory_error"),
			"bind":   args,
		}).Error("_com_redis_failure")
	}
	defer c.Close()
	startExecTime := time.Now()
	reply, err := c.Do(commandName, args...)
	endExecTime := time.Now()

	if err != nil {
		log.WithFields(log.Fields{
			"method":    commandName,
			"err":       err,
			"bind":      args,
			"proc_time": fmt.Sprintf("%fs", endExecTime.Sub(startExecTime).Seconds()),
		}).Error("_com_redis_failure")
	} else {
		replyStr, _ := redis.String(reply, nil)
		log.WithFields(log.Fields{
			"method":    commandName,
			"bind":      args,
			"reply":     replyStr,
			"proc_time": fmt.Sprintf("%fs", endExecTime.Sub(startExecTime).Seconds()),
		}).Info("_com_redis_success")
	}

	return reply, err
}
