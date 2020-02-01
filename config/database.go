package config

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func InitSounDiverseDBConfig() DBConfig {
	return DBConfig{
		Host:     GetEnv("SOUNDIVERSE_DB_HOST"),
		Port:     GetEnv("SOUNDIVERSE_DB_PORT"),
		User:     GetEnv("SOUNDIVERSE_DB_USER"),
		Password: GetEnv("SOUNDIVERSE_DB_PASSWORD"),
		Name:     GetEnv("SOUNDIVERSE_DB_NAME"),
	}
}
