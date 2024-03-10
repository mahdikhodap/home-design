package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const serviceName = "home"

func ErrorLogger(funcName string) *zerolog.Event {
	return setLogger(log.Error(), funcName)
}

func InfoLogger(funcName string) *zerolog.Event {
	return setLogger(log.Info(), funcName)
}

func DebugLogger(funcName string) *zerolog.Event {
	return setLogger(log.Debug(), funcName)
}

func setLogger(event *zerolog.Event, funcName string) *zerolog.Event {
	return event.Str("service", serviceName).Str("function", funcName)
}
