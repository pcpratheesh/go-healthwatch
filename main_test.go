package gohealthwatch_test

import (
	"net/http"
	"testing"

	gohealthwatch "github.com/pcpratheesh/go-healthwatch"
	"github.com/pcpratheesh/go-healthwatch/constants"
	"github.com/pcpratheesh/go-healthwatch/errors"
	"github.com/pcpratheesh/go-healthwatch/models"
	"github.com/stretchr/testify/require"
)

func TestHealthWatch(t *testing.T) {
	t.Run("", func(t *testing.T) {

		checker := gohealthwatch.NewChecker(
			gohealthwatch.WithIntegrations([]models.HealthCheckConfig{
				{
					Name:       "public-entries",
					URL:        "https://api.publicapis.org/entries",
					Type:       constants.External,
					StatusCode: http.StatusOK,
					Interval:   -1,
				},
			}),
			gohealthwatch.WithServiceStatusWebHook(func(check models.HealthCheckConfig, statusCode constants.HealthCheckStatus, err errors.Error) {
				if check.GetName() == "public-entries" {
					require.Equal(t, constants.Success, statusCode)
				}
			}),
		)
		checker.Check()
	})
}

func TestHealthWatchWithCustomhandler(t *testing.T) {
	t.Run("", func(t *testing.T) {

		checker := gohealthwatch.NewChecker(
			gohealthwatch.WithIntegrations([]models.HealthCheckConfig{
				{
					Name:       "public-entries",
					URL:        "no-url",
					Type:       constants.External,
					StatusCode: http.StatusOK,
					Interval:   -1,
				},
			}),
			gohealthwatch.WithServiceStatusWebHook(func(check models.HealthCheckConfig, statusCode constants.HealthCheckStatus, err errors.Error) {
				if check.GetName() == "public-entries" {
					require.Equal(t, constants.Success, statusCode)
				}
			}),
		)

		// add custom handler for public-entries
		checker.AddCheck("public-entries", func(check models.HealthCheckConfig) errors.Error {
			return nil
		})
		checker.Check()
	})
}

func TestHealthWatchWithCustomhandlerFailure(t *testing.T) {
	t.Run("", func(t *testing.T) {

		checker := gohealthwatch.NewChecker(
			gohealthwatch.WithIntegrations([]models.HealthCheckConfig{
				{
					Name:       "public-entries",
					URL:        "no-url",
					Type:       constants.External,
					StatusCode: http.StatusOK,
					Interval:   -1,
				},
			}),
			gohealthwatch.WithServiceStatusWebHook(func(check models.HealthCheckConfig, statusCode constants.HealthCheckStatus, err errors.Error) {
				if check.GetName() == "public-entries" {
					require.Equal(t, constants.Failure, statusCode)
				}
			}),
		)

		// add custom handler for public-entries
		checker.AddCheck("public-entries", func(check models.HealthCheckConfig) errors.Error {
			return errors.New("trigger-failure", "")
		})

		checker.Check()
	})
}
