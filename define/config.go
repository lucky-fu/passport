// Package define...
//
// File:  config
//
// Description: config
//
// Date: 2021/5/19 下午11:38
package define

// Config 用到的配置
type Config struct {
	Base  Base  `json:"base"`
	Redis Redis `json:"redis"`
}

// Base ...
type Base struct {
	Port   int    `yml:"port"`
	Env    string `yml:"env"`
	Domain string `yml:"domain"`
	Shard  int    `yml:"shard"`
}

// Redis ...
type Redis struct {
	Host     string       `yml:"host"`
	Port     int          `yml:"port"`
	DB       int          `yml:"db"`
	Password string       `yml:"password"`
	Timeout  RedisTimeout `yml:"timeout"`
}

// RedisTimeout ...
type RedisTimeout struct {
	Connect int `yml:"connect"`
	Read    int `yml:"read"`
	Write   int `yml:"write"`
}

// ConfigRedis ...
type ConfigRedis struct {
	Host     string `mapstructure:"host,omitempty"`
	Port     string `mapstructure:"port,omitempty"`
	DB       int    `mapstructure:"db,omitempty"`
	Password string `mapstructure:"password,omitempty"`

	ConnectTimeout float32 `mapstructure:"connect_timeout,omitempty"`
	WriteTimeout   float32 `mapstructure:"write_timeout,omitempty"`
	ReadTimeout    float32 `mapstructure:"read_timeout,omitempty"`
}
