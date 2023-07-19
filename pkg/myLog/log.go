package myLog

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type logFormatter struct{}

func (t *logFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s\n", timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

var Log = NewLog()

func NewLog() *logrus.Logger {
	mLog := logrus.New()
	mLog.SetOutput(os.Stdout)
	mLog.SetReportCaller(true)
	mLog.SetFormatter(&logFormatter{})
	mLog.SetLevel(logrus.DebugLevel)
	return mLog
}
