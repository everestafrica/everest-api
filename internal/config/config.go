package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"reflect"
	"strings"
)

const tagName = "env"

// Config is the configuration struct
type Config struct {
	Port              string `env:"PORT"`
	JWTSecret         string `env:"JWT_SECRET"`
	DatabaseURL       string `env:"DATABASE_URL"`
	RedisURL          string `env:"REDIS_URL"`
	RedisPassword     string `env:"REDIS_PASSWORD"`
	Env               string `env:"ENV"`
	LogLevel          string `env:"LOG_LEVEL"`
	ExpiryTime        string `env:"TOKEN_EXP_TIME"`
	MonoWebhookSecret string `env:"MONO_WEBHOOK_SECRET"`
	MonoSecretKey     string `env:"MONO_SECRET_KEY"`
	BscApiKey         string `env:"BSC_API_KEY"`
	EthApiKey         string `env:"ETH_API_KEY"`
	NewsApiUrl        string `env:"NEWS_API_URL"`
	NewsApiKeyPry     string `env:"NEWS_API_KEY_PRY"`
	NewsApiKeySec     string `env:"NEWS_API_KEY_SEC"`
	EmailSecretKey    string `env:"EMAIL_SECRET_KEY"`
	EmailFrom         string `env:"EMAIL_FROM"`
	EmailTo           string `env:"EMAIL_TO"`
	EmailDomainUrl    string `env:"EMAIL_DOMAIN_URL"`
	ProdSmsSecretKey  string `env:"PROD_SMS_SECRET_KEY"`
	ProdSmsPublicKey  string `env:"PROD_SMS_PUBLIC_KEY"`
}

var config *Config

func LoadConfig() (*Config, error) {
	godotenv.Load()
	var cfg = Config{}

	tempCfg := reflect.TypeOf(cfg)

	for i := 0; i < tempCfg.NumField(); i++ {
		f := tempCfg.Field(i)
		v := reflect.ValueOf(&cfg).Elem().FieldByName(f.Name)
		tagData := strings.Split(f.Tag.Get(tagName), ",")

		if len(tagData) == 0 {
			return nil, errors.New("invalid tag format")
		}

		var (
			envKey             = tagData[0]
			envValue, envFound = os.LookupEnv(envKey)
			required           = true
			isPtr              = false
		)

		if v.Kind() == reflect.String {
			// The field is a string
		} else if v.Kind() == reflect.Ptr && v.Type().Elem().Kind() == reflect.String {
			// The field is a *string
			required = false
			isPtr = true
		} else {
			// We don't support that type yet :(
			return nil, fmt.Errorf("warning: struct field %s must be of type string or *string", f.Name)
		}

		if required && !envFound {
			return nil, fmt.Errorf("env %s is required but not set", envKey)
		}

		if !v.CanSet() {
			return nil, fmt.Errorf("cannot set field %s", f.Name)
		}

		// Expand the environment variables before setting
		envValue = os.ExpandEnv(envValue)
		os.Setenv(envKey, envValue)

		if isPtr {
			if envFound {
				v.Set(reflect.ValueOf(&envValue))
			}
		} else {
			v.Set(reflect.ValueOf(envValue))
		}
	}

	config = &cfg

	return &cfg, nil
}

func GetConf() *Config {
	return config
}
