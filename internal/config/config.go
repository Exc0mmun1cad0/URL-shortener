package config

import (
	"flag"
	"os"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string     `yaml:"env"`
	HTTPServer HTTPServer `yaml:"http_server"`
	PostgreSQL PostgreSQL `yaml:"postgresql"`
	Redis      Redis      `yaml:"redis"`
}

type HTTPServer struct {
	Host        string        `yaml:"host" env-default:"localhost"`
	Port        int           `yaml:"port" env-default:"8080"`
	ReadTimeout time.Duration `yaml:"read_timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

type PostgreSQL struct {
}

type Redis struct {
}

var (
	config Config
	once   sync.Once
)

func MustLoad() *Config {
	once.Do(func() {
		configPath := fetchConfigPath()

		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			panic("config file does not exist: " + configPath)
		}

		if err := cleanenv.ReadConfig(configPath, &config); err != nil {
			panic("cannot read config: " + err.Error())
		}
	})

	return &config
}

func fetchConfigPath() (path string) {
	flag.StringVar(&path, "config", "", "path to config file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}

	return path
}
