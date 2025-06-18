package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// Создаем конфиг который будет читать наш local.yaml файл.
// yaml:"env" - Наименование поля в yaml.
// env-default:"dev" - Если в yaml нет параметра, установим по дефолту выбранный здесь.
// env-required:"true" - Если поле не задано в yaml, крашим запуск приложения.

const AliasMaxLength = 6

type Config struct {
	Env        string      `yaml:"env" env-default:"dev"`
	Postgres   PostgresCfg `yaml:"postgres" env-required:"true"`
	Sqlite     SqliteCfg   `yaml:"sqlite" env-required:"true"`
	HTTPServer HTTPServer  `yaml:"http_server"`
	DBType     string      `yaml:"db_type" env-required:"true"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	User string `yaml:user env-required:"true"`
	Password string `yaml:password env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}

type SqliteCfg struct {
	Path string `yaml:"storage_path" env-required:"true"`
}
type PostgresCfg struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     int    `yaml:"port" env-required:"true"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	DBName   string `yaml:"dbname" env-required:"true"`
	SSLMode  string `yaml:"sslmode" env-default:"disable"`
}

// Функция, которая прочитает файл с конфигом и заполнит объект Config.
// Функция называется MustLoad, Must - обозначает, что вместо ошибки, функция будет паниковать. (Делается в редких случаях)
func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// Проверяем, что файл существует.
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file is not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg

}
