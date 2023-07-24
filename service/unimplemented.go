package service

import (
	"context"
	"fmt"

	"github.com/pcpratheesh/go-healthwatch/config"
	"github.com/pcpratheesh/go-healthwatch/utils/errors"
)

type UnImplementedService struct {
	Type string
}

func (impl *UnImplementedService) Check(ctx context.Context) errors.Error {
	return errors.New(fmt.Errorf("service Check for type [%s] is not on borded yet", impl.Type), "")
}
func (impl *UnImplementedService) Validate() error {
	return fmt.Errorf("service Validate for type [%s] is not on borded yet", impl.Type)
}
func (impl *UnImplementedService) GetCheck() config.HealthCheckConfig {
	return config.HealthCheckConfig{}
}
func (impl *UnImplementedService) GetWebHook() config.ServiceStatusNotificationHook {
	return nil
}
