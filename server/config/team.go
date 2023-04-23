package config

type Team struct {
	Create uint `mpstructure:"create" json:"create" yaml:"create"`
	Update uint `mpstructure:"update" json:"update" yaml:"update"`
}
