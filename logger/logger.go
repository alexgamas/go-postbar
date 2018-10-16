package logger

import (
	"os"

	"github.com/op/go-logging"
)

// Log ...
var Log = logging.MustGetLogger("postbar")

func init() {

	var format = logging.MustStringFormatter(
		`%{color} %{time:2006-01-02T15:04:05.000} %{longfile} [%{level}] %{color:reset}%{message}`,
	)

	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	//backendLeveled := logging.AddModuleLevel(backend)
	//backendLeveled.SetLevel(logging.ERROR, "")
	//logging.SetBackend(backendLeveled, backendFormatter)
	logging.SetBackend(backendFormatter)

}
