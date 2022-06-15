package config

type ServerConfig struct {
    Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
    Mysql  Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
