package debuglog

import (
	"log"
	"sync"
)

type logger struct {
	debugEnabled bool
	output       *log.Logger
	mutex        sync.RWMutex
}

var out *logger

func init() {
	out = &logger{
		output: log.New(log.Writer(), "sumologic", log.Flags()),
	}
}

func Enable() {
	out.mutex.Lock()
	defer out.mutex.Unlock()
	out.debugEnabled = true
}

func SetLogger(log *log.Logger) {
	out.mutex.Lock()
	defer out.mutex.Unlock()

	out.output = log
}

func Log(in interface{}) {
	out.mutex.RLock()
	defer out.mutex.RUnlock()

	if !out.debugEnabled {
		return
	}

	out.output.Print(in)
}

func Logf(s string, args ...interface{}) {
	out.mutex.RLock()
	defer out.mutex.RUnlock()

	out.output.Printf(s, args...)
}

func Logln(args ...interface{}) {
	out.mutex.RLock()
	defer out.mutex.RUnlock()

	out.output.Println(args...)
}

func Debug(in interface{}) {
	out.mutex.RLock()
	defer out.mutex.RUnlock()

	out.output.Print(in)
}

func Debugf(s string, args ...interface{}) {
	out.mutex.RLock()
	defer out.mutex.RUnlock()

	if !out.debugEnabled {
		return
	}

	out.output.Printf(s, args...)
}

func Debugln(args ...interface{}) {
	out.mutex.RLock()
	defer out.mutex.RUnlock()

	if !out.debugEnabled {
		return
	}

	out.output.Println(args...)
}

func Fatal(err error) {
	if err == nil {
		return
	}

	out.mutex.RLock()
	defer out.mutex.RUnlock()
	out.output.Fatal(err)
}
