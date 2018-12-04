# echo-relic
Echo middleware for [New Relic](https://newrelic.com/)

[![GoDoc](https://godoc.org/github.com/jessie-codes/echo-relic?status.svg)](https://godoc.org/github.com/jessie-codes/echo-relic)
[![Build Status](https://travis-ci.org/jessie-codes/echo-relic.svg?branch=master)](https://travis-ci.org/jessie-codes/echo-relic)

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
	"github.com/newrelic/go-agent"
	"github.com/jessie-codes/echo-relic"
)

func main() {
	e := echo.New()
	config := newrelic.NewConfig("__APP_NAME__", "__NEW_RELIC_LICENSE_KEY__")
	app, _ := newrelic.NewApplication(config)
	e.Use(echorelic.Middleware(app))

	e.GET("/", func(c echo.Context) error {
		txn := c.Get("newRelicTransaction")
		//route handle code
		return c.JSON(http.StatusOK, result)
	})
	e.Start(":8080")
}

```
