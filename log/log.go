package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog = log.New(os.Stdout, "\033[34m[info ]\033[0m", log.LstdFlags|log.Lshortfile)
	loggers = []*log.Logger{errorLog, infoLog}
	mu sync.Mutex
)

// log methods
var (
	Error = errorLog.Println
	Errorf = errorLog.Printf
	Info = infoLog.Println
	Infof = infoLog.Printf
)

// log level
const (
	Infolevel = iota
	ErrorLvel
	Disabled
)

// SetLevel control log level
func SetLevel(level int){
	mu.Lock()
	defer mu.Unlock()

	for _,logger := range loggers{
		logger.SetOutput(os.Stdout)
	}
	if ErrorLvel < level {
		errorLog.SetOutput(ioutil.Discard)
	}
	if Infolevel < level {
		errorLog.SetOutput(ioutil.Discard)
	}
}