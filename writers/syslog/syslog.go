package syslog

import (
	"fmt"
	"github.com/getevo/log/logger"
	"io"
	"log/syslog"
	"os"
)

type Stdio struct {
	Stdout io.Writer
	Stderr io.Writer
	Format   string
	DateFormat string
}

func New() Stdio {
	syslog.Dial()
	return Stdio{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Format: "%color[%t][%lvl][%fl:%ln pkg:%pkg]:%white %msg",
		DateFormat: "2006-01-02 15:04:05",
	}
}

func (f Stdio)Write(trace *logger.Trace) {

	state := 0
	arg := ""
	var writer io.Writer
	if trace.Level == logger.ERROR || trace.Level == logger.CRITICAL{
		writer = f.Stderr
	}else{
		writer = f.Stdout
	}

	for _,chr := range f.Format{
		if state == 0 && chr != '%'{
			writer.Write([]byte( string(chr) ))
			continue
		}else if state == 0 && chr == '%'{
			state = 1
			continue
		}else if state == 1{
			if chr >= 'a' && chr <= 'z'{
				arg += string(chr)
				continue
			}else{

				f.renderArg(writer,arg,trace)
				writer.Write([]byte( string(chr) ))
				arg = ""
				state = 0
				continue
			}
		}
	}
	if arg != ""{
		f.renderArg(writer,arg,trace)
	}
	writer.Write([]byte( string("\n") ))


}

func (f *Stdio)renderArg(writer io.Writer,arg string, trace *logger.Trace){
	switch arg {
	case "ts":
		writer.Write( []byte(fmt.Sprint(trace.Time.Unix()) ))
	case "t":
		writer.Write( []byte(fmt.Sprint(trace.Time.Format(f.DateFormat))))
	case "lvl":
		writer.Write( []byte(trace.Level.String()))
	case "msg":
		writer.Write( []byte(trace.Message))
	case "pkg":
		writer.Write( []byte(trace.Package))
	case "fn":
		writer.Write( []byte(trace.Function))
	case "fl":
		writer.Write( []byte(trace.File))
	case "ln":
		writer.Write( []byte(fmt.Sprint(trace.Line) ))
	}
}