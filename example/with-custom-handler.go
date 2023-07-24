package main

import (
	"net/http"

	gohealthwatch "github.com/pcpratheesh/go-healthwatch"
	"github.com/pcpratheesh/go-healthwatch/config"
	"github.com/pcpratheesh/go-healthwatch/constants"
	"github.com/pcpratheesh/go-healthwatch/utils/errors"
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

	// register the handler
	checker.AddCheck("public-entries", func(check config.HealthCheckConfig) errors.Error {
		// check some other api status before
		res, err := http.Get("https://catfact.ninja/fact")
		if err != nil {
			return errors.New("first api checked failed", "API FAILED")
		}

		if res.StatusCode != http.StatusOK {
			return errors.New("Status code check failure", "API FAILED")
		}

		// uncomment to make it fail
		// return errors.New("making this fail", "CUSTOM FAIL")

		// if not then check the actual endpoint
		res, err = http.Get(check.URL)
		if err != nil {
			return errors.New("first api checked failed", "API FAILED")
		}

		return nil
	})

	// Check
	checker.Check()
}
