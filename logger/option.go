package logger


type Option struct {
	PrintLevel Level
	WriteLevel Level
	Writers    []Writer
}

var DefaultOption = Option{
	PrintLevel: DEBUG,
	WriteLevel: ERROR,

}

