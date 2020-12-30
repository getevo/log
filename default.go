package log

import (
	"github.com/getevo/log/logger"
	"github.com/getevo/log/writers/stdio"
)

var DefaultLogger = logger.Logger{
	Option: logger.Option{
		PrintLevel: logger.DEBUG,
		WriteLevel: logger.ERROR,
		Writers:    []logger.Writer{cmd.New() },
		//Writers:    []logger.Writer{stdio.New() },
	},
}
func Warning(input ...interface{}) {
	DefaultLogger.Warning(input...)
}

func Critical(input ...interface{}) {
	DefaultLogger.Critical(input...)
}

func Error(input ...interface{}) {
	DefaultLogger.Error(input...)
}

func Debug(input ...interface{}) {
	DefaultLogger.Debug(input...)
}

func Notice(input ...interface{}) {
	DefaultLogger.Notice(input...)
}

func Info(input ...interface{}) {
	DefaultLogger.Info(input...)
}

func Print(input ...interface{}) {
	DefaultLogger.Print(input...)
}
