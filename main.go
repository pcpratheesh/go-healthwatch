package gohealthwatch

import (
	"context"
	"net/http"
	"time"

	"github.com/pcpratheesh/go-healthwatch/constants"
	"github.com/pcpratheesh/go-healthwatch/errors"
	"github.com/pcpratheesh/go-healthwatch/models"
	"github.com/pcpratheesh/go-healthwatch/service"
	"github.com/pcpratheesh/go-healthwatch/worker"
	"github.com/sirupsen/logrus"
)

type Options func(health *HealthCheck)

type HealthCheck struct {
	Checks                    []models.HealthCheckConfig
	StatusNotificationWebhook models.ServiceStatusNotificationHook
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
func WithIntegrations(checks []models.HealthCheckConfig) Options {
	return func(health *HealthCheck) {
		health.Checks = checks
	}
}

// WithServiceFailureHandler
func WithServiceStatusWebHook(handler models.ServiceStatusNotificationHook) Options {
	return func(health *HealthCheck) {
		health.StatusNotificationWebhook = handler
	}
}

// Append new integration
func (health *HealthCheck) AddIntegration(integration models.HealthCheckConfig) *HealthCheck {
	health.Checks = append(health.Checks, integration)
	return health
}

// Custom checking
func (health *HealthCheck) AddCheck(name string, callback models.CustomHandler) {
	models.CustomHandlerMap[name] = callback
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
		WithIntegrations([]models.HealthCheckConfig{
			{
				Name:       "profile-api",
				URL:        "https://profile-dev.my.mtn.com/api/v1/health",
				Type:       constants.External,
				StatusCode: http.StatusOK,
				Interval:   time.Second * 1,
			},
		}),
		// WithServiceStatusWebHook(func(check models.HealthCheckConfig, statusCode constants.HealthCheckStatus, err errors.Error) {
		// 	switch statusCode {
		// 	case constants.Success:
		// 		logrus.Infof("Custom Handler [%v] health check success\n", check.GetName())

		// 	case constants.Failure:
		// 		logrus.Errorf("Custom Handler  [%v] service check failing due to : %v", check.GetName(), err.Reason())
		// 	}
		// }),
	)

	checker.AddCheck("profile-api", func(check models.HealthCheckConfig) errors.Error {
		return errors.New("trigger-failure", "")
	})

	checker.Check()
}
