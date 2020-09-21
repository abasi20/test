package main

import (
	seelog "github.com/cihub/seelog"
)

func writ() {
	logger.Debug("dddddd")
	// logger.traceWithCallDepth(2, "ssss")
	logger.Error("abc")
}

var logger, _ = seelog.LoggerFromConfigAsFile("seelog.xml")

func main() {
	// logger, err := seelog.LoggerFromConfigAsFile("seelog.xml")
	// defer seelog.Flush()
	// if err != nil {
	// 	seelog.Critical("err parsing config log file", err)
	// 	return
	// }
	defer seelog.Flush()
	seelog.ReplaceLogger(logger)

	seelog.Error("seelog error")
	seelog.Info("seelog info")
	seelog.Debug("seelog debug")
	writ()
}
