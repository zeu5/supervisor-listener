package config

type Config interface{}

func ParseConfig(flags map[string]string) *Config {
	var config Config
	return &config
}
