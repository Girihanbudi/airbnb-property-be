package config

type Config struct {
	Name      string `mapstructure:"name"`
	Separator string `mapstructure:"separator"`
}
