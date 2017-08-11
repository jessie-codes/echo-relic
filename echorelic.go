package echorelic

import (
	"github.com/labstack/echo"
	"github.com/newrelic/go-agent"
)

// EchoRelic contains the connection to New Relic
type EchoRelic struct {
	app  newrelic.Application
	name string
}

// Init is used to create a newrelic.Application
func (r *EchoRelic) Init(appName string, licenseKey string) (newrelic.Application, error) {
	config := newrelic.NewConfig(appName, licenseKey)
	app, err := newrelic.NewApplication(config)
	r.app = app
	r.name = appName
	return app, err
}

// EchoRelicMiddleware starts a newrelic.Transactions and adds it to the request's echo.Context
func (r *EchoRelic) EchoRelicMiddleware() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			txn := r.app.StartTransaction(c.Path(), c.Response(), c.Request())
			c.Set("RelicTransaction", txn)
			return h(c)
		}
	}
}
