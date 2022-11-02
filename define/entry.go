package define

// Config
type Config struct {
	Base  Base
	Redis Redis
	Mysql Mysql
	Log   Log
}

// Redis
type Redis struct {
	Host     string       `yml:"host"`
	Port     int          `yml:"port"`
	DB       int          `yml:"db"`
	Password string       `yml:"password"`
	Timeout  RedisTimeout `yml:"timeout"`
}

// Mysql
type Mysql struct {
	Host     string `yml:"host"`
	Port     int    `yml:"port"`
	Name     string `yml:"name"`
	User     string `yml:"user"`
	Password string `yml:"password"`
	Level    int    `yml:"level"`
	Max      int    `yml:"max"`
}

// Base ...
type Base struct {
	Env string `yml:"env"`
}

// RedisTimeout ...
type RedisTimeout struct {
	Connect int `yml:"connect"`
	Read    int `yml:"read"`
	Write   int `yml:"write"`
}

// Log ...
type Log struct {
	AccessLog    LogConfig
	APILog       LogConfig
	SQLLog       LogConfig
	BuesinessLog LogConfig
}

// LogConfig ...
type LogConfig struct {
	Path    string `yml:"path"`
	Level   int    `yml:"level"`
	Buffer  int    `yml:"buffer"`
	Console bool   `yml:"console"`
}
