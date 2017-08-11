# echo-relic
Echo middleware for [New Relic](https://newrelic.com/)

## install

`go get github.com/jessie-codes/echo-relic`

## Usage

Echo Relic starts a new transaction for each request and binds the transaction
to the request context. See [go-agent](https://github.com/newrelic/go-agent)'s
documentation for how to use the transaction interface.

```golang

package main

import (
	"github.com/labstack/echo"
	"github.com/jessie-codes/echo-relic"
)

func main() {
	e := echo.New()
	er := new(echorelic.EchoRelic)
	er.Init("__APP_NAME", "__NEW_RELIC_LICENSE_KEY")
	e.Use(er.EchoRelicMiddleware())

	e.GET("/", func(c echo.Context) error {
		txn := c.Get("transaction")
		//route handle code
		return c.JSON(http.StatusOK, result)
	})
	e.Start(":8080")
}

```
