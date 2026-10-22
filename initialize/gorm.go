package initialize

import (
	"fmt"
	"log"

	"github.com/ve-weiyi/blog-gin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Gorm(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s&parseTime=True&loc=Local",
		cfg.Mysql.Username, cfg.Mysql.Password, cfg.Mysql.Host,
		cfg.Mysql.Port, cfg.Mysql.Database, cfg.Mysql.Config,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("connect mysql failed: %v", err)
	}

	return db
}
