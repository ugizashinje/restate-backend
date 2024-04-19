package config

import (
	"crypto/ed25519"
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"gopkg.in/yaml.v3"
)

type DBConfig struct {
	Host      string
	Port      int
	User      string
	Password  string
	Dbname    string
	SslMode   string
	DbLogging bool
}

type RedisConfig struct {
	Password string `yaml:"password"`
	Username string `yaml:"username"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

type AuthConfig struct {
	JwtPemFile       string `yaml:"jwtPemFile"`
	SigningMethod    string `yaml:"signingMethod"`
	TsMailHost       string `yaml:"tsMailHost"`
	TsMailPort       int    `yaml:"tsMailPort"`
	WarrantServer    string `yaml:"warrantServer"`
	RateLimitPerSlot int    `yaml:"rateLimitPerSlot"`
	RateLimitTotal   int    `yaml:"rateLimitTotal"`
}

type FormatConfig struct {
	DateFormat string `yaml:"dateFormat"`
	TimeFormat string `yaml:"timeFormat"`
}

type StorageConfig struct {
	Endpoint        string `yaml:"endpoint"`
	Bucket          string `yaml:"bucket"`
	AccessKeyID     string `yaml:"accessKeyId"`
	SecretAccessKey string `yaml:"secretAccessKey"`
	UseSsl          bool   `yaml:"useSsl"`
}

type Config struct {
	Db       DBConfig       `yaml:"db"`
	Redis    RedisConfig    `yaml:"redis"`
	Auth     AuthConfig     `yaml:"auth"`
	Format   FormatConfig   `yaml:"format"`
	Storage  StorageConfig  `yaml:"storage"`
	Superset SupersetConfig `yaml:"superset"`
}
type SupersetConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Url      string `yaml:"url"`
}

func Init(env string) {
	var config Config
	dir := "./"
	if env == "testing" {
		dir = "../"
	}
	yamlFile, err := os.ReadFile(dir + "conf.yaml")
	if err != nil {
		log.Fatal("yamlFile.Get err  # ", err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		os.Exit(1)
	}
	Db = config.Db
	Redis = config.Redis
	Auth = config.Auth
	Format = config.Format
	Storage = config.Storage
	Superset = config.Superset

	if err := loadPrivateKeys(config.Auth, dir); err != nil {
		panic(err)
	}
}

var Db DBConfig
var Redis RedisConfig
var Auth AuthConfig
var Format FormatConfig
var JwtPrivateKey ed25519.PrivateKey
var Storage StorageConfig
var Superset SupersetConfig

func loadPrivateKeys(auth AuthConfig, dir string) error {

	data, err := os.ReadFile(dir + auth.JwtPemFile)
	if err != nil {
		return fmt.Errorf("createToken: failed to read p8 key: %v", err)
	}

	crytoPrivate, err := jwt.ParseEdPrivateKeyFromPEM(data)
	if err != nil {
		return fmt.Errorf("Cant parse ed25519 private key from pem : %v", err)
	}

	// if err != nil {
	// 	return fmt.Errorf("createToken: AuthKey must be of type ecdsa.PrivateKey %v", err)
	// }
	JwtPrivateKey = crytoPrivate.(ed25519.PrivateKey)

	return nil
}
