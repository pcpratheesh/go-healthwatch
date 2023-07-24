package service

import (
	"context"

	"github.com/pcpratheesh/go-healthwatch/constants"
	"github.com/pcpratheesh/go-healthwatch/errors"
	"github.com/pcpratheesh/go-healthwatch/models"
)

type ServiceChecker interface {
	Validate() error
	Check(context.Context) errors.Error
	GetCheck() models.HealthCheckConfig
	GetWebHook() models.ServiceStatusNotificationHook
}

func InitService(check models.HealthCheckConfig, webhook models.ServiceStatusNotificationHook) ServiceChecker {
	// check where the custom handler is initialized
	if _, ok := models.CustomHandlerMap[check.Name]; ok {
		return NewCustomHandlerService(check, models.CustomHandlerMap[check.GetName()], webhook)
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
