package main

import (
	"net/http"

	gohealthwatch "github.com/pcpratheesh/go-healthwatch"
	"github.com/pcpratheesh/go-healthwatch/constants"
	"github.com/pcpratheesh/go-healthwatch/models"
)

func main() {
	checker := gohealthwatch.NewChecker(
		gohealthwatch.WithIntegrations([]models.HealthCheckConfig{
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
