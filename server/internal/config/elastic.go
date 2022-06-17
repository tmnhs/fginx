package config

//Elastic搜索引擎
type Elastic struct {
	Path     string `mapstructure:"path" json:"path" yaml:"path"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`             // 端口
	UserName string `mapstructure:"username" json:"username" yaml:"username"` // 用户名
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
}
