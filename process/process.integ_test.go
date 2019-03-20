package process

import (
	"fmt"
	"os"
	"testing"

	"github.com/pulpfree/univsales-pdf-url/config"
	"github.com/pulpfree/univsales-pdf-url/model"
	"github.com/stretchr/testify/suite"
)

const (
	defaultsFilePath = "../config/defaults.yml"
	quoteNumber      = 1083
	// quoteNumber = 1064
	quoteVersion = 1
	// quoteVersion = 2
	requestType = "quote"
)

// IntegSuite struct
type IntegSuite struct {
	suite.Suite
	cfg     *config.Config
	request *model.Request
	process *Process
}

// SetupTest method
func (suite *IntegSuite) SetupTest() {
	os.Setenv("Stage", "test")
	suite.request = &model.Request{
		Number:  quoteNumber,
		Type:    requestType,
		Version: quoteVersion,
	}
	suite.cfg = &config.Config{DefaultsFilePath: defaultsFilePath}
	err := suite.cfg.Load()
	suite.NoError(err)
	suite.IsType(new(config.Config), suite.cfg)

	suite.process, err = New(suite.request, suite.cfg)
	suite.NoError(err)
	suite.IsType(new(Process), suite.process)
}

// TestCreateURL method
func (suite *IntegSuite) TestCreateSignedURL() {
	url, err := suite.process.CreateURL()
	fmt.Printf("url %s\n", url)
	suite.NoError(err)
	suite.NotEqual("", url, "Expected url to be populated")
}

// TestAssemblePrefix method
func (suite *IntegSuite) TestAssemblePrefix() {
	expectPrefix := "/quote/qte-1083-r1.pdf"
	prefix, _ := suite.process.assemblePrefix()
	suite.Equal(expectPrefix, prefix)
}

// TestIntegSuite function
func TestIntegSuite(t *testing.T) {
	suite.Run(t, new(IntegSuite))
}
