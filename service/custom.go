package service

import (
	"context"

	"github.com/pcpratheesh/go-healthwatch/errors"
	"github.com/pcpratheesh/go-healthwatch/models"
)

type customHandler struct {
	check   models.HealthCheckConfig
	handler models.CustomHandler
	webHook models.ServiceStatusNotificationHook
}

func NewCustomHandlerService(check models.HealthCheckConfig, handler models.CustomHandler, webHook models.ServiceStatusNotificationHook) *customHandler {
	return &customHandler{
		check:   check,
		handler: handler,
		webHook: webHook,
	}
}

// Check the integration
func (custom *customHandler) Check(ctx context.Context) errors.Error {
	return custom.handler(custom.check)
}

func (custom *customHandler) Validate() error {
	return nil
}
func (custom *customHandler) GetCheck() models.HealthCheckConfig {
	return custom.check
}
func (custom *customHandler) GetWebHook() models.ServiceStatusNotificationHook {
	return custom.webHook
}
