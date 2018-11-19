package asset

import (
	os "os"

	"fmt"

	logging "github.com/op/go-logging"
)

var logger *logging.Logger

var format = logging.MustStringFormatter(
	`%{message}`,
)

func init() {
	logFile, err := os.OpenFile("/opt/ytrace/log/log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("err reason : ", err.Error())
		panic(err)
	}
	backend1 := logging.NewLogBackend(logFile, "", 0)
	//backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	backend2Formatter := logging.NewBackendFormatter(backend1, format)
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.INFO, "")

	logging.SetBackend(backend1Leveled, backend2Formatter)

	logger = logging.MustGetLogger("asset")

}
