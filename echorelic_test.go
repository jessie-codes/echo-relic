package echorelic

import (
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	newrelic "github.com/newrelic/go-agent"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	echoRelic *EchoRelic
	e         *echo.Echo
}

func (suite *TestSuite) SetupTest() {
	echoRelic, err := New("test", "1234567890123456789012345678901234567890")
	if err != nil {
		suite.Fail("Failed to create new EchoRelic")
	}
	suite.echoRelic = echoRelic
	suite.e = echo.New()
}

func (suite *TestSuite) TestUseMiddleware() {
	var t newrelic.Transaction
	suite.e.Use(suite.echoRelic.Transaction)
	req := httptest.NewRequest(echo.GET, "/", nil)
	res := httptest.NewRecorder()
	c := suite.e.NewContext(req, res)
	txn := c.Get("newRelicTransaction")
	suite.IsType(t, txn)
}

func TestMethodSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
