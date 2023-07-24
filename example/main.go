package main

import (
	"net/http"

	gohealthwatch "github.com/pcpratheesh/go-healthwatch"
	"github.com/pcpratheesh/go-healthwatch/config"
	"github.com/pcpratheesh/go-healthwatch/constants"
)

func main() {
	checker := gohealthwatch.NewChecker(
		gohealthwatch.WithIntegrations([]config.HealthCheckConfig{
			{
				Name:       "public-entries",
				URL:        "https://api.publicapis.org/entries",
				Type:       constants.External,
				StatusCode: http.StatusOK,
				Interval:   1,
			},
		}),
	)
	checker.Check()
}
