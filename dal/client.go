package dal

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/vahdet/tafalk-user-store/app"
	log "github.com/sirupsen/logrus"
)

var client *redis.Client

func InitDataStoreClient() {
	redisUrl := app.Config.DataStoreUrl

	client = redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.WithFields(log.Fields{
			"url": redisUrl,
		}).Error(fmt.Sprintf("redis dial failed!\nPong: '%v' \nError: '%v'", pong, err))
		return
	}
	fmt.Println(pong)
	return
}

func CloseDataStoreClient() {
	client.Close()
}
