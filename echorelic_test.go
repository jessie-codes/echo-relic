package echorelic

import (
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/newrelic/go-agent"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	app newrelic.Application
	e   *echo.Echo
}

func (suite *TestSuite) SetupTest() {
	config := newrelic.NewConfig("test", "1234567890123456789012345678901234567890")
	app, _ := newrelic.NewApplication(config)
	suite.app = app
	suite.e = echo.New()
}

func (suite *TestSuite) TestMiddleware() {
	m := Middleware(suite.app)
	suite.NotEmpty(m)
}

func (suite *TestSuite) TestUseMiddleware() {
	var t newrelic.Transaction
	suite.e.Use(Middleware(suite.app))
	req := httptest.NewRequest(echo.GET, "/", nil)
	res := httptest.NewRecorder()
	c := suite.e.NewContext(req, res)
	txn := c.Get("newRelicTransaction")
	suite.IsType(t, txn)
}

func TestMethodSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
