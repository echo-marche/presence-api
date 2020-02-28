package config

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func InitNicheFarmDBConfig() DBConfig {
	return DBConfig{
		Host:     GetEnv("NICHE_FARM_DB_HOST"),
		Port:     GetEnv("NICHE_FARM_DB_PORT"),
		User:     GetEnv("NICHE_FARM_DB_USER"),
		Password: GetEnv("NICHE_FARM_DB_PASSWORD"),
		Name:     GetEnv("NICHE_FARM_DB_NAME"),
	}
}
