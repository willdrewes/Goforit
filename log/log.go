package log

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
)

func Init(env, filename string) error {
	if len(env) == 0 || env == "dev" {
		formater := &logrus.TextFormatter{}
		formater.ForceColors = true
		logrus.SetFormatter(formater)
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			return err
		}

		logrus.SetOutput(f)
	}

	return nil
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

func Error(err error, args ...interface{}) {
	logrus.WithField("error", err).Error(args...)
}

func ErrorMessage(args ...interface{}) {
	logrus.Error(args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args)
}

func LogExternalHTTP(ctx context.Context, URI string, start time.Time, method string, statusCode int) {
	requestURL, _ := url.Parse(URI)
	statusText := http.StatusText(statusCode)
	if statusCode == 422 {
		statusText = "unprocessable entity"
	}
	msg := fmt.Sprintf("External %s %s %v %s in %v", method, requestURL.RequestURI(), statusCode, statusText, time.Since(start))
	if statusCode >= 400 {
		ErrorMessage(msg)
	} else {
		Info(msg)
	}
}
