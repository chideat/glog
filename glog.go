package glog

import (
	"fmt"
	logger "log"
	"os"
	"path"
	"runtime"
	"sync"
)

const (
	DEFAULT_LOG_PATH string = "logs"
)

type LEVEL uint8

const (
	INFO  = 1
	WARN  = 2
	ERROR = 4
	PANIC = 8
)

func (l LEVEL) String() string {
	switch l {
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case PANIC:
		return "PANIC"
	default:
		return "UNKNOWN"
	}
}

var (
	debug       bool
	level       LEVEL
	logFilePath string
	opendFile   *os.File

	lock *sync.RWMutex
)

func SetLevel(l LEVEL) {
	lock.RLock()
	level = l
	lock.RUnlock()
}

func SetDebug(d bool) {
	lock.Lock()
	defer func() {
		debug = d
		lock.Unlock()
	}()

	var err error
	if !d && logFilePath != "" {
		opendFile, err = os.OpenFile(logFilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err == nil {
			logger.SetOutput(opendFile)
		} else {
			panic(err)
		}
	} else {
		logger.SetOutput(os.Stdout)

		if opendFile != nil {
			if err = opendFile.Close(); err != nil {
				Error(err)
			}
		}
	}
}

func init() {
	lock = &sync.RWMutex{}

	info, err := os.Stat(DEFAULT_LOG_PATH)
	if err == nil && info.IsDir() {
		logFilePath = path.Join(DEFAULT_LOG_PATH, fmt.Sprintf("%s.log", os.Args[0]))
	} else {
		logFilePath = fmt.Sprintf("%s.log", os.Args[0])
	}

	SetLevel(INFO | WARN | ERROR | PANIC)
	SetDebug(os.Getenv("DEBUG") == "debug")
}

func print(level LEVEL, data string) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 1
	}
	logger.Printf("[%s] %s:%d %s\n", level.String(), file, line, data)
}

func Info(v ...interface{}) {
	lock.RLock()
	defer lock.RUnlock()

	if INFO&level == INFO {
		print(INFO, fmt.Sprintf("%v", v...))
	}
}

func Infof(format string, v ...interface{}) {
	lock.RLock()
	defer lock.RUnlock()

	if INFO&level == INFO {
		print(INFO, fmt.Sprintf(format, v...))
	}
}

func Warn(v ...interface{}) {
	lock.RLock()
	defer lock.RUnlock()

	if WARN&level == WARN {
		print(WARN, fmt.Sprintf("%v", v...))
	}
}

func Warnf(format string, v ...interface{}) {
	lock.RLock()
	defer lock.RUnlock()

	if WARN&level == WARN {
		print(WARN, fmt.Sprintf(format, v...))
	}
}

func Error(v ...interface{}) {
	lock.RLock()
	defer lock.RUnlock()

	if ERROR&level == ERROR {
		print(ERROR, fmt.Sprintf("%v", v...))
	}
}

func Errorf(format string, v ...interface{}) {
	lock.RLock()
	defer lock.RUnlock()

	if ERROR&level == ERROR {
		print(ERROR, fmt.Sprintf(format, v...))
	}
}

func Panic(v ...interface{}) {
	lock.RLock()
	defer lock.RUnlock()

	if PANIC&level == PANIC {
		print(PANIC, fmt.Sprintf("%v", v...))
		panic(fmt.Sprintf("%v", v...))
	}
}

func Panicf(format string, v ...interface{}) {
	lock.RLock()
	defer lock.RUnlock()

	if PANIC&level == PANIC {
		print(PANIC, fmt.Sprintf(format, v...))
		panic(fmt.Sprintf(format, v...))
	}
}
