package data_source

import (
	"github.com/go-redis/redis"
)

var (
	Client *redis.Client
	Nil    = redis.Nil
)

type SliceCmd = redis.SliceCmd
type StringStringMapCmd = redis.StringStringMapCmd

// Init 初始化连接
func Init() (err error) {
	Client = redis.NewClient(&redis.Options{
		Addr:     "192.168.10.3:6379", // no password set
		DB:       0,                   // use default DB
		PoolSize: 100,
	})

	_, err = Client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = Client.Close()
}
