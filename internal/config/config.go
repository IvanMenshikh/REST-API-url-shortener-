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

type Config struct {
	Env         string `yaml:"env" env-default:"dev"`
	PostgresCfg `yaml:"storage_path" env-required:"true"`
	Sqlite      `yaml:"sqlite" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
	DBType      string `yaml:"db_type" env-required:"true"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Sqlite struct {
	Path string `yaml:"storage_path" env-required:"true"`
}
type PostgresCfg struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     int    `yaml:"port" env-default:"5432"`
	User     string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password" env-default:""` // чувствительные данные
	DBName   string `yaml:"dbname" env-default:"url-shortener"`
	SSLMode  string `yaml:"" env-default:"disable"`
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
