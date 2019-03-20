package config

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

// to test config package use: go test -v ./config

type IntegSuite struct {
	suite.Suite
	cfg *Config
}

// SetupTest method
func (suite *IntegSuite) SetupTest() {

	suite.cfg = &Config{}

	os.Setenv("Stage", "test")
	suite.cfg.setDefaults()
	suite.cfg.setEnvVars()
}

// TestSetDefaults method
func (suite *IntegSuite) TestSetDefaults() {
	err := suite.cfg.setDefaults()
	suite.NoError(err)
}

// TestSetEnvVars method
func (suite *IntegSuite) TestSetEnvVars() {

	var err error
	err = suite.cfg.setEnvVars()
	suite.NoError(err)

	// Change a var
	os.Setenv("Stage", "no exist")
	err = suite.cfg.setEnvVars()
	suite.Error(err)

	// Reset to valid stage
	os.Setenv("Stage", "test")
	suite.cfg.setEnvVars()
}

// TestValidateStage method
func (suite *IntegSuite) TestValidateStage() {
	err := suite.cfg.validateStage()
	suite.NoError(err)
}

// TestSetSSMParams function
// this test assumes that the CognitoClientID in defaults is empty
func (suite *IntegSuite) TestSetSSMParams() {
	CognitoClientIDBefore := defs.CognitoClientID
	err := suite.cfg.setSSMParams()
	suite.NoError(err)

	CognitoClientIDAfter := defs.CognitoClientID
	suite.True(strings.Compare(CognitoClientIDBefore, CognitoClientIDAfter) != 0)
}

// TestSetFinal function
func (suite *IntegSuite) TestSetFinal() {

	var se StageEnvironment
	err := suite.cfg.setFinal()

	suite.NoError(err)
	suite.Equal(suite.cfg.AWSRegion, defs.AWSRegion, "Expected Config.AWSRegion (%s) to equal defs.AWSRegion (%s)", suite.cfg.AWSRegion, defs.AWSRegion)
	suite.IsType(se, suite.cfg.Stage)
}

// TestIntegrationSuite function
func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, new(IntegSuite))
}
