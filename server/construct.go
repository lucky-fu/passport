package server

import "github.com/passport/middler"

func construct() {
	initContainer()
}

func initContainer() {
	middler.InitRuntime()
}
