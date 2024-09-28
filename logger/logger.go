package log

import (
	"github.com/mrido10/ido-log/constanta"
)

// Debug : to print debug level
func Debug(msg ...any) {
	doPrint(constanta.DEBUG, 0, func() string { return "" }, msg...)
}

// Info : to print info level
func Info(msg ...any) {
	doPrint(constanta.INFO, 0, func() string { return "" }, msg...)
}

// Error : to print error level
func Error(msg ...any) {
	doPrint(constanta.ERROR, 0, FuncErrorFilePositions, msg...)
}

// Panic : to print panic level
func Panic(msg ...any) {
	doPrintPanic(constanta.PANIC, 0, msg...)
}

// Fatal : to print fatal level
func Fatal(msg ...any) {
	doPrintFatal(constanta.FATAL, 0, msg...)
}
