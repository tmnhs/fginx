package config

//Kafka消息队列
type Kafka struct {
	Use   bool   `mapstructure:"use" json:"use" yaml:"use"`       //是否使用kafka消息队列中间件,
	Path  string `mapstructure:"path" json:"path" yaml:"path"`    //路径
	Port  string `mapstructure:"port" json:"port" yaml:"port"`    // 端口
	Topic string `mapstructure:"topic" json:"topic" yaml:"topic"` // 端口
}
