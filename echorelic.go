package echorelic

import (
	"github.com/labstack/echo"
	"github.com/newrelic/go-agent"
)

var ctxKey = "newRelicTransaction"

// Middleware starts a newrelic.Transactions and adds it to the request's echo.Context
func Middleware(app newrelic.Application) echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			name := c.Request().Method + " " + c.Path()
			txn := app.StartTransaction(name, c.Response().Writer, c.Request())
			txn.AddAttribute("RealIP", c.RealIP())
			txn.AddAttribute("IsTLS", c.IsTLS())
			txn.AddAttribute("IsWebSocket", c.IsWebSocket())
			txn.AddAttribute("Query", c.QueryString())
			defer txn.End()

			c.Set(ctxKey, txn)
			return h(c)
		}
	}
}
