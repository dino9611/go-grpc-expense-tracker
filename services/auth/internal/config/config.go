package config

type Config struct {
	Application AppConfig
	Database    DBConfig
}

type AppConfig struct {
	Environment string
	Port        int
	BaseUrl     string
}

type DBConfig struct {
	Host     string
	Port     int
	Password string
	User     string
	Name     string
}
