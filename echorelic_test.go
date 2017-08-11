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
	r *Relic
	e *echo.Echo
}

func (suite *TestSuite) SetupTest() {
	suite.r = new(Relic)
	suite.e = echo.New()
}

func (suite *TestSuite) TestInitNoLicense() {
	_, err := suite.r.Init("test", "")
	suite.Error(err)
}

func (suite *TestSuite) TestInitWithLicense() {
	app, err := suite.r.Init("test", "1234567890123456789012345678901234567890")
	suite.NoError(err)
	suite.NotEmpty(app)
}

func (suite *TestSuite) TestEchoRelicMiddleware() {
	m := suite.r.EchoRelicMiddleware()
	suite.NotEmpty(m)
}

func (suite *TestSuite) TestUseMiddleware() {
	var t newrelic.Transaction
	suite.e.Use(suite.r.EchoRelicMiddleware())
	req := httptest.NewRequest(echo.GET, "/", nil)
	res := httptest.NewRecorder()
	c := suite.e.NewContext(req, res)
	txn := c.Get("transaction")
	suite.IsType(t, txn)
}

func TestMethodSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
