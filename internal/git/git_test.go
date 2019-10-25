package git

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type GitTestSuite struct {
	suite.Suite
}

func (suite *GitTestSuite) TestPrint() {
	path := Print("./")
	assert.Equal(suite.T(), "", path)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(GitTestSuite))
}