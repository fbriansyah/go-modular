package config

type Secret struct {
	Database DatabaseSecret `mapstructure:"database"`
}

type DatabaseSecret struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
