package config

type Mini struct {
	Appid  string `mapstructure:"appid" json:"appid" yaml:"appid"`
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"`
}
