package echorelic

import (
	"github.com/labstack/echo"
	"github.com/newrelic/go-agent"
)

type EchoRelic interface {
	Init(appName string, licenseKey string) (newrelic.Application, error)
	EchoRelicMiddleware() echo.MiddlewareFunc
}

type Relic struct {
	app  newrelic.Application
	name string
}

func (r *Relic) Init(appName string, licenseKey string) (newrelic.Application, error) {
	config := newrelic.NewConfig(appName, licenseKey)
	app, err := newrelic.NewApplication(config)
	r.app = app
	r.name = appName
	return app, err
}

func (r *Relic) EchoRelicMiddleware() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			txn := r.app.StartTransaction(c.Path(), c.Response(), c.Request())
			defer txn.End()
			c.Set("RelicTransaction", txn)
			return h(c)
		}
	}
}
