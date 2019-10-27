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
	shell := getShell()
	assert.Equal(suite.T(), "aaaa", shell)
}

func TestGitTestSuite(t *testing.T) {
	suite.Run(t, new(ShellTestSuite))
}