package validate

import (
	"os"
	"testing"

	"github.com/pulpfree/univsales-pdf-url/model"
	"github.com/stretchr/testify/suite"
)

const (
	quoteNumber  = 1083
	quoteVersion = 1
	requestType  = "quote"
)

// IntegSuite struct
type IntegSuite struct {
	suite.Suite
	request *model.Request
}

// SetupTest method
func (suite *IntegSuite) SetupTest() {
	os.Setenv("Stage", "test")
	suite.request = &model.Request{
		Number:  quoteNumber,
		Type:    requestType,
		Version: quoteVersion,
	}
	suite.IsType(new(model.Request), suite.request)
}

// TestRequestInputFalse method
func (suite *IntegSuite) TestRequestInputFalse() {
	req := &model.Request{
		Number: quoteNumber,
		Type:   requestType,
	}
	err := RequestInput(req)
	suite.Error(err)
}

// TestRequestInputFalse method
func (suite *IntegSuite) TestRequestInput() {
	req := &model.Request{
		Number: quoteNumber,
		Type:   "invoice",
	}
	err := RequestInput(req)
	suite.NoError(err)
}

// TestUnitSuite function
func TestUnitSuite(t *testing.T) {
	suite.Run(t, new(IntegSuite))
}
