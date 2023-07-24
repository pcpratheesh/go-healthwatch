package gohealthwatch

import (
	"context"
	"net/http"
	"time"

	"github.com/pcpratheesh/go-healthwatch/config"
	"github.com/pcpratheesh/go-healthwatch/constants"
	"github.com/pcpratheesh/go-healthwatch/service"
	"github.com/pcpratheesh/go-healthwatch/utils/errors"
	"github.com/pcpratheesh/go-healthwatch/utils/worker"
	"github.com/sirupsen/logrus"
)

type Options func(health *HealthCheck)

type HealthCheck struct {
	Checks                    []config.HealthCheckConfig
	StatusNotificationWebhook config.ServiceStatusNotificationHook
}

// NewChecker
func NewChecker(options ...Options) *HealthCheck {
	checker := &HealthCheck{}

	for _, opt := range options {
		opt(checker)
	}

	return checker
}

// AddIntegrations
func WithIntegrations(checks []config.HealthCheckConfig) Options {
	return func(health *HealthCheck) {
		health.Checks = checks
	}
}

// WithServiceFailureHandler
func WithServiceStatusWebHook(handler config.ServiceStatusNotificationHook) Options {
	return func(health *HealthCheck) {
		health.StatusNotificationWebhook = handler
	}
}

// Append new integration
func (health *HealthCheck) AddIntegration(integration config.HealthCheckConfig) *HealthCheck {
	health.Checks = append(health.Checks, integration)
	return health
}

// Custom checking
func (health *HealthCheck) AddCheck(name string, callback config.CustomHandler) {
	config.CustomHandlerMap[name] = callback
}

// Check the services status integrated
func (health *HealthCheck) Check() {
	workers := worker.NewTasks()

	for _, integration := range health.Checks {
		srv := service.InitService(integration, health.StatusNotificationWebhook)

		if err := srv.Validate(); err != nil {
			logrus.Errorf("[%v] service validation failed : %v", integration.GetName(), err)
		}

		workers = workers.Add(integration.GetName(), srv, integration.Interval)
	}

	workers.Start(context.Background())
}

func main() {
	checker := NewChecker(
		WithIntegrations([]config.HealthCheckConfig{
			{
				Name:       "profile-api",
				URL:        "https://profile-dev.my.mtn.com/api/v1/health",
				Type:       constants.External,
				StatusCode: http.StatusOK,
				Interval:   time.Second * 1,
			},
		}),
		// WithServiceStatusWebHook(func(check config.HealthCheckConfig, statusCode constants.HealthCheckStatus, err errors.Error) {
		// 	switch statusCode {
		// 	case constants.Success:
		// 		logrus.Infof("Custom Handler [%v] health check success\n", check.GetName())

		// 	case constants.Failure:
		// 		logrus.Errorf("Custom Handler  [%v] service check failing due to : %v", check.GetName(), err.Reason())
		// 	}
		// }),
	)

	checker.AddCheck("profile-api", func(check config.HealthCheckConfig) errors.Error {
		return errors.New("trigger-failure", "")
	})

	checker.Check()
}
