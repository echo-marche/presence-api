package config

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func InitDevCompassDBConfig() DBConfig {
	return DBConfig{
		Host:     GetEnv("DEV_COMPASS_DB_HOST"),
		Port:     GetEnv("DEV_COMPASS_DB_PORT"),
		User:     GetEnv("DEV_COMPASS_DB_USER"),
		Password: GetEnv("DEV_COMPASS_DB_PASSWORD"),
		Name:     GetEnv("DEV_COMPASS_DB_NAME"),
	}
}
