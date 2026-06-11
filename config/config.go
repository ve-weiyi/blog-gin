package config

type Config struct {
	Server   ServerConf   `json:"system"`
	JWT      JWTConf      `json:"jwt"`
	Mysql    MysqlConf    `json:"mysql"`
	Redis    RedisConf    `json:"redis"`
	RabbitMQ RabbitMQConf `json:"rabbitmq"`
	Zap      Zap          `json:"zap"`
}

// 系统配置
type ServerConf struct {
	Port        int    `json:"port" example:"8080"`              // 运行端口
	Env         string `json:"env" example:"dev"`                // 运行环境 dev、test、prod
	Version     string `json:"version" example:"v1.0.0"`         // 程序版本
	RuntimePath string `json:"runtime-path" example:"./runtime"` // 运行时目录
}

// jwt鉴权
type JWTConf struct {
	AccessSecret  string `json:"access_secret"`
	AccessExpire  int64  `json:"access_expire"`
	RefreshSecret string `json:"refresh_secret"`
	RefreshExpire int64  `json:"refresh_expire"`
}

// mysql数据库
type MysqlConf struct {
	Host     string `json:"host" example:"localhost"`
	Port     string `json:"port" example:"3306"`
	Username string `json:"username" example:"root"`
	Password string `json:"password" example:"123456"`
	Database string `json:"database"`
	Config   string `json:"config" example:"charset=utf8mb4&parseTime=True&loc=Local"`
}

// redis缓存
type RedisConf struct {
	Host     string `json:"host" example:"localhost"` // 服务器地址
	Port     string `json:"port" example:"6379"`      // 端口
	Password string `json:"password" example:""`      // 密码
	DB       int    `json:"db" example:"0"`           // redis的哪个数据库
}

// rabbitmq消息队列
type RabbitMQConf struct {
	Host     string `json:"host" example:"localhost"`
	Port     string `json:"port" example:"5672"`
	Username string `json:"username" example:"guest"`
	Password string `json:"password" example:"guest"`
}

type Zap struct {
	Level      string `json:"level" example:"debug"`
	Mode       string `json:"mode" example:"console|file"` // console|file
	Filename   string `json:"file_name" example:"./runtime/app.log"`
	MaxSize    int    `json:"max_size" example:"100"`
	MaxBackups int    `json:"max_backups" example:"3"`
	MaxAge     int    `json:"max_age" example:"30"`
	Compress   bool   `json:"compress" example:"true"`
}
