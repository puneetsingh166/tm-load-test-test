package types_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/puneetsingh166/tm-load-test/testutil/testdata"
	sdk "github.com/puneetsingh166/tm-load-test/types"
)

type testMsgSuite struct {
	suite.Suite
}

func TestMsgTestSuite(t *testing.T) {
	suite.Run(t, new(testMsgSuite))
}

func (s *testMsgSuite) TestMsg() {
	addr := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	accAddr := sdk.AccAddress(addr)

	msg := testdata.NewTestMsg(accAddr)
	s.Require().NotNil(msg)
	s.Require().True(accAddr.Equals(msg.GetSigners()[0]))
	s.Require().Equal("TestMsg", msg.Route())
	s.Require().Equal("Test message", msg.Type())
	s.Require().Nil(msg.ValidateBasic())
	s.Require().NotPanics(func() { msg.GetSignBytes() })
}

func (s *testMsgSuite) TestMsgTypeURL() {
	s.Require().Equal("/testdata.TestMsg", sdk.MsgTypeURL(new(testdata.TestMsg)))
}
