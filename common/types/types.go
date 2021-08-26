package types

type LogType string

type OutputType int

type LogData struct {
	LogType     LogType
	SubService  string
	Method      string
	Description string
}

type LoggerConfig struct {
	OutputType    OutputType
	FileName      string
	FileSizeLimit int64
}
