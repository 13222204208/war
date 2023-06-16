package config

type Setting struct {
	WinExperience  int `mapstructure:"win-experience" json:"winExperience" yaml:"win-experience"`
	LoseExperience int `mapstructure:"lose-experience" json:"loseExperience" yaml:"lose-experience"`
	WinScore       int `mapstructure:"win-score" json:"winScore" yaml:"win-score"`
	LoseScore      int `mapstructure:"lose-score" json:"loseScore" yaml:"lose-score"`
}
