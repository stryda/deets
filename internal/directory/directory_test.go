package directory

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DirectoryTestSuite struct {
	suite.Suite
}

func (suite *DirectoryTestSuite) TestPrint() {
	base := Print("/this/is/a/file.test")
	assert.Equal(suite.T(), "file.test", base)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(DirectoryTestSuite))
}
