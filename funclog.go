package funclog

import (
"log"
	"os"
	"path/filepath"
"runtime"
"strings"
)

type ErrorLogWriter struct{}

type InfoLogWriter struct{}

func (f ErrorLogWriter) Write(p []byte) (n int, err error) {
	log.SetOutput(os.Stderr)
	pc, file, line, ok := runtime.Caller(4)
	if !ok {
		file = "?"
		line = 0
	}

	fn := runtime.FuncForPC(pc)
	var fnName string
	if fn == nil {
		fnName = "?()"
	} else {
		dotName := filepath.Ext(fn.Name())
		fnName = strings.TrimLeft(dotName, ".") + "()"
	}

	log.Printf("%s:%d %s: %s", filepath.Base(file), line, fnName, p)
	return len(p), nil
}

func (f InfoLogWriter) Write(p []byte) (n int, err error) {
	log.SetOutput(os.Stdout)
	pc, file, line, ok := runtime.Caller(4)
	if !ok {
		file = "?"
		line = 0
	}

	fn := runtime.FuncForPC(pc)
	var fnName string
	if fn == nil {
		fnName = "?()"
	} else {
		dotName := filepath.Ext(fn.Name())
		fnName = strings.TrimLeft(dotName, ".") + "()"
	}

	log.Printf("%s:%d %s: %s", filepath.Base(file), line, fnName, p)
	return len(p), nil
}

func NewErrorLogger(prefix string) *log.Logger {
	return log.New(ErrorLogWriter{}, prefix, 0)
}

func NewInfoLogger(prefix string) *log.Logger {
	return log.New(InfoLogWriter{}, prefix, 0)
}