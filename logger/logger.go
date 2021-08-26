package logger

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/battulasaivinesh/CustomLog/common"
	"github.com/battulasaivinesh/CustomLog/common/constants"
	"github.com/battulasaivinesh/CustomLog/common/types"
)

type Logger struct {
	Output        types.OutputType
	FileName      string
	FilePath      string
	FileSizeLimit int64
}

var LoggerInstance *Logger

func NewLoggerInstance(config types.LoggerConfig) *Logger {
	LoggerInstance = &Logger{
		Output:        config.OutputType,
		FileName:      config.FileName,
		FileSizeLimit: config.FileSizeLimit,
	}
	LoggerInstance.UpdateCurrentFile()
	return LoggerInstance
}

func (log Logger) PrintLog(data types.LogData) {
	switch log.Output {
	case constants.OutputTypes.Console:
		common.LogToConsole(data)
	case constants.OutputTypes.File:
		log.CheckAndUpdateCurrentFile()
		common.LogToFile(data, log.FilePath)
	}
}

func (log Logger) Warnf(desc string) {
	log.PrintLog(types.LogData{
		LogType:     constants.LogTypes.Warning,
		Description: desc,
	})
}

func (log Logger) Infof(desc string) {
	log.PrintLog(types.LogData{
		LogType:     constants.LogTypes.Info,
		Description: desc,
	})
}

func (log Logger) Errorf(desc string) {
	log.PrintLog(types.LogData{
		LogType:     constants.LogTypes.Error,
		Description: desc,
	})
}

func (log *Logger) UpdateCurrentFile() {
	absPath, err := filepath.Abs("../")
	if err != nil {
		panic(err)
	}

	root := absPath + "/logs/"
	var files []string
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if path != root {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	currentFile := ""
	count := 0
	for _, file := range files {
		info, err := os.Stat(file)
		if err != nil {
			continue
		}
		x := info.Size()
		if x < log.FileSizeLimit {
			currentFile = file
		}
		count++
	}

	var file string
	if currentFile != "" {
		file = currentFile
	} else {
		file = absPath + "/logs/" + log.FileName + "." + strconv.Itoa(count)
		_, err = os.Create(file)
		if err != nil {
			panic(err)
		}
	}

	log.FilePath = file
}

func (log *Logger) CheckAndUpdateCurrentFile() {
	info, err := os.Stat(log.FilePath)
	if err != nil {
		panic(err)
	}
	if info.Size() >= log.FileSizeLimit {
		log.UpdateCurrentFile()
	}
}
