package svctx

import (
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/blog-gin/config"
	"github.com/ve-weiyi/blog-gin/infra/middleware"
	"github.com/ve-weiyi/blog-gin/infra/middleware/permissionx"
	"github.com/ve-weiyi/blog-gin/infra/tokenx"
	"github.com/ve-weiyi/blog-gin/initialize"
)

const (
	// tokenKeyPrefix 令牌在 Redis 中的 key 前缀，与 blog-cloud 的 cachekey.TokenStorePrefixApp 一致
	tokenKeyPrefix = "blog:app:token"
	// tokenIssuer 令牌签发者标识
	tokenIssuer = "blog-gin"
)

type ServiceContext struct {
	Config       *config.Config
	DB           *gorm.DB
	Redis        *redis.Client
	RabbitMQ     *amqp.Connection
	TokenManager tokenx.Manager

	UserAuth     gin.HandlerFunc
	AdminAuth    gin.HandlerFunc
	RateLimit    gin.HandlerFunc
	Permission   gin.HandlerFunc
	OperationLog gin.HandlerFunc
}

func NewServiceContext(cfg *config.Config) *ServiceContext {
	rds := initialize.Redis(cfg)

	// TokenManager: Redis 存储的 JWT Token，支持吊销和单点登录
	tokenManager, err := tokenx.New(tokenx.Config{
		Signer:          tokenx.NewJWTSigner([]byte(cfg.JWT.AccessSecret), tokenIssuer),
		Store:           tokenx.NewRedisStore(rds),
		KeyPrefix:       tokenKeyPrefix,
		AccessTokenTTL:  time.Duration(cfg.JWT.AccessExpire) * time.Second,
		RefreshTokenTTL: time.Duration(cfg.JWT.RefreshExpire) * time.Second,
	})
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:       cfg,
		DB:           initialize.Gorm(cfg),
		Redis:        rds,
		RabbitMQ:     initialize.RabbitMQ(cfg),
		TokenManager: tokenManager,
		UserAuth:     middleware.UserAuth(tokenManager),
		AdminAuth:    middleware.AdminAuth(tokenManager),
		RateLimit:    middleware.Limit(rds),
		Permission:   middleware.Permission(permissionx.NewNoopEnforcer()),
		OperationLog: middleware.OperationLog(),
	}
}
