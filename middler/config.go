package middler

import (
	"fmt"

	"github.com/passport/define"
	"github.com/passport/util"
)

var (
	Config *define.Config
)

// InitConfig ...
func InitConfig() {
	Config = &define.Config{}

	if err := util.Parse.ParseYml(fmt.Sprintf("config/%s.yml", ENV), Config); err != nil {
		panic(fmt.Sprintf("parse yaml failed: %v", err.Error()))
	}

	fmt.Printf("初始化配置成功\n")
}
