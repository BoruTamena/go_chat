package dto

type Config struct {
	Db     DbConfig `mapstructure:"database"`
	Server Server   `mapstructure:"server"`
}

type DbConfig struct {
	Url  string `mapstructure:"url"`
	Name string `mapstructure:"db"`
}

type Server struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}
