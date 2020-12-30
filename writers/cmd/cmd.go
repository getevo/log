package cmd

import (
	"fmt"
	"github.com/getevo/log/logger"
	"github.com/gookit/color"
)


//     %pid       	Process id (int)
//     %unix      	Time when log occurred (time.Time)
//     %ts	      	Time when log occurred (time.Time)
//     %level  		Log level (Level)
//     %package		Module (string)
//     %program     Basename of os.Args[0] (string)
//     %file	    Basename of file
//     %message	    Message (string)
//     %fn  		Function name
//     %path		Path
//     %line		Line
//     %color[]     Color


type Cmd struct {
	Colorful bool
	Format   string
	DateFormat string
}

func New() Cmd {
	return Cmd{
		Colorful: true,
		Format: "%color[%t][%lvl][%fl:%ln pkg:%pkg]:%white %msg",
		DateFormat: "2006-01-02 15:04:05",
	}
}

func (f Cmd)Write(trace *logger.Trace) {

	state := 0
	arg := ""
	var print = color.White.Print
	for _,chr := range f.Format{
		if state == 0 && chr != '%'{
			print(string(chr))
			continue
		}else if state == 0 && chr == '%'{
			state = 1
			continue
		}else if state == 1{
			if chr >= 'a' && chr <= 'z'{
				arg += string(chr)
				continue
			}else{

				print = f.renderArg(print,arg,trace)
				print(string(chr))
				arg = ""
				state = 0
				continue
			}
		}
	}
	if arg != ""{
		 f.renderArg(print,arg,trace)
	}
	fmt.Print("\n")
	//fmt.Println( colorRegex.ReplaceAllString(f.Format,genColor(trace.Level) ) )
}

func (f *Cmd)renderArg(p func(args ...interface{}), arg string, trace *logger.Trace) func(args ...interface{}) {
	switch arg {
		case "ts":
			p(trace.Time.Unix())
		case "t":
			p(trace.Time.Format(f.DateFormat))
		case "lvl":
			p(trace.Level.String())
		case "msg":
			p(trace.Message)
		case "pkg":
			p(trace.Package)
		case "fn":
			p(trace.Function)
		case "fl":
			p(trace.File)
		case "ln":
			p(trace.Line)
		case "white","reset":
			return color.White.Print
		case "red":
			return color.Red.Print
		case "cyan":
			return color.Cyan.Print
		case "green":
			return color.Green.Print
		case "yellow":
			return color.Yellow.Print
		case "color":
			if trace.Level == logger.DEBUG || trace.Level == logger.WARNING{
				return color.Debug.Print
			}else if trace.Level == logger.ERROR || trace.Level == logger.CRITICAL{
				return color.Error.Print
			}else if trace.Level == logger.NOTICE{
				return color.Notice.Print
			}else if trace.Level == logger.INFO{
				return color.Info.Print
			}else{
				return color.White.Print
			}
	}
	return p
}



func genColor(level logger.Level) string {
	color.Debug.Print()
	if level == logger.DEBUG || level == logger.WARNING{
		return color.Debug.Render("$1")
	}else if level == logger.ERROR || level == logger.CRITICAL{
			return color.Error.Render("$1")
	}else if level == logger.NOTICE{
		return color.Notice.Render("$1")
	}else if level == logger.INFO{
		return color.Info.Render("$1")
	}
	return color.Debug.Render("$1")
}
