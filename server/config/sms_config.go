package config

type SmsConfig struct {
	Signature        string `mapstructure:"signature" json:"signature" yaml:"signature"`                        // 短信签名
	VerificationCode string `mapstructure:"verification-code" json:"verificationCode" yaml:"verification-code"` // 短信验证码模板
	Key              string `mapstructure:"key" json:"key" yaml:"key"`                                          // 短信验证码key
	Secret           string `mapstructure:"secret" json:"secret" yaml:"secret"`
}
