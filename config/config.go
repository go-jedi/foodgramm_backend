package config

import (
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type LoggerConfig struct {
	IsJSON     bool   `yaml:"is_json"`
	AddSource  bool   `yaml:"add_source"`
	Level      string `yaml:"level"`
	SetFile    bool   `yaml:"set_file"`
	FileName   string `yaml:"file_name"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
}

type PostgresConfig struct {
	Host          string `yaml:"host"`
	User          string `yaml:"user"`
	Password      string `yaml:"password"`
	DBName        string `yaml:"dbname"`
	Port          int    `yaml:"port"`
	SSLMode       string `yaml:"sslmode"`
	PoolMaxConns  int    `yaml:"pool_max_conns"`
	MigrationsDir string `yaml:"migrations_dir"`
	QueryTimeout  int64  `yaml:"query_timeout"`
}

type HTTPServerConfig struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	GinMode string `yaml:"gin_mode"`
}

type Config struct {
	Logger     LoggerConfig     `yaml:"logger"`
	Postgres   PostgresConfig   `yaml:"postgres"`
	HTTPServer HTTPServerConfig `yaml:"httpserver"`
}

// LoadConfig load config file.
func LoadConfig() string {
	var cf string

	flag.StringVar(&cf, "config", "config.yaml", "config file path")
	flag.Parse()

	return cf
}

// ParseConfig parse config file.
func ParseConfig(configFile string) (config Config, err error) {
	f, err := os.Open(configFile)
	if err != nil {
		return config, err
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Printf("error closing the file: %v", err)
		}
	}(f)

	err = yaml.NewDecoder(f).Decode(&config)

	return config, err
}

// GetConfig get config.
func GetConfig() (config Config, err error) {
	cf := LoadConfig()

	return ParseConfig(cf)
}
