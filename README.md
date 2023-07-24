# go-healthwatch
Go-HealthWatch is a lightweight Golang package that provides health check functionality for applications and services. It allows you to monitor the health of various components, such as external APIs, databases, caches, and more. With GoHealthWatch, you can easily set up health checks and receive notifications when any issues arise.

## Installation
To install GoHealthWatch, use the following command:
```
go get github.com/pcpratheesh/go-healthwatch
```

## Usage
### Creating a Health Check

To create a health check, you can use the `NewChecker` function and provide the necessary options. You can specify the health checks to be performed and configure a service status notification webhook.

```go
healthchecker := gohealthwatch.NewChecker(
    gohealthwatch.WithIntegrations([]gohealthwatch.HealthCheckConfig{
        {
            Name:       "public-entries",
            URL:        "https://api.publicapis.org/entries",
            Type:       constants.External,
            StatusCode: http.StatusOK,
            Interval:   time.Second * 1,
        },
    }),
)
```

### HealthCheckConfig

This is a configuration struct used to hold the settings for a health check. It contains the following fields: 

Field | Description | required
--- | --- | --- 
Name | The unique name of the health check | True
URL | The URL to check. This can be anything like an external API URL, database URL, cache URL, etc | True
Type | The type of the health check, which can be one of the constants defined in the constants.Kind enum | True
Interval | The interval duration for the health check in seconds | False
HTTPHeader | An array of HTTP headers to include in the health check request | False
StatusCode | The expected HTTP status code for the health check response | True

### Run your own Checks for individual checklist

You can also override the checks to the health check configuration using the `AddCheck` method. This allows you to override checks based on your specific requirements.

The first argument in the checker.AddCheck() function should be the name of the health check. The second argument is a callback function that takes a `HealthCheckConfig` object as input and returns an `errors.Error` object.

```go
checker.AddCheck("public-entries", func(check gohealthwatch.HealthCheckConfig) errors.Error {
    return errors.New("trigger-failure", "")
})
```


### Performing Health Checks

To perform the health checks, simply call the `Check` method on the `HealthCheck` instance.
```go
checker.Check()
```


### Service Status Notification WebHook

The `WithServiceStatusWebHook` function is used to configure a service status notification webhook for the `HealthCheck` struct. It takes a callback function as an argument, which is invoked when a health check status changes.

```go
gohealthwatch.WithServiceStatusWebHook(func(check gohealthwatch.HealthCheckConfig, statusCode constants.HealthCheckStatus, err errors.Error) {
    switch statusCode {
    case constants.Success:
        logrus.Infof("Custom Handler [%v] health check success\n", check.GetName())
    case constants.Failure:
        logrus.Errorf("Custom Handler  [%v] service check failing due to : %v", check.GetName(), err.Reason())
    }
})
```

This allows you to customize the behavior when a health check status changes. You can perform actions such as logging, sending notifications, or triggering other processes based on the status and error information.


## Example
See more samples at [here](/example/)

## Contributing
Contributions to GoHealthWatch are welcome! If you find any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request on the GitHub repository.

## License
GoHealthWatch is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.



