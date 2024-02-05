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
	LogLevelAlert
)

var LogLevelToString = map[uint]string{
	LogLevelInfo:    "INFO",
	LogLevelCaution: "CAUTION",
	LogLevelWarning: "WARN",
	LogLevelError:   "ERROR",
	LogLevelAlert:   "ALERT",
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
	LogError(string)
	LogAlert(string)
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

func (pm PrintMonitor) LogError(msg string) {
	pm.log(LogLevelError, msg, pm.Operation)
}

func (pm PrintMonitor) LogAlert(msg string) {
	pm.log(LogLevelAlert, msg, pm.Operation)
}

func LogByLevel(monitor Monitor, logLevel uint, msg string) {
	switch logLevel {
	case LogLevelInfo:
		monitor.LogInfo(msg)
	case LogLevelCaution:
		monitor.LogCaution(msg)
	case LogLevelWarning:
		monitor.LogWarning(msg)
	case LogLevelError:
		monitor.LogError(msg)
	case LogLevelAlert:
		monitor.LogAlert(msg)
	default:
		monitor.LogAlert(fmt.Sprintf("Invalid logging level %d used by monitor at %s [msg='%s']", logLevel, monitor.GetOperation(), msg))
	}
}

func (pm PrintMonitor) log(level uint, message string, operation string) {
	logLine := fmt.Sprintf("[%s] %s: %s", LogLevelToString[level], operation, message)
	log.Println(logLine)
}
