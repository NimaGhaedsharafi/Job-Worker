package config

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/go-playground/validator/v10"
	"time"
)

// nolint:gochecknoglobals
var (
	// Cfg represents config of the project loaded into a struct
	Cfg Config
)

func Reset() {
	viper.Reset()
	Cfg = Config{}
}

// Init initialize config
func Init() {
	v := viper.New()
	v.SetConfigType("yaml")
	v.AddConfigPath("/etc")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		logrus.Panicf("error loading default configs: %s", err)
		return
	}

	err := v.MergeInConfig()
	if err != nil {
		logrus.Info("no config file found. Using defaults and environment variables")
	}

	if err := v.UnmarshalExact(&Cfg); err != nil {
		logrus.Panicf("invalid configuration: %s", err)
	}

	if err := Cfg.Validate(); err != nil {
		logrus.Panicf("invalid configuration: %s", err)
	}

	if Cfg.Debug {
		cfgBytes, _ := json.Marshal(Cfg)
		logrus.Infof("config: loaded config: %s", string(cfgBytes))
	}
}

type (
	Config struct {
		Env   string `mapstructure:"env" validate:"required"`
		Debug bool   `mapstructure:"debug"`

		Logger Logger `mapstructure:"logger" validate:"required"`
		Mysql Mysql `mapstructure:"database" validate:"required"`
		Storage  Storage  `mapstructure:"storage"`
	}

	Logger struct {
		Level string `mapstructure:"level" validate:"required"`
	}

	Mysql struct {
		Host               string        `mapstructure:"host" validate:"required"`
		Port               int           `mapstructure:"port" validate:"required"`
		Username           string        `mapstructure:"user" validate:"required"`
		Password           string        `mapstructure:"pass" validate:"required"`
		DBName             string        `mapstructure:"dbname" validate:"required"`
		ConnectTimeout     time.Duration `mapstructure:"connect-timeout" validate:"required"`
		ConnectionLifetime time.Duration `mapstructure:"connection-lifetime" validate:"required"`
		MaxOpenConnections int           `mapstructure:"max-open-connections"`
		MaxIdleConnections int           `mapstructure:"max-idle-connections"`
	}

	Storage struct {
		Adrs   string `mapstructure:"adrs" validate:"required"`
		Bucket string `mapstructure:"bucket" validate:"required"`
		Domain string `mapstructure:"domain" validate:"required"`
	}
)

func (c Config) Validate() error {
	return validator.New().Struct(c)
}
