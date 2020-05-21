package debuglog

import "log"

type Logger interface {
	Println(...interface{})
}

type wrapper struct {
	logger Logger
}

func (l wrapper) Write(in []byte) (int, error) {
	l.logger.Println(string(in))
	return len(in), nil
}

func SetDebugLogger(in Logger) {
	out.mutex.Lock()
	defer out.mutex.Unlock()

	out.output = log.New(wrapper{logger: in}, "sumologic", log.Flags())
}
