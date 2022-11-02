package middler

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	ENV      string
	Port     string
	RootPath string
)

// InitVariables ...
func InitVariables() {
	var (
		err error
		ok  bool
	)

	if RootPath, err = os.Getwd(); err != nil {
		panic(fmt.Sprintf("get root path failed : %v", err.Error()))
	}

	if ENV, ok = os.LookupEnv("ENV"); !ok {
		panic("get env failed")
	}

	if Port, ok = os.LookupEnv("PORT"); !ok {
		panic("get port failed")
	}

	fmt.Printf("初始化变量成功\n")
}
