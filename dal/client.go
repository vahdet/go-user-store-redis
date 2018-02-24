package dal

import (
	"fmt"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	redisUrlKey = "USER_STORE_REDIS_URL"
)

var client *redis.Client

// Initializes the data store client (i.e. Redis)
func InitDataStoreClient() {
	redisUrl := os.Getenv(redisUrlKey)
	log.WithFields(log.Fields{
		"url": redisUrl,
	}).Info(fmt.Sprintf("redis client starting..."))
	client = redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: "", //TODO: Set Password
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.WithFields(log.Fields{
			"url": redisUrl,
		}).Error(fmt.Sprintf("redis dial failed!\nPong: '%v' \nError: '%v'", pong, err))
		return
	} else {
		log.Debug(pong)
	}
	return
}

// Ends the data store client (i.e. Redis)
func CloseDataStoreClient() {
	log.Info(fmt.Sprintf("redis client closing..."))
	client.Close()
}
