package log

import (
	"fmt"
	"log"
	"strings"

	"github.com/mrido10/ido-log/constanta"
	"github.com/mrido10/ido-log/runtime"
)

func doPrint(level constanta.Info, caller int, fun func() string, msg ...any) {
	call, m := setCallerAndMsg(caller, msg...)
	log.Println(level, runtime.GetRuntimeCaller(call), m, fun())
}

func doPrintPanic(level constanta.Info, caller int, err ...any) {
	call, m := setCallerAndMsg(caller, err...)
	log.Println(level, runtime.GetRuntimeCaller(call), m)
	panic(m)
}

func doPrintFatal(level constanta.Info, caller int, err ...any) {
	call, m := setCallerAndMsg(caller, err...)
	log.Fatalln(level, runtime.GetRuntimeCaller(call), m)
}

func setCallerAndMsg(caller int, msg ...any) (int, string) {
	call := 3
	if caller > 0 {
		call = caller
	}

	var m []string
	for _, v := range msg {
		m = append(m, fmt.Sprintf("%v", v))
	}
	return call, strings.Join(m, " ")
}

func funcErrorFilePositions() string {
	var errFilePositions []string
	maxIterate := 10
	for i := 5; i <= maxIterate; i++ {
		dir := strings.TrimSpace(runtime.GetRuntimeCaller(i))
		if dir == "/" || strings.Contains(dir, ">:") {
			break
		}
		errFilePositions = append(errFilePositions, dir)
		if i == maxIterate {
			errFilePositions = append(errFilePositions, "...")
		}
	}

	return "\n\t" + strings.Join(errFilePositions, "\n\t")
}
