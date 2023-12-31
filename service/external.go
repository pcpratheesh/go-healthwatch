package service

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/pcpratheesh/go-healthwatch/config"
	"github.com/pcpratheesh/go-healthwatch/utils/errors"
)

type ExternalAPI struct {
	check   config.HealthCheckConfig
	webHook config.ServiceStatusNotificationHook
}

// NewThirdPartyServiceCheck
func NewThirdPartyServiceCheck(check config.HealthCheckConfig, webHook config.ServiceStatusNotificationHook) *ExternalAPI {
	return &ExternalAPI{
		check:   check,
		webHook: webHook,
	}
}

// Check the external api
func (external *ExternalAPI) Check(ctx context.Context) errors.Error {
	req, err := http.NewRequest(http.MethodGet, external.check.GetUrl(), nil)
	if err != nil {
		return errors.New(err.Error(), "request-failure")
	}

	// add headers
	if len(external.check.HTTPHeader) > 0 {
		for _, header := range external.check.HTTPHeader {
			req.Header.Add(header.Key, header.Value)
		}
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.New(err.Error(), "request-failure")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.New(err.Error(), "")
	}

	if res.StatusCode != external.check.StatusCode {
		return errors.New("status code checking failed", string(body))
	}

	return nil
}

// Validate the data
func (external *ExternalAPI) Validate() error {
	return nil
}

func (external *ExternalAPI) GetCheck() config.HealthCheckConfig {
	return external.check
}

func (external *ExternalAPI) GetWebHook() config.ServiceStatusNotificationHook {
	return external.webHook
}
