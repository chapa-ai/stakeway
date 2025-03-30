package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type App struct {
	Port   string `yaml:"port" env:"APP_PORT" env-default:"8080"`
	Env    string `yaml:"env" env:"APP_ENV" env-default:"local"`
	ApiUrl string `yaml:"apiUrl" env:"PAY_AGGREGATOR_API" env-default:"http://localhost"`
}

type Db struct {
	Host    string `yaml:"host" env:"DB_HOST" env-default:"postgres"`
	Port    string `yaml:"port" env:"DB_PORT" env-default:"5432"`
	Name    string `yaml:"name" env:"DB_NAME" env-default:"postgres"`
	User    string `yaml:"user" env:"DB_USER" env-default:"postgres"`
	Pass    string `yaml:"pass" env:"DB_PASS" env-default:"test"`
	MaxConn int    `yaml:"maxConn" env:"DB_MAX_CONN" env-default:"50"`
	MaxIdle int    `yaml:"maxIdle" env:"DB_MAX_IDLE" env-default:"10"`
	Schema  string `yaml:"schema" env:"DB_SCHEMA" env-default:"public"`
}

type Log struct {
	Log string `yaml:"level" env:"LOG_LEVEL" env-default:"INFO"`
}

type BlockchainConfig struct {
	Enabled    bool   `yaml:"enabled"`
	PrivateKey string `yaml:"private_key"`
}

type Config struct {
	App        App              `yaml:"app"`
	DB         Db               `yaml:"db"`
	Log        Log              `yaml:"log"`
	Blockchain BlockchainConfig `yaml:"blockchain"`
}

func (cfg Config) GetDbConfig() Db {
	return cfg.DB
}

func (c Db) GetPort() string {
	return c.Port
}

func (c Db) GetUser() string {
	return c.User
}

func (c Db) GetPassword() string {
	return c.Pass
}

func (c Db) GetDatabase() string {
	return c.Name
}

func (c Db) GetDsn() string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		c.User,
		c.Pass,
		c.Host,
		c.Port,
		c.Name,
	)
}

func New() (Config, error) {
	cfg := Config{}
	err := cleanenv.ReadConfig("config/config.yml", &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
