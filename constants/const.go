package constants

type Kind string

var (
	API      Kind = "API"
	External Kind = "External"
)

type HealthCheckStatus string

var (
	Success HealthCheckStatus = "Success"
	Failure HealthCheckStatus = "Failure"
)
