package main

import (
	"fmt"
	"net/http"
	"time"

	gohealthwatch "github.com/pcpratheesh/go-healthwatch"
	"github.com/pcpratheesh/go-healthwatch/config"
	"github.com/pcpratheesh/go-healthwatch/constants"
	"github.com/pcpratheesh/go-healthwatch/utils/errors"
)

var webhookHandler = func(check config.HealthCheckConfig, statusCode constants.HealthCheckStatus, err errors.Error) {
	if statusCode == constants.Success {
		fmt.Printf("[%v] health check complete with %v status code \n", check.GetName(), statusCode)
	}

	if statusCode == constants.Failure {
		fmt.Printf("[%v] health check failed due to %v \n", check.GetName(), err.Reason())
	}
}

func main() {
	checker := gohealthwatch.NewChecker(
		gohealthwatch.WithIntegrations([]config.HealthCheckConfig{
			{
				Name:       "public-entries",
				URL:        "https://api.publicapis.org/entries123",
				Type:       constants.External,
				StatusCode: http.StatusOK,
				Interval:   time.Second * 4,
			},
		}),
		gohealthwatch.WithServiceStatusWebHook(webhookHandler),
	)

	// Check it
	checker.Check()
}
