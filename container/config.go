package container

// 将配置读取到内存中
import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/mitchellh/mapstructure"
	"github.com/passport/define"
)

var (
	configMap *sync.Map
)

// init ...
func init() {
	configMap = &sync.Map{}
}

// getAbsPath 获取绝对路径
func getAbsPath(base, path, suffix string) string {
	if len(suffix) <= 0 {
		suffix = ".yml"
	}
	return fmt.Sprintf("%s/%s%s", base, path, suffix)
}

// 读取文件内容
func readFile(base, path, suffix string) ([]byte, error) {
	abPath := getAbsPath(base, path, suffix)
	file, err := os.OpenFile(abPath, os.O_RDONLY, 0644)

	if err != nil {
		return nil, err
	}

	defer func() {
		if file != nil {
			_ = file.Close()
		}
	}()

	return ioutil.ReadFile(abPath)
}

// 读取内存中的配置
func get(base, path string) (interface{}, error) {
	var (
		exists bool
		err    error
		value  interface{}
	)

	value, exists = configMap.Load(path)

	if !exists {
		value, err = readFile(base, path, "")
		if err != nil {
			return nil, err
		}
	}

	configMap.Store(path, value)

	return value, nil
}

// GetRedisConfigList 获取redis的配置列表
func GetRedisConfig(path string) (*define.ConfigRedis, error) {
	var redisConfig *define.ConfigRedis

	rawConfig, err := get("", path)

	if err != nil {
		return nil, err
	}

	err = mapstructure.Decode(rawConfig, &redisConfig)

	return redisConfig, nil
}

// GetRootPath 获取根目录
func GetRootPath() (rootPath string, err error) {
	rootPath, err = os.Getwd()
	return
}

// GetLogConfig 获取日志配置
func GetLogConfig(path string) (*define.Log, error) {
	var log define.Log

	err := getConfig(path, &log)

	return &log, err
}

// getConfig ...
func getConfig(path string, decodeConf interface{}) error {
	rawConfig, err := get("", path)

	if err != nil {
		return err
	}

	err = mapstructure.Decode(rawConfig, decodeConf)

	return err
}
