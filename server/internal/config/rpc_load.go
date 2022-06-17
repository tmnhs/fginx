package config

//rpc远程调用服务
type RpcLoad struct {
	Init       string `mapstructure:"init" json:"init" yaml:"init"`
	AppKey     string `mapstructure:"app-key" json:"appKey" yaml:"app-key"`
	ServiceUri string `mapstructure:"service-uri" json:"serviceUri" yaml:"service-uri"`
}
