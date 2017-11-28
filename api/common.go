package api

type LOG_LEVEL uint

const (
	LOG_DEBUG LOG_LEVEL = iota
	LOG_INFO
	LOG_WARNING
	LOG_ERROR
	LOG_FATAL
)

type LoggerConfig struct {
	Filename string
	loglevel LOG_LEVEL
}
