package config

const (
	Host     = "localhost"
	Port     = "5432"
	User     = "postgres"
	Password = "1234"
	DbName   = "postgres"

	Region  = "us-east-1"
	Enpoint = "http://localhost:8000"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

type Config struct {
	Type           DatabaseConfig
	DynamoDBConfig DynamoConfig
}

type DynamoConfig struct {
	TableName string
	Region    string
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
		DynamoDBConfig: DynamoConfig{
			TableName: "User",
			Region:    Region,
		},
	}
}
