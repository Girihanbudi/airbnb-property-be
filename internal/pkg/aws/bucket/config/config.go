package config

type Config struct {
	Region string `mapstructure:"region"`
	Bucket string `mapstructure:"bucket"`
	Key    string `mapstructure:"key"`
}
