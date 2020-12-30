package logger

import (
	"fmt"
	"github.com/alecthomas/repr"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type Trace struct {
	Time     time.Time
	Package  string
	Function string
	Line     int
	File     string
	Path     string
	Level    Level
	Message  string
}

type Logger struct {
	Option Option
}

func Sprint(input ...interface{}) string {
	var out = ""
	if len(input) == 0{
		return out
	}
	if v,ok := input[0].(string); ok{
		if len(input) == 1{
			return fmt.Sprint(input[0])
		}else {
			return fmt.Sprintf(v, input[1:])
		}
	}
	for _,in := range input{
		if v,ok := in.(error); ok{
			out += v.Error() + "\r\n"
		}else {
			out += repr.String(in) + "\r\n"
		}
	}
	return	strings.TrimSpace(out)

}


func (logger Logger)Warning(input ...interface{}) {
	msg := Sprint(input...)
	logger.Write(msg, WARNING)
}

func (logger Logger)Critical(input ...interface{}) {
	msg := Sprint(input...)
	logger.Write(msg, CRITICAL)
}

func (logger Logger)Error(input ...interface{}) {
	msg := Sprint(input...)
	logger.Write(msg, ERROR)
}

func (logger Logger)Debug(input ...interface{}) {
	msg := Sprint(input...)
	logger.Write(msg, DEBUG)
}

func (logger Logger)Notice(input ...interface{}) {
	msg := Sprint(input...)
	logger.Write(msg, NOTICE)
}

func (logger Logger)Info(input ...interface{}) {
	msg := Sprint(input...)
	logger.Write(msg, INFO)
}

func (logger Logger)Print(input ...interface{}) {
	msg := Sprint(input...)
	fmt.Println(msg)
}

func (logger Logger)Write(msg string, level Level) {

	pc, file, line, _ := runtime.Caller(3)
	fn := runtime.FuncForPC(pc)
	chunks := strings.Split(fn.Name(),".")

	var trace = Trace{
		Time: time.Now(),
		Line: line,
		File: filepath.Base(file),
		Path: filepath.Dir(file),
		Package: chunks[0],
		Function:  chunks[1],
		Level: level,
		Message: msg,
	}
	for _,writer := range logger.Option.Writers {
		writer.Write(&trace)
	}

}