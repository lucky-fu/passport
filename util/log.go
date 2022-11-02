package util

import (
	"sync"
)

// Logger ...
type Logger struct {
	level      int
	logType    string
	buf        chan string
	isRunning  int32
	wg         *sync.WaitGroup
	stop       chan struct{}
	flush      chan *sync.WaitGroup
	providers  []LogProvider
	fileOffset int // 文件获取时使用
}
