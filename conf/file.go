package conf

type file struct {
	Dir     string `mapstructure:"dir"`
	MaxSize int    `mapstructure:"max_size"`
}
