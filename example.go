package main

import (
	log "github.com/mrido10/ido-log/logger"
)

func main() {
	example1()
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

	log.Info("this is INFO", ex)
	log.Debug("this is DEBUG", ex.Field1, ex.Field2)
	log.Error("this is ERROR", nil)
	log.Panic("this is PANIC")
	log.Fatal("this is FATAL")
}
