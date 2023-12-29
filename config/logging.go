package config

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
)

func Logger() *log.Logger {

    logger:= log.NewWithOptions(os.Stderr, log.Options{
        TimeFunction: time.Now().Local().UTC, //.UTC, // 2023-11-26T02:21:15Z with RFC3339
                                      //time.Now().Local 2023-11-25T23:19:43-03:00
        TimeFormat: time.RFC3339, // this shows a Z instead of 00:00 in utc
        // TimeFormat: time.DateTime, // this is good for dev,,, not for prod
        ReportCaller: true,
        ReportTimestamp: true,
        Prefix: "log",
        Level: log.DebugLevel,
        // Formatter: log.JSONFormatter,
    })
    return logger
}
