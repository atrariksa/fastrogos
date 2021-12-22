package configs

import (
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Config is a struct that will receive configuration options via environment
// variables.
type Config struct {
	App struct {
		Hostname string `mapstructure:"HOST_NAME"`
		LogLevel string `mapstructure:"LOG_LEVEL"`

		Shutdown struct {
			CleanupPeriodSeconds int64 `mapstructure:"CLEANUP_PERIOD_SECONDS"`
			GracePeriodSeconds   int64 `mapstructure:"GRACE_PERIOD_SECONDS"`
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

		Page struct {
			TemplatesFolder string `mapstructure:"TEMPLATES_FOLDER"`
			AssetsFolder    string `mapstructure:"ASSETS_FOLDER"`
		} `mapstructure:"PAGE"`
	}

	DB struct {
		MySQL struct {
			Read struct {
				Host     string `mapstructure:"HOST"`
				Port     string `mapstructure:"PORT"`
				Username string `mapstructure:"USER"`
				Password string `mapstructure:"PASSWORD"`
				Name     string `mapstructure:"NAME"`
				Timezone string `mapstructure:"TIMEZONE"`
			}
			Write struct {
				Host     string `mapstructure:"HOST"`
				Port     string `mapstructure:"PORT"`
				Username string `mapstructure:"USER"`
				Password string `mapstructure:"PASSWORD"`
				Name     string `mapstructure:"NAME"`
				Timezone string `mapstructure:"TIMEZONE"`
			}
		}
	}

	Rula struct {
		Hostname string `mapstructure:"HOSTNAME"`
	} `mapstructure:"RULA"`
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

// GetFrom are responsible to load env and get data an return the struct
func GetFrom(path string) *Config {
	viper.SetConfigFile(path)
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
