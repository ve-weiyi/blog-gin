package svctx

import (
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/ve-weiyi/vkit/adapter/storex"
	"github.com/ve-weiyi/vkit/adapter/storex/tokenstore"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-gin/config"
	"github.com/ve-weiyi/ve-blog-gin/infra/middleware"
	"github.com/ve-weiyi/ve-blog-gin/infra/middleware/permissionx"
	"github.com/ve-weiyi/ve-blog-gin/initialize"
)

type ServiceContext struct {
	Config     *config.Config
	DB         *gorm.DB
	Redis      *redis.Client
	RabbitMQ   *amqp.Connection
	TokenStore tokenstore.TokenStore

	UserAuth     gin.HandlerFunc
	AdminAuth    gin.HandlerFunc
	RateLimit    gin.HandlerFunc
	Permission   gin.HandlerFunc
	OperationLog gin.HandlerFunc
}

func NewServiceContext(cfg *config.Config) *ServiceContext {
	rds := initialize.Redis(cfg)

	// TokenStore: Redis 存储的 JWT Token，支持吊销和单设备登录
	tokenStore := tokenstore.NewJwtTokenStore(
		storex.NewRedisStore(rds, storex.WithPrefix("blog:app:token")),
		cfg.JWT.AccessSecret,
		"ve-blog-gin",
		cfg.JWT.AccessExpire,
		cfg.JWT.RefreshExpire,
	)

	return &ServiceContext{
		Config:       cfg,
		DB:           initialize.Gorm(cfg),
		Redis:        rds,
		RabbitMQ:     initialize.RabbitMQ(cfg),
		TokenStore:   tokenStore,
		UserAuth:     middleware.UserAuth(tokenStore),
		AdminAuth:    middleware.AdminAuth(tokenStore),
		RateLimit:    middleware.Limit(rds),
		Permission:   middleware.Permission(permissionx.NewNoopEnforcer()),
		OperationLog: middleware.OperationLog(),
	}
}
