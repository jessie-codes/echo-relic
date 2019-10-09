package echorelic

import (
	"github.com/labstack/echo"
	newrelic "github.com/newrelic/go-agent"
)

//EchoRelic stores the configured newrelic.Application
type EchoRelic struct {
	app newrelic.Application
}

//New creates an instance of type EchoRelic
func New(appName, licenseKey string) (*EchoRelic, error) {
	config := newrelic.NewConfig(appName, licenseKey)
	app, err := newrelic.NewApplication(config)
	if err != nil {
		return nil, err
	}
	return &EchoRelic{
		app: app,
	}, nil
}

//Transaction is an echo middleware function which creates a new transaction for each request
func (e *EchoRelic) Transaction(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Request().Method + " " + c.Path()
		txn := e.app.StartTransaction(name, c.Response().Writer, c.Request())
		txn.AddAttribute("RealIP", c.RealIP())
		txn.AddAttribute("IsTLS", c.IsTLS())
		txn.AddAttribute("IsWebSocket", c.IsWebSocket())
		txn.AddAttribute("Query", c.QueryString())
		defer txn.End()
		next(c)
		return nil
	}
}
