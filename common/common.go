package common

import (
	"fmt"
	"os"
	"time"

	"github.com/battulasaivinesh/CustomLog/common/types"
)

func LogToConsole(data types.LogData) {
	fmt.Printf(getLogString(data))
}

func getLogString(data types.LogData) string {
	var log string
	timeNow := time.Now()
	log += timeNow.Format("2006-01-02 15:04:05 - ")
	log += "(" + string(data.LogType) + ")" + " "
	if data.SubService != "" {
		log += "SubService - " + data.SubService + ", "
	}

	if data.Method != "" {
		log += "Method - " + data.Method + ", "
	}

	log += data.Description + "\n"
	return log
}

func LogToFile(data types.LogData, file string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(getLogString(data)); err != nil {
		panic(err)
	}
}
