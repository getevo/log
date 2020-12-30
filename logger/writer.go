package logger

type Writer interface {
	Write(trace *Trace)
}
