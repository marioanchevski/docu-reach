package config

type Config struct {
	ListenAddr string
}

func NewStandardConfig() *Config {
	return &Config{
		ListenAddr: ":8080",
	}
}
