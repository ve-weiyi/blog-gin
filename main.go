package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/vkit/adapter/logz"

	"github.com/ve-weiyi/ve-blog-gin/api"
	"github.com/ve-weiyi/ve-blog-gin/config"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

var (
	configFile     = flag.String("f", "config.yaml", "the config file")
	nacosHost      = flag.String("nacos-host", "veweiyi.cn", "Input Your Nacos Host")
	nacosPort      = flag.Uint64("nacos-port", 8848, "Input Your Nacos Port")
	nacosUsername  = flag.String("nacos-username", "nacos", "Input Your Nacos Username")
	nacosPassword  = flag.String("nacos-password", "nacos", "Input Your Nacos Password")
	nacosNamespace = flag.String("nacos-namespace", "dev", "Input Your Nacos NameSpaceId")
	nacosGroup     = flag.String("nacos-group", "blog", "nacos group")
	nacosDataID    = flag.String("nacos-data-id", "ve-blog-gin", "Input Your Nacos DataId")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

// @title						ve-blog-gin API
// @version					1.0.0
// @description				ve-blog-gin 博客系统 API 文档
// @termsOfService			https://github.com/ve-weiyi/ve-blog-gin
// @contact.name			API Support
// @contact.url			https://github.com/ve-weiyi
// @contact.email			support@swagger.io
// @license.name			MIT
// @license.url			https://opensource.org/licenses/MIT
// @host						localhost:9090
// @BasePath
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						x-token
func main() {
	flag.Parse()

	var err error
	var c *config.Config

	if *configFile != "" {
		c, err = config.LoadConfigFromFile(*configFile, "yaml")
	} else {
		c, err = config.LoadConfigFromNacos(&config.LoadNacosConfigOption{
			ConfigType:      "yaml",
			NacosHost:       *nacosHost,
			NacosPort:       *nacosPort,
			NacosNamespace:  *nacosNamespace,
			NacosDataID:     *nacosDataID,
			NacosGroup:      *nacosGroup,
			NacosRuntimeDir: "runtime/nacos",
			NacosUsername:   *nacosUsername,
			NacosPassword:   *nacosPassword,
		})
	}
	if err != nil {
		log.Fatalf("failed to initialize config: %v\n", err)
	}

	// 初始化zap日志库
	logz.SetLog(&logz.LogConfig{
		Level:      c.Zap.Level,
		Mode:       c.Zap.Mode,
		Filename:   c.Zap.Filename,
		MaxSize:    c.Zap.MaxSize,
		MaxBackups: c.Zap.MaxBackups,
		MaxAge:     c.Zap.MaxAge,
		Compress:   c.Zap.Compress,
	})
	logz.L().Sugar().Infof("zap log init success. mode:%v, level:%v", c.Zap.Mode, c.Zap.Level)

	ctx := svctx.NewServiceContext(c)

	// 设置ReleaseMode则不会打印路由注册日志
	gin.SetMode(gin.DebugMode)
	engine := gin.Default()
	api.RegisterRouters(engine, ctx)

	logz.L().Sugar().Infof("register router success")

	address := fmt.Sprintf(":%d", c.Server.Port)
	s := &http.Server{
		Addr:           address,
		Handler:        engine,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	logz.L().Sugar().Infof("run server on http://localhost:%v success", c.Server.Port)

	fmt.Printf(`
    欢迎使用 ve-blog-gin
    当前版本: %s
    默认接口文档地址: http://localhost:%v/api/v1/swagger/index.html
`, c.Server.Version, c.Server.Port)
	fmt.Println(s.ListenAndServe().Error())
}
