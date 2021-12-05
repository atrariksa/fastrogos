package configs

import (
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Hostname string `mapstructure:"HOST_NAME"`
		LogLevel string `mapstructure:"LOG_LEVEL"`

		Shutdown struct {
			GracePeriodSeconds int64 `mapstructure:"GRACE_PERIOD_SECONDS"`
		} `mapstructure:"SHUTDOWN"`

		CORS struct {
			Enable           bool     `mapstructure:"ENABLE"`
			AllowedOrigins   []string `mapstructure:"ALLOWED_ORIGINS"`
			AllowedMethods   []string `mapstructure:"ALLOWED_METHODS"`
			AllowedHeaders   []string `mapstructure:"ALLOWED_HEADERS"`
			AllowCredentials bool     `mapstructure:"ALLOW_CREDENTIALS"`
			MaxAgeSeconds    int      `mapstructure:"MAX_AGE_SECONDS"`
		} `mapstructure:"CORS"`

		DefaultRespHeaders struct {
			XFrameOptions       string   `mapstructure:"X_FRAME_OPTIONS"`
			XContentTypeOptions string   `mapstructure:"X_CONTENT_TYPE_OPTIONS"`
			XXSSProtection      string   `mapstructure:"X_XSS_PROTECTION"`
			CSP                 []string `mapstructure:"CSP"`
		} `mapstructure:"DEFAULT_RESP_HEADERS"`
	} `mapstructure:"APP"`

	Cache struct {
		Host          string        `mapstructure:"HOST"`
		Port          string        `mapstructure:"PORT"`
		DialTimeout   time.Duration `mapstructure:"DIAL_TIMEOUT"`
		ReadTimeout   time.Duration `mapstructure:"READ_TIMEOUT"`
		WriteTimeout  time.Duration `mapstructure:"WRITE_TIMEOUT"`
		IdleTimeout   time.Duration `mapstructure:"IDLE_TIMEOUT"`
		MaxConnAge    time.Duration `mapstructure:"MAX_CONN_AGE"`
		MinIdleConns  int           `mapstructure:"MIN_IDLE_CONNS"`
		Namespace     int           `mapstructure:"NAMESPACE"`
		Password      string        `mapstructure:"PASSWORD"`
		CacheDuration time.Duration `mapstructure:"CACHE_DURATION"`
	} `mapstructure:"CACHE"`

	DB struct {
		SQLITE struct {
			Name                 string `mapstructure:"NAME"`
			AdditionalParameters string `mapstructure:"ADDITIONAL_PARAMETERS"`
		} `mapstructure:"SQLITE"`
	} `mapstructure:"DB"`
}

var (
	conf Config
	once sync.Once
)

// Get are responsible to load env and get data an return the struct
func Get() *Config {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal().Err(err).Msg("Failed reading config file")
	}

	once.Do(func() {
		log.Info().Msg("Service configuration initialized.")
		err = viper.Unmarshal(&conf)
		if err != nil {
			log.Fatal().Err(err).Msg("")
		}
	})

	return &conf
}
