package config

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func InitNicheyeDBConfig() DBConfig {
	return DBConfig{
		Host:     GetEnv("NICHEYE_DB_HOST"),
		Port:     GetEnv("NICHEYE_DB_PORT"),
		User:     GetEnv("NICHEYE_DB_USER"),
		Password: GetEnv("NICHEYE_DB_PASSWORD"),
		Name:     GetEnv("NICHEYE_DB_NAME"),
	}
}
