package pkg

import (
	"errors"
	"notes/logger/logrus/logger"
)

func Test2() {

	err := errors.New("errors Test2")
	logger.Log.Error(err)

}
