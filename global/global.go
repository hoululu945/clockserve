package global

import (
	"gorm.io/gorm"
)
import redis "github.com/redis/go-redis/v9"

var (
	Backend_DB    *gorm.DB
	Backend_REDIS *redis.Client
)
