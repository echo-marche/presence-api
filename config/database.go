package config

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func InitHackTechTipsDBConfig() DBConfig {
	return DBConfig{
		Host:     GetEnv("HACK_TECH_TIPS_DB_HOST"),
		Port:     GetEnv("HACK_TECH_TIPS_DB_PORT"),
		User:     GetEnv("HACK_TECH_TIPS_DB_USER"),
		Password: GetEnv("HACK_TECH_TIPS_DB_PASSWORD"),
		Name:     GetEnv("HACK_TECH_TIPS_DB_NAME"),
	}
}
