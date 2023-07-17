package shared

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func LogError(message, directory, function string, error error, args ...any) {
	logrus.WithFields(logrus.Fields{
		"directory": directory,
		"function":  function,
		"args":      fmt.Sprintf("%+v", args),
		"error":     error,
	}).Error(message)
}

func LogWarn(message, directory, function string, error error, args ...any) {
	logrus.WithFields(logrus.Fields{
		"directory": directory,
		"function":  function,
		"args":      fmt.Sprintf("%+v", args),
		"error":     error,
	}).Warning(message)
}

func LogInfo(message, directory, function string, error error, args ...any) {
	logrus.WithFields(logrus.Fields{
		"directory": directory,
		"function":  function,
		"args":      fmt.Sprintf("%+v", args),
		"error":     error,
	}).Info(message)
}

func LogRequest(uuid, method, url, body string, request, headers any) {
	logrus.WithFields(logrus.Fields{
		"method":  method,
		"url":     url,
		"body":    body,
		"headers": headers,
	}).Infof("HTTP_REQUEST=%s", uuid)
}

func LogResponse(uuid, status, body, method, url string, headers any) {
	logrus.WithFields(logrus.Fields{
		"method":  method,
		"url":     url,
		"status":  status,
		"body":    body,
		"headers": headers,
	}).Infof("HTTP_RESPONSE=%s", uuid)
}
