package main

import (
	"encoding/json"
	log "github.com/mrido10/ido-log/logger"
	"github.com/mrido10/ido-log/runtime"
)

func main() {
	// set the log is print file info or not
	// default false > is not print info
	log.SetIsPrintCaller(true)

	// get file info
	LoggingCallFileInfo()

	// example to logging
	example1()
}

func LoggingCallFileInfo() {
	log.Debug(runtime.GetRuntimeCaller(1))
}

func example1() {
	example2()
}

func example2() {
	example3()
}
func example3() {
	example4()
}
func example4() {
	example5()
}

type example struct {
	Field1 string
	Field2 int
	Field3 float32
	Field4 bool
	Field5 interface{}
}

func example5() {
	ex := example{
		Field1: "example",
		Field2: 10,
		Field3: 20.01,
		Field4: true,
		Field5: nil,
	}

	b, _ := json.Marshal(ex)
	log.Info("this is INFO", string(b))
	log.Debug("this is DEBUG", ex.Field1, ex.Field2)
	log.Error("this is ERROR", nil)
	//log.Panic("this is PANIC")
	log.Fatal("this is FATAL")
}
