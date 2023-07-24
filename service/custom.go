package service

import (
	"context"

	"github.com/pcpratheesh/go-healthwatch/config"
	"github.com/pcpratheesh/go-healthwatch/utils/errors"
)

type customHandler struct {
	check   config.HealthCheckConfig
	handler config.CustomHandler
	webHook config.ServiceStatusNotificationHook
}

func NewCustomHandlerService(check config.HealthCheckConfig, handler config.CustomHandler, webHook config.ServiceStatusNotificationHook) *customHandler {
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
func (custom *customHandler) GetCheck() config.HealthCheckConfig {
	return custom.check
}
func (custom *customHandler) GetWebHook() config.ServiceStatusNotificationHook {
	return custom.webHook
}
