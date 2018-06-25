package config

type handlerConfig struct {
	name   string
	config interface{}
}

type eventConfig struct {
	name     string
	handlers []handlerConfig
}

type Config struct {
	handlers []handlerConfig
	events   []eventConfig
}

func ParseConfig(file string) (Config, error) {
	return Config{}, nil
}
