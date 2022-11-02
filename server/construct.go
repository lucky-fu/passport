package server

import "github.com/passport/middler"

func construct() {
	initContainer()
}

func initContainer() {
	middler.InitRuntime()
	middler.InitVariables()
	middler.InitConfig()
	middler.InitRedis()
}
