package logging

import (
	"github.com/sirupsen/logrus"
	"sync"
)

// ToDo: заменить Singleton на передачу логгера везде где он понадобится?
var instance *logrus.Entry = nil
var once sync.Once

func getInstance() *logrus.Entry {
	once.Do(func() {
		l := logrus.New()
		// ToDo: настроить форматтер для вывода правильного места логирования
		l.SetFormatter(&logrus.JSONFormatter{})
		instance = logrus.NewEntry(l)
	})
	return instance
}

func SetLevel(level string) {
	parsedlevel, err := logrus.ParseLevel(level)
	if err != nil {
		return
	}
	getInstance().Logger.SetLevel(parsedlevel)
}

func Trace(args ...interface{}) {
	getInstance().Trace(args)
}

func Debug(args ...interface{}) {
	getInstance().Debug(args)
}

func Info(args ...interface{}) {
	getInstance().Info(args)
}

func Warn(args ...interface{}) {
	getInstance().Warn(args)
}

func Error(args ...interface{}) {
	getInstance().Error(args)
}

func Fatal(args ...interface{}) {
	getInstance().Fatal(args)
}

func Panic(args ...interface{}) {
	getInstance().Panic(args)
}
