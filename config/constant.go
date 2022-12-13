package config

import "time"

const (
	AppName = "shop-search-api"

	HeaderAuthField     = "Authorization"
	HeaderAuthDateField = "Authorization-Date"
	AuthorizationExpire = time.Minute * 30

	RunModeDev  = "dev"
	RunModeProd = "prod"

	DefaultMySQLClient = "default-mysql"
	DefaultRedisClient = "default-redis"

	RedisKeyPrefixSignature       = "sign:"
	RedisSignatureCacheSeconds    = 300 * time.Second
	HeaderSignTokenTimeoutSeconds = 300 * time.Second
)
