package util

import "context"

type LogProvider interface {
	Init() error
	WriteMsg(ctx *context.Context, msg string, trace string, level int) error
	Destroy() error
	Flush() error
}
