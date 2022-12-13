package config

import (
	"github.com/fsnotify/fsnotify"
	"time"

	"github.com/spf13/viper"
)

var Cfg = &Config{}

type Config struct {
	App           App           `mapstructure:"app"`
	Mysql         Mysql         `mapstructure:"mysql"`
	MongoDB       MongoDB       `mapstructure:"mongodb"`
	Elasticsearch Elasticsearch `mapstructure:"elasticsearch"`
	Redis         Redis         `mapstructure:"redis"`
	Prometheus    Prometheus    `mapstructure:"prometheus"`
}

type App struct {
	AppSignExpire   int64         `mapstructure:"app_sign_expire"`
	RunMode         string        `mapstructure:"run_mode"`
	HttpPort        int           `mapstructure:"http_port"`
	ReadTimeout     time.Duration `mapstructure:"read_timeout"`
	WriteTimeout    time.Duration `mapstructure:"write_timeout"`
	RuntimeRootPath string        `mapstructure:"runtime_root_path"`
	AppLogPath      string        `mapstructure:"app_log_path"`
}

type Mysql struct {
	DBName            string        `mapstructure:"dbname"`
	User              string        `mapstructure:"user"`
	Password          string        `mapstructure:"password"`
	Host              string        `mapstructure:"host"`
	MaxOpenConn       int           `mapstructure:"max_open_conn"`
	MaxIdleConn       int           `mapstructure:"max_idle_conn"`
	ConnMaxLifeSecond time.Duration `mapstructure:"conn_max_life_second"`
	TablePrefix       string        `mapstructure:"table_prefix"`
}

type MongoDB struct {
	DBName   string   `mapstructure:"dbname"`
	User     string   `mapstructure:"user"`
	Password string   `mapstructure:"password"`
	Host     []string `mapstructure:"host"`
}

type Elasticsearch struct {
	Host           []string `mapstructure:"host"`
	User           string   `mapstructure:"user"`
	Password       string   `mapstructure:"password"`
	BulkActionNum  int      `mapstructure:"bulk_action_num"`
	BulkActionSize int      `mapstructure:"bulk_action_size"` //kb
	BulkWorkersNum int      `mapstructure:"bulk_workers_num"`
}

type Redis struct {
	Host        string `mapstructure:"host"`
	DB          int    `mapstructure:"db"`
	Password    string `mapstructure:"password"`
	MinIdleConn int    `mapstructure:"min_idle_conn"`
	PoolSize    int    `mapstructure:"pool_size"`
	MaxRetries  int    `mapstructure:"max_retries"`
}

type Prometheus struct {
	Host string `mapstructure:"host"`
}

// LoadConfig 加载配置
func LoadConfig() {
	v := viper.New()
	v.SetConfigFile("config/config.yml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(Cfg); err != nil {
		panic(err)
	}

	// 动态刷新配置
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		if err := v.Unmarshal(Cfg); err != nil {
			panic(err)
		}
	})
}
