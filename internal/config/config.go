package config

type Config struct {
	ListenHost string
	ListenPort int
}

// NewConfig creates a default configuration
func NewConfig() *Config {
	return &Config{
		ListenHost: "localhost",
		ListenPort: 8080,
	}
}

// NewConfigWithOptions creates a configuration with parameters
func NewConfigWithOptions(host string, port int) *Config {
	return &Config{
		ListenHost: host,
		ListenPort: port,
	}
}
