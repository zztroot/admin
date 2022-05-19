package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var GlobalConfig *Config

type Config struct {
	AppConfig  *AppConfig    `json:"app_config"`
	CronConfig []*CronConfig `json:"cron_config"`
}

func InitConfig(path string) {
	// 这里其实可以加一个监控，实现热更新配置
	read, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	c := new(Config)
	if err := json.Unmarshal(read, c); err != nil {
		log.Fatalln(err)
	}
	// 全局赋值
	GlobalConfig = c
}

// 应用配置
type AppConfig struct {
	AppName        string               `json:"app_name"`
	Host           string               `json:"host"`
	Port           int                  `json:"port"`
	TrustedProxies []string             `json:"trusted_proxies"` // 信任代理
	GinMode        string               `json:"gin_mode"`        // 运行模式 debug|release
	Database       map[string]*Database `json:"database"`        // string首先拿default配置
	Redis          *Redis               `json:"redis"`
	Jaeger         *Jaeger              `json:"jaeger"`
}

// 定时任务配置
type CronConfig struct {
	CronName    string `json:"cron_name"`
	Schedule    string `json:"schedule"`
	IsRun       bool   `json:"is_run"`
	Parallelism int    `json:"parallelism"`
}

type Database struct {
	Dialect        string `json:"dialect"`
	Database       string `json:"database"`
	User           string `json:"user"`
	Password       string `json:"password"`
	Host           string `json:"host"`
	Port           int    `json:"port"`
	Charset        string `json:"charset"`
	MaxIdleConnNum int    `json:"max_idle_conn_num"`
	MaxOpenConnNum int    `json:"max_open_conn_num"`
}

type Redis struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Password    string `json:"password"`
	MaxIdle     int    `json:"max_idle"`
	MaxActive   int    `json:"max_active"`
	IdleTimeout int    `json:"idle_timeout"`
	KeyPrefix   string `json:"key_prefix"`
}

type Jaeger struct {
	Host string  `json:"host"`
	Port int     `json:"port"`
	Name string  `json:"name"`
	Rate float64 `json:"rate"` // 采样率
}
