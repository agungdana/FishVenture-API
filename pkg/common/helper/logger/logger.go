package logger

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/sirupsen/logrus"
)

var (
	logger *Logger
	once   sync.Once
)

func extractLabel(ctx context.Context) (label string) {
	if reqId, ok := ctxutil.GetRequestID(ctx); ok {
		label += "[R: " + reqId.String() + " ]"
	}
	return label + " "
}

type Logger struct {
	Log   *logrus.Logger
	Entry *logrus.Entry
}

type loggerHook struct {
	level    logrus.Level
	formater logrus.Formatter
}

func (l *loggerHook) Levels() []logrus.Level {
	return logrus.AllLevels[:l.level]
}

func (l *loggerHook) Fire(entry *logrus.Entry) error {
	now := time.Now()
	years := now.Year()
	month := now.Month()
	day := now.Day()
	dateString := now.Format("2006-01-02")

	path := fmt.Sprintf("%v/%v/%v/%v", "Log", years, month, day)

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	fileName := fmt.Sprintf("%v/%v(%v).json", path, "log", dateString)

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	save, err := l.formater.Format(entry)
	if err != nil {
		return err
	}
	_, err = file.Write(save)
	if err != nil {
		return err
	}
	return nil
}

func SetupLogger(debugStatus string) {
	if logger == nil {
		once.Do(func() {
			log := logrus.New()
			log.SetLevel(logrus.InfoLevel)
			if debugStatus == "true" || debugStatus == "" {
				log.SetLevel(logrus.TraceLevel)
			}

			entry := logrus.NewEntry(log)

			logger = &Logger{
				Log:   log,
				Entry: entry,
			}
		})
	}

	logger.Log.AddHook(&loggerHook{
		level: logger.Log.GetLevel(),
		formater: &logrus.JSONFormatter{
			TimestampFormat:   "2006-02-01",
			DisableHTMLEscape: true,
			PrettyPrint:       true,
		},
	})

}

func GetLogger() *Logger {
	return logger
}

func Trace(format string, msg ...any) {
	messages := fmt.Sprintf("Trace:"+format, msg...)
	go logger.Entry.Trace(messages)
}

func Debug(format string, msg ...any) {
	messages := fmt.Sprintf("Debug:"+format, msg...)
	go logger.Entry.Debug(messages)
}

func Info(format string, msg ...any) {
	messages := fmt.Sprintf("Info:"+format, msg...)
	go logger.Entry.Info(messages)
}

func Error(format string, msg ...any) {
	messages := fmt.Sprintf("Error:"+format, msg...)
	go logger.Entry.Error(messages)
}

func Warn(format string, msg ...any) {
	messages := fmt.Sprintf("Warn:"+format, msg...)
	go logger.Entry.Warn(messages)
}

func Fatal(format string, msg ...any) {
	messages := fmt.Sprintf("Fatal:"+format, msg...)
	go logger.Entry.Fatal(messages)
}

func Panic(format string, msg ...any) {
	messages := fmt.Sprintf("Panic:"+format, msg...)
	go logger.Entry.Panic(messages)
}

func TraceWithContext(ctx context.Context, format string, msg ...any) {
	messages := fmt.Sprintf("Trace:"+extractLabel(ctx)+format, msg...)
	logger.Entry.Context = ctx
	go logger.Entry.Trace(messages)
}

func DebugWithContext(ctx context.Context, format string, msg ...any) {
	messages := fmt.Sprintf("Debug:"+extractLabel(ctx)+format, msg...)
	logger.Entry.Context = ctx
	go logger.Entry.Debug(messages)
}

func InfoWithContext(ctx context.Context, format string, msg ...any) {
	messages := fmt.Sprintf("Info:"+extractLabel(ctx)+format, msg...)
	logger.Entry.Context = ctx
	go logger.Entry.Info(messages)
}

func WarnWithContext(ctx context.Context, format string, msg ...any) {
	messages := fmt.Sprintf("Warn:"+extractLabel(ctx)+format, msg...)
	logger.Entry.Context = ctx
	go logger.Entry.Warn(messages)
}

func ErrorWithContext(ctx context.Context, format string, msg ...any) {
	messages := fmt.Sprintf("Error:"+extractLabel(ctx)+format, msg...)
	logger.Entry.Context = ctx
	go logger.Entry.Error(messages)
}

func FatalWithContext(ctx context.Context, format string, msg ...any) {
	messages := fmt.Sprintf("Fatal:"+extractLabel(ctx)+format, msg...)
	logger.Entry.Context = ctx
	go logger.Entry.Fatal(messages)
}

func PanicWithContext(ctx context.Context, format string, msg ...any) {
	messages := fmt.Sprintf("Panic:"+extractLabel(ctx)+format, msg...)
	logger.Entry.Context = ctx
	go logger.Entry.Panic(messages)
}
