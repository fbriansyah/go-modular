package config

type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
}

type DatabaseConfig struct {
	URL     string `mapstructure:"url"`
	Host    string `mapstructure:"host"`
	Port    string `mapstructure:"port"`
	Name    string `mapstructure:"name"`
	SSLMode string `mapstructure:"sslmode"`
}
