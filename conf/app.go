package conf

type app struct {
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	AuthKey    string `mapstructure:"auth_key"`
	AuthSecret string `mapstructure:"auth_secret"`
	AuthExpire int    `mapstructure:"auth_expire"`
}
