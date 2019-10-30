package shell

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ShellTestSuite struct {
	suite.Suite
}

func (suite *ShellTestSuite) TestGetShell() {
	s := Print()
	assert.Equal(suite.T(), "", s)
}

func (suite *ShellTestSuite) TestEnvironment() {
	printEnv()
	assert.True(suite.T(), len(getEnv()) > 0)
}

func (suite *ShellTestSuite) TestGetCommand() {
	c := getCommand("whoami", []string{}, true)
	assert.NotEmpty(suite.T(), c)
}

func (suite *ShellTestSuite) TestEnvironmentVariable() {
	v := getEnvVariable("SHELL")
	assert.Equal(suite.T(), "/bin/bash", v)
}

func TestGitTestSuite(t *testing.T) {
	suite.Run(t, new(ShellTestSuite))
}
