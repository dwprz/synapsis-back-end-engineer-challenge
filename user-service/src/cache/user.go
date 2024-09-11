package cache

import (
	"user-service/src/common/log"
	"user-service/src/interface/cache"
	"user-service/src/model/entity"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type UserImpl struct {
	redis  *redis.ClusterClient
}

func NewUser(r *redis.ClusterClient) cache.User {
	return &UserImpl{
		redis:  r,
	}
}

func (c *UserImpl) Cache(ctx context.Context, user *entity.User) {
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "cache.UserImpl/Cache", "section": "json.Marshal"}).Error(err)
		return
	}

	key := fmt.Sprintf("user:%s", user.Email)
	const expire = 24 * time.Hour

	if _, err := c.redis.SetEx(ctx, key, string(jsonData), expire).Result(); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "cache.UserImpl/Cache", "section": "redis.SetEx"}).Error(err)
	}
}

func (c *UserImpl) FindByEmail(ctx context.Context, email string) *entity.User {
	res, err := c.redis.Get(ctx, fmt.Sprintf("user:%s", email)).Result()

	if err != nil && err != redis.Nil {
		log.Logger.WithFields(logrus.Fields{"location": "cache.UserImpl/FindByEmail", "section": "redis.Get"}).Error(err)
		return nil
	}

	if res == "" {
		return nil
	}

	user := new(entity.User)

	if err := json.Unmarshal([]byte(res), user); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "cache.UserImpl/FindByEmail", "section": "json.Unmarshal"}).Error(err)
		return nil
	}

	return user
}

func (c *UserImpl) DeleteByEmail(ctx context.Context, email string) {
	key := fmt.Sprintf("user:%s", email)

	if _, err := c.redis.Del(ctx, key).Result(); err != nil && err != redis.Nil {
		log.Logger.WithFields(logrus.Fields{"location": "cache.UserImpl/DeleteByEmail", "section": "redis.Del"}).Error(err)
	}
}
