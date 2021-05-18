package utils

import (
	"github.com/sirupsen/logrus"
)

var defaultEntry = logrus.NewEntry(logrus.StandardLogger())

// log2LogrusWriter exploits the documented fact that the standard
// log pkg sends each log entry as a single io.Writer.Write call:
// https://golang.org/pkg/log/#Logger
type log2LogrusWriter struct {
	entry *logrus.Entry
}

func (w *log2LogrusWriter) Write(b []byte) (int, error) {
	n := len(b)
	if n > 0 && b[n-1] == '\n' {
		b = b[:n-1]
	}
	w.entry.Info(string(b))
	return n, nil
}
