package keeper_test

import (
	"testing"

	apptesing "github.com/nghuyenthevinh2000/nebula/app/apptesting"
	"github.com/nghuyenthevinh2000/nebula/x/launchpad/types"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	apptesing.KeeperTestHelper
}

func (suite *KeeperTestSuite) SetupTest() {
	// setup KeeperTestSuite
	suite.SetupKeeperTestHelper()

	// set params of global_project_id
	suite.App.LaunchpadKeeper.SetParams(suite.Ctx, types.Params{})
	suite.App.LaunchpadKeeper.SetNextProjectID(suite.Ctx, 0)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
