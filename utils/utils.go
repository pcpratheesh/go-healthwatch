package utils

import (
	"strings"

	"github.com/pcpratheesh/go-healthwatch/constants"
	"github.com/pcpratheesh/go-healthwatch/service"
	"github.com/pcpratheesh/go-healthwatch/utils/errors"
	"github.com/sirupsen/logrus"
)

var errNoScheme = errors.New("no scheme", "no scheme")
var errEmptyURL = errors.New("URL cannot be empty", "URL cannot be empty")

// schemeFromURL returns the scheme from a URL string
func SchemeFromURL(url string) (string, error) {
	if url == "" {
		return "", errEmptyURL
	}

	i := strings.Index(url, "://")

	// No : or : is the first character.
	if i < 1 {
		return "", errNoScheme
	}

	return url[0:i], nil
}

// Send notification to webhook
func SendStatusNotification(service service.ServiceChecker, statusCode constants.HealthCheckStatus, err errors.Error) {
	if webhook := service.GetWebHook(); webhook != nil {
		webhook(service.GetCheck(), statusCode, err)
		return
	}

	switch statusCode {
	case constants.Success:
		check := service.GetCheck()
		logrus.Infof("[%v] health check success\n", check.GetName())

	case constants.Failure:
		check := service.GetCheck()
		logrus.Errorf("[%v] service check failing due to : %v", check.GetName(), err.Reason())
	}
}
