package someframes

import (
	"context"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

// https://redis.uptrace.dev/guide/go-redis.html#connecting-to-redis-server

func RedisServerOpr(getkeyname string) {
	// Every Redis command accepts a context that you can use to set timeouts or
	// propagate some information, for example, tracing context.
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if val, err := rdb.Get(ctx, getkeyname).Result(); err == nil {
		// normal process for val
		log.WithFields(log.Fields{"key": getkeyname}).Info("Redis get result:", val)
	} else if err == redis.Nil {
		log.WithFields(log.Fields{"key": getkeyname}).Warn("Redis reply - key does not exist!")
	} else {
		log.Error(err)
	}

	log.Info("End.")
}
