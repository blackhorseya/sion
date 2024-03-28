package biz

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type suiteTester struct {
	suite.Suite
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}
