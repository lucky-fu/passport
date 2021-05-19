package container

// 将配置读取到内存中
import (
	"sync"
)

var (
	configMap *sync.Map
)

// init ...
func init() {
	configMap = &sync.Map{}
}

// getAbsPath 获取绝对路径
func getAbsPath(base, path string) string {
	return ""
}

// 读取文件内容
func readFile(base, path string) ([]byte, error) {
	var err error
	return nil, err
}

// parseYaml 解析yaml文件内容
func parseYaml(content []byte) (interface{}, error) {
	var err error
	return nil, err
}
