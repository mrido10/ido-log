package log

import (
	"github.com/mrido10/ido-log/constanta"
)

func Debug(msg ...any) {
	doPrint(constanta.DEBUG, 0, func() string { return "" }, msg...)
}

func Info(msg ...any) {
	doPrint(constanta.INFO, 0, func() string { return "" }, msg...)
}

func Error(msg ...any) {
	doPrint(constanta.ERROR, 0, funcErrorFilePositions, msg...)
}

func Panic(msg ...any) {
	doPrintPanic(constanta.PANIC, 0, msg...)
}

func Fatal(msg ...any) {
	doPrintFatal(constanta.FATAL, 0, msg...)
}
