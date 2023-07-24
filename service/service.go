package service

import (
	"context"

	"github.com/pcpratheesh/go-healthwatch/config"
	"github.com/pcpratheesh/go-healthwatch/constants"
	"github.com/pcpratheesh/go-healthwatch/utils/errors"
)

type ServiceChecker interface {
	Validate() error
	Check(context.Context) errors.Error
	GetCheck() config.HealthCheckConfig
	GetWebHook() config.ServiceStatusNotificationHook
}

func InitService(check config.HealthCheckConfig, webhook config.ServiceStatusNotificationHook) ServiceChecker {
	// check where the custom handler is initialized
	if _, ok := config.CustomHandlerMap[check.Name]; ok {
		return NewCustomHandlerService(check, config.CustomHandlerMap[check.GetName()], webhook)
	}

	switch check.Type {
	case constants.External:
		return NewThirdPartyServiceCheck(check, webhook)
	default:
		return &UnImplementedService{
			Type: string(check.Type),
		}
	}

}
