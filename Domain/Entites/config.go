package Entites

type Config struct {
	Enable bool
	Port   string
}

func NewConfig() *Config {
	return &Config{
		Enable: true,
		Port:   "9090",
	}
}
