package monitoring

import (
	"fmt"
	"log"
)

const (
	LogLevelInfo = iota
	LogLevelCaution
	LogLevelWarning
	LogLevelError
)

var LogLevelToString = map[uint]string{
	LogLevelInfo:    "INFO",
	LogLevelCaution: "CAUTION",
	LogLevelWarning: "WARN",
	LogLevelError:   "ERROR",
}

type Log struct {
	Level     uint
	Message   string
	Operation string
}

func NewLog(level uint, message string, operation string) *Log {
	return &Log{Level: level, Message: message, Operation: operation}
}

type Monitor interface {
	GetOperation() string
	LogInfo(string)
	LogCaution(string)
	LogWarning(string)
	LogError(error)
}

type PrintMonitor struct {
	Operation string
}

func (pm PrintMonitor) GetOperation() string {
	return pm.Operation
}

func NewPrintMonitor(operation string) *PrintMonitor {
	return &PrintMonitor{Operation: operation}
}

func (pm PrintMonitor) LogInfo(msg string) {
	pm.log(LogLevelInfo, msg, pm.Operation)
}

func (pm PrintMonitor) LogCaution(msg string) {
	pm.log(LogLevelCaution, msg, pm.Operation)
}

func (pm PrintMonitor) LogWarning(msg string) {
	pm.log(LogLevelWarning, msg, pm.Operation)
}

func (pm PrintMonitor) LogError(err error) {
	pm.log(LogLevelError, err.Error(), pm.Operation)
}

func (pm PrintMonitor) log(level uint, message string, operation string) {
	logLine := fmt.Sprintf("[%s] %s: %s", LogLevelToString[level], operation, message)
	log.Println(logLine)
}
