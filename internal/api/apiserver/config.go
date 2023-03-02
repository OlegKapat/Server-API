package apiserver



type Config struct {
	BindAdrr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	DataBaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		BindAdrr: ":8080",
		LogLevel: "debug",
		
	}
}
