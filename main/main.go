package main

import (
	"strconv"

	"github.com/battulasaivinesh/CustomLog/common/constants"
	"github.com/battulasaivinesh/CustomLog/common/types"
	"github.com/battulasaivinesh/CustomLog/logger"
)

func main() {
	logCust := logger.NewLoggerInstance(types.LoggerConfig{
		OutputType:    constants.OutputTypes.File,
		FileName:      "test.log",
		FileSizeLimit: 1024,
	})

	for i := 0; i < 100; i++ {
		logCust.PrintLog(types.LogData{
			LogType:     constants.LogTypes.Error,
			SubService:  "Mail" + strconv.Itoa(i),
			Method:      "SendMail",
			Description: "Invalid Email Id",
		})
	}

	for i := 0; i < 10; i++ {
		logCust.Warnf("Test Warning")
		logCust.Infof("Test Info")
		logCust.Errorf("Test Error")
	}
}
