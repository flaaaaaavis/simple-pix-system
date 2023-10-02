package config

const (
	Host     = "localhost"
	Port     = "5432"
	User     = "root"
	Password = "password"
	DbName   = "postgres"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

type Config struct {
	Type DatabaseConfig
}

func NewConfig() *Config {
	return &Config{
		Type: DatabaseConfig{
			Host:     Host,
			Port:     Port,
			User:     User,
			Password: Password,
			DbName:   DbName,
		},
	}
}
