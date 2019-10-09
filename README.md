# echo-relic
Echo middleware for [New Relic](https://newrelic.com/)

[![GoDoc](https://godoc.org/github.com/jessie-codes/echo-relic?status.svg)](https://godoc.org/github.com/jessie-codes/echo-relic)
[![Build Status](https://travis-ci.org/jessie-codes/echo-relic.svg?branch=master)](https://travis-ci.org/jessie-codes/echo-relic)
[![Coverage Status](https://coveralls.io/repos/github/jessie-codes/echo-relic/badge.svg?branch=master)](https://coveralls.io/github/jessie-codes/echo-relic?branch=master)

## install

`go get github.com/jessie-codes/echo-relic`

## Usage

Echo Relic starts a new transaction for each request, binds the transaction to the request context, and end the transaction after handling has been completed. It uses the following convention for naming transaction: `<Method> <Path>`. It automatically adds attributes for `RealIP`, `IsTLS`, `IsWebSocket`, and `Query`. See [go-agent](https://github.com/newrelic/go-agent)'s documentation for how to use the transaction interface.

```golang

package main

import (
	"github.com/labstack/echo"
	"github.com/newrelic/go-agent"
	"github.com/jessie-codes/echo-relic/v3"
)

func main() {
	e := echo.New()
	relic = echorelic.New("__APP_NAME__", "__NEW_RELIC_LICENSE_KEY__")
	e.Use(relic.Transaction)

	e.GET("/", func(c echo.Context) error {
		txn := c.Get("newRelicTransaction")
		//route handle code
		return c.JSON(http.StatusOK, result)
	})
	e.Start(":8080")
}

```
