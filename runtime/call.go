package runtime

import (
	"fmt"
	"runtime"
	"strings"
)

func GetRuntimeCaller(skip int) string {
	pc, fileName, line, _ := runtime.Caller(skip)
	fn := runtime.FuncForPC(pc)

	fnSplit := strings.Split(fn.Name(), ".")
	fnName := ""
	if len(fnSplit) == 2 {
		fnName = fnSplit[1]
	}
	return fmt.Sprintf("%s>%s:%d", fileName, fnName, line)
}
