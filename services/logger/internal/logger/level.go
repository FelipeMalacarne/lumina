package logger

type Level string

const (
	INFO  Level = "INFO"
	WARN  Level = "WARN"
	ERROR Level = "ERROR"
	DEBUG Level = "DEBUG"
	FATAL Level = "FATAL"
)
