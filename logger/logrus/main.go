package main

import (
	"notes/logger/logrus/logger"
	"notes/logger/logrus/pkg"
)

func main() {
	logger.MakeLogger()
	pkg.Test2()

	logger.Log.Info("infotext")

}
