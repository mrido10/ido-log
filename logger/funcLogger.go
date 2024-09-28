package log

import (
	"fmt"
	"log"
	"strings"

	"github.com/mrido10/ido-log/constanta"
	"github.com/mrido10/ido-log/runtime"
)

var isPrintCaller bool

func SetIsPrintCaller(isPrint bool) {
	isPrintCaller = isPrint
}
func doPrint(level constanta.Info, caller int, fun func() string, msg ...any) {
	log.Println(level, setMessage(caller, msg...), fun())
}

func doPrintPanic(level constanta.Info, caller int, err ...any) {
	m := setMessage(caller, err...)
	log.Println(level, m)
	panic(m)
}

func doPrintFatal(level constanta.Info, caller int, err ...any) {
	log.Fatalln(level, setMessage(caller, err...))
}

func setMessage(caller int, msg ...any) string {
	call := 5
	if caller > 0 {
		call = caller
	}

	var m []string
	setFileCaller(call, &m)
	for _, v := range msg {
		m = append(m, fmt.Sprintf("%v", v))
	}
	return strings.Join(m, " ")
}

func setFileCaller(call int, m *[]string) {
	if isPrintCaller {
		*m = append(*m, runtime.GetRuntimeCaller(call))
		return
	}
}

func FuncErrorFilePositions() string {
	var errFilePositions []string
	maxIterate := 10
	start := 5
	if !isPrintCaller {
		start = 4
	}
	for i := start; i <= maxIterate; i++ {
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
