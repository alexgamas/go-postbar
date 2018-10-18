package logger

import (
	"fmt"
	"os"

	"github.com/op/go-logging"
)

// LogWriter ...
type LogWriter struct {
}

func (l LogWriter) Write(p []byte) (n int, err error) {
	Log.Debug(fmt.Sprintf("%s", p))
	return len(p), nil
}

// Log ...
var Log = logging.MustGetLogger("postbar")

// DebugLogWriter ...
var DebugLogWriter = LogWriter{}

// https://astaxie.gitbooks.io/build-web-application-with-golang/en/images/2.3.init.png?raw=true
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
