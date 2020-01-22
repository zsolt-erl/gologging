package loglib

import "github.com/jinzhu/copier"

// Logger implements the CommonLog interface
type Logger struct {
	Wrap Wrapper
	KV   Pairs
}

// Info is the implementations of the CommonLog Info function
func (l Logger) Info(args ...interface{}) {
	l.Wrap.Info(l.KV, args...)
}

// With implements With from the CommomLog interface
func (l Logger) With(key string, value string) Logger {
	newLogger := Logger{KV: Pairs{}}
	copier.Copy(&l, &newLogger)
	newLogger.KV[key] = value
	return newLogger
}
