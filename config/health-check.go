package config

import (
	"time"

	"github.com/pcpratheesh/go-healthwatch/constants"
	"github.com/pcpratheesh/go-healthwatch/utils/errors"
)

type CustomHandler func(check HealthCheckConfig) errors.Error
type ServiceStatusNotificationHook func(check HealthCheckConfig, statusCode constants.HealthCheckStatus, err errors.Error)

var CustomHandlerMap = make(map[string]CustomHandler, 0)

type HealthCheckConfig struct {
	Name       string
	URL        string // Url to HealthCheckConfig : This would be anything like external api url , database url, cache url...etc
	Type       constants.Kind
	TimeOut    int
	Interval   time.Duration
	HTTPHeader []HTTPHeader
	StatusCode int
}

type HTTPHeader struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"Value,omitempty"`
}

// Get the Name
func (check *HealthCheckConfig) GetName() string {
	return check.Name
}

// Get the Name
func (check *HealthCheckConfig) GetUrl() string {
	return check.URL
}

// Get the Type
func (check *HealthCheckConfig) GetType() string {
	return string(check.Type)
}
