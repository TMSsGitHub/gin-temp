package conf

type database struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Table     string `mapstructure:"table"`
	Charset   string `mapstructure:"charset"`
	ParseTime bool   `mapstructure:"parse_time"`
	Loc       string `mapstructure:"loc"`

	Prefix        string `mapstructure:"prefix"`
	SingularTable bool   `mapstructure:"singular_table"`
	LogDir        string `mapstructure:"log_dir"`
}
