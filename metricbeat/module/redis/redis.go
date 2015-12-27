package redis

import (
	//"fmt"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/metricbeat/helper"
	"github.com/garyburd/redigo/redis"
)

func init() {
	Module.Register()
}

var Module = helper.NewModule("redis", Redis{})

var Config = &RedisModuleConfig{}

type RedisModuleConfig struct {
	Metrics map[string]interface{}
	Hosts   []string
}

type Redis struct {
	Name   string
	Config RedisModuleConfig
}

func (e Redis) Setup() {

	// Loads module config
	// This is module specific config object
	Module.LoadConfig(&Config)
}

func Connect() redis.Conn {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		logp.Err("Redis connection error: %v", err)
	}

	//defer conn.Close()
	return conn
}
