package loglib

// Pairs is key value pairs for structured logging
type Pairs map[string]string

// CommonLog is the interface that packages will require from a logger
type CommonLog interface {
	Info(kv Pairs, args ...interface{})
	With(key string, value string) CommonLog
}

// Wrapper is the interface that a log library wrapper will implement
type Wrapper interface {
	Info(Pairs, ...interface{})
}
