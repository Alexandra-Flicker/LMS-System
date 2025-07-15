package config

type Config struct {
	PostgresDB struct {
		User     string
		Password string
		Host     string
		Port     int
	}
	Server struct {
		Port string
	}
}
