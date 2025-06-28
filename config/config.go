package config

type Config struct {
    ServerAddress string `mapstructure:"SERVER_ADDRESS"`
    DBUri string `mapstructure:"DB_URI"`
}