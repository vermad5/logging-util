package api

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

func (logInst *LoggerConfig) Debug(logMessage string, args ...interface{}) {
	log.Debug(logMessage, args)
}
func (logInst *LoggerConfig) Info(logMessage string, args ...interface{}) {
	log.Info(logMessage, args)
}
func (logInst *LoggerConfig) Warning(logMessage string, args ...interface{}) {
	log.Warning(logMessage, args)
}
func (logInst *LoggerConfig) Error(logMessage string, args ...interface{}) {
	log.Error(logMessage, args)
}
func (logInst *LoggerConfig) Fatal(logMessage string, args ...interface{}) {
	log.Fatal(logMessage, args)
}

func (logInst *LoggerConfig) Debugf(logMessage string, args ...interface{}) {
	log.Debugf(logMessage, args)
}
func (logInst *LoggerConfig) Infof(logMessage string, args ...interface{}) {
	log.Infof(logMessage, args)
}
func (logInst *LoggerConfig) Warningf(logMessage string, args ...interface{}) {
	log.Warningf(logMessage, args)
}
func (logInst *LoggerConfig) Errorf(logMessage string, args ...interface{}) {
	log.Errorf(logMessage, args)
}
func (logInst *LoggerConfig) Fatalf(logMessage string, args ...interface{}) {
	log.Fatalf(logMessage, args)
}

func (logInst *LoggerConfig) Create() {
	Formatter := new(log.TextFormatter)
	Formatter.FullTimestamp = true

	log.SetFormatter(Formatter)
	setLogLevel(os.Getenv("FRACTAL_LOG_LEVEL"))

	// Create the log file if doesn't exist. And append to it if it already exists.
	if logInst.Filename == "" {
		logInst.Warning("Logger file name is empty! Defaulting to fractal.log")
		logInst.Filename = "fractal-web.log"
	} else {
		file, err := os.OpenFile(logInst.Filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			logInst.Warning(fmt.Sprintf("Unable to create file [%s] Error Description:[%s]. Defaulting to console logging ", logInst.Filename, err))
		} else {
			log.SetOutput(file)
			logInst.Info("Successfully configured logger")
		}
	}
}

func setLogLevel(level string) {
	switch level {
	case "INFO":
		log.SetLevel(log.InfoLevel)
		break
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
		break
	case "WARNING":
		log.SetLevel(log.WarnLevel)
		break
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
		break
	case "FATAL":
		log.SetLevel(log.FatalLevel)
		break
	default:
		log.SetLevel(log.InfoLevel)
	}
}
