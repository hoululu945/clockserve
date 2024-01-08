package global

import (
	"gorm.io/gorm"
	"serve/config"
)
import redis "github.com/redis/go-redis/v9"

var (
	Backend_DB    *gorm.DB
	Backend_REDIS *redis.Client
	GVA_CONFIG    config.Server
)
