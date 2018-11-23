package logs

import (
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

var Log *log.Logger

func PrintLogger() *log.Logger  {
	if Log != nil {
		return Log
	}

	Log = log.New()
	Log.SetLevel(log.InfoLevel)
	formatter := &log.TextFormatter{
		TimestampFormat:"2006-01-01 20:00:00",
		FullTimestamp:true,
	}

	Log.SetReportCaller(true)
	pathMap := lfshook.PathMap{
		log.InfoLevel : "./logs/info.log",
		log.DebugLevel : "./logs/debug.log",
		log.ErrorLevel : "./logs/error.log",
		log.FatalLevel : "./logs/fatal.log",
	}

	Log.Hooks.Add(lfshook.NewHook(
		pathMap,
		formatter,
		))

	return Log
}