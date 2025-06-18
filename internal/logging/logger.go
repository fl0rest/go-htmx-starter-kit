package logging

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"project/internal/config"
)

type Logger struct {
	sys   *log.Logger
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
}

var Log *Logger

func init() {
	cfg := config.Load()
	var sysWriter, infoWriter, warnWriter, errorWriter io.Writer

	if cfg.LogFile != "" {
		file, err := os.OpenFile(cfg.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			if os.IsNotExist(err) {
				log.Println("[SYS] Log dir not found, creating:", filepath.Dir(cfg.LogFile))

				if err := os.MkdirAll(filepath.Dir(cfg.LogFile), 0755); err != nil {
					log.Println("[SYS] Failed to create dir:", err)
				}

				file, err = os.OpenFile(cfg.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
				if err != nil {
					log.Println("[SYS] Failed to open logfile after creating dir:", err)
				}
			} else {
				log.Println("[SYS] Failed to open logfile: ", err)
			}
		}

		errorFile, err := os.OpenFile(cfg.ErrorLogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			if os.IsNotExist(err) {
				log.Println("[SYS] Log dir not found, creating:", filepath.Dir(cfg.LogFile))

				if err := os.MkdirAll(filepath.Dir(cfg.LogFile), 0755); err != nil {
					log.Println("[SYS] Failed to create dir:", err)
				}

				file, err = os.OpenFile(cfg.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
				if err != nil {
					log.Println("[SYS] Failed to open error logfile after creating dir:", err)
				}
			} else {
				log.Println("[SYS] Failed to open error logfile: ", err)
			}
		}

		if cfg.LogToScreen {
			sysWriter = io.MultiWriter(os.Stdout, file)
			infoWriter = io.MultiWriter(os.Stdout, file)
			warnWriter = io.MultiWriter(os.Stdout, file)
			errorWriter = io.MultiWriter(os.Stdout, errorFile)
		} else {
			sysWriter = file
			infoWriter = file
			warnWriter = file
			errorWriter = errorFile
		}

	} else {
		sysWriter = os.Stdout
		infoWriter = os.Stdout
		warnWriter = os.Stdout
		errorWriter = os.Stderr
	}

	Log = &Logger{
		sys:   log.New(sysWriter, "[SYS] ", log.LstdFlags),
		info:  log.New(infoWriter, "[INFO] ", log.LstdFlags),
		warn:  log.New(warnWriter, "[WARN] ", log.LstdFlags),
		error: log.New(errorWriter, "[ERROR] ", log.LstdFlags),
	}
}

func (l *Logger) Sys(v ...interface{}) {
	l.sys.Println(v...)
}

func (l *Logger) Sysf(format string, v ...interface{}) {
	l.sys.Printf(format, v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.warn.Println(v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warn.Printf(format, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.error.Println(v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.error.Printf(format, v...)
}
