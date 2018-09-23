package logger

import (
	"os"

	"github.com/op/go-logging"
)

// type Password string

// func (p Password) Redacted() interface{} {
// 	return logging.Redact(string(p))
// }

// func Info(args ...interface{}) {
// 	log.Info(args)
// 	/*
// 		log.Debugf("debug %s", Password("secret"))
// 		log.Notice("notice")
// 		log.Error("err")
// 		log.Critical("crit")
// 	*/
// }

// func Fatal(args ...interface{}) {
// 	log.Fatal(args)
// 	/*
// 		log.Debugf("debug %s", Password("secret"))
// 		log.Notice("notice")
// 		log.Error("err")
// 		log.Critical("crit")
// 	*/
// }

// func Warning(msg string) {
// 	log.Warning(msg)
// }

// func Debug(msg string) {
// 	log.Debug(msg)
// }

//func Debug(string msg)

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
	/*

		var logpath = build.Default.GOPATH + "/src/chat/logger/info.log"

		flag.Parse()
		var file, err1 = os.Create(logpath)

		if err1 != nil {
			panic(err1)
		}

		Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
		Log.Println("LogFile : " + logpath)
	*/
}
