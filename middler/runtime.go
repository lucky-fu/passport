package middler

import (
	"fmt"
	"runtime"
)

// InitRuntime
func InitRuntime() {
	// 设置Processor和线程的数量
	cpuNum := runtime.NumCPU()
	process := runtime.GOMAXPROCS(cpuNum)

	fmt.Printf("cpu的数目为：%d, 设置的调度器数目为：%d", cpuNum, process)
}
