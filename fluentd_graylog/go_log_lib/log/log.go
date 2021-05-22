package log

import (
	"io"
	glog "log"
	"strings"
	"sync"

	fluentd "github.com/joonix/log"
	"github.com/sirupsen/logrus"
)

var (
	// the default logger, which is updated by Initialize().
	defaultLog Log = NewJson(false)

	// mutex protecting default log
	mu sync.RWMutex
)

// LogCloser is a Log interface that must be closed in order to flush
// logs on shutdown.
type LogCloser interface {
	Log
	io.Closer
}

type logrusLog struct {
	*logrus.Logger
}

// Log represents the standard logger interface.
type Log interface {
	logrus.FieldLogger
}

func (l *logrusLog) Close() error {
	return nil
}

// SetDefaultLog sets the given Log instance as the default, meaning:
// * Standard logging will be redirected to the given logger.
// * Exported log functions will use the given logger (deprecated).
func SetDefaultLog(lg Log) {
	if lg == nil {
		panic("SetDefaultLog received a nil log instance")
	}

	mu.Lock()
	defer mu.Unlock()

	redirectStdLog(lg)
	defaultLog = lg
}

func redirectStdLog(lg Log) {
	glog.SetFlags(glog.Lshortfile)
	glog.SetOutput(WriterFunc(func(p []byte) (int, error) {
		str := strings.TrimSpace(string(p))
		lg.Info(str)
		return len(p), nil
	}))
}

func NewJson(debug bool) LogCloser {
	tf := &logrus.JSONFormatter{}

	return newLog(tf, debug)
}

func NewFluentd(debug bool) LogCloser {
	tf := &fluentd.FluentdFormatter{}

	return newLog(tf, debug)
}

func newLog(formatter logrus.Formatter, debug bool) LogCloser {
	l := logrus.New()
	l.Formatter = formatter

	l.Level = logrus.InfoLevel
	if debug {
		l.Level = logrus.DebugLevel
	}

	return &logrusLog{
		Logger: l,
	}
}

// WriterFunc allows functions to implement the io.Write interface.
type WriterFunc func(p []byte) (n int, err error)

func (f WriterFunc) Write(p []byte) (n int, err error) {
	return f(p)
}
