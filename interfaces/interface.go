package main

import "fmt"

func main() {
	consoleLogger := ConsoleLogger{}
	fileLogger := FileLogger{}

	logService1 := LogService{logger: consoleLogger}
	logService2 := LogService{logger: fileLogger}

	logService1.logger.Log("console log value")
	logService2.logger.Log("file log value")
}

type LogService struct {
	logger Logger // Interface
}

type Logger interface {
	Log(str string) string
}

type ConsoleLogger struct{}

func (c ConsoleLogger) Log(str string) string {
	fmt.Println("logging to console: " + str)
	return str
}

type FileLogger struct{}

func (c FileLogger) Log(str string) string {
	fmt.Println("logging to file: " + str)
	return str
}
