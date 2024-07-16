package logger

import "fmt"

type logger struct{}

func NewLogger() *logger {
	return &logger{}
}

func (l *logger) Debug(msg string) {
	fmt.Printf("[DEBUG] %s\n", msg)
}

func (l *logger) Error(err error) {
	fmt.Printf("[ERROR] %s\n", err.Error())
}
