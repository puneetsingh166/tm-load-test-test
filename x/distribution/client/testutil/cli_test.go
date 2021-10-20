// +build norace

package testutil

import (
	"testing"

	"github.com/puneetsingh166/tm-load-test/testutil/network"

	"github.com/stretchr/testify/suite"
)

func TestIntegrationTestSuite(t *testing.T) {
	cfg := network.DefaultConfig()
	cfg.NumValidators = 1
	suite.Run(t, NewIntegrationTestSuite(cfg))
}

func TestGRPCQueryTestSuite(t *testing.T) {
	suite.Run(t, new(GRPCQueryTestSuite))
}
