package bot

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/testsuite"
)

type WorkflowTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite

	env *testsuite.TestWorkflowEnvironment
}

func TestWorkflowTestSuit(t *testing.T) {
	suite.Run(t, new(WorkflowTestSuite))
}

func (s *WorkflowTestSuite) SetupTest() {
	s.env = s.NewTestWorkflowEnvironment()
}

func (s *WorkflowTestSuite) AfterTest(suiteName, testName string) {
	s.env.AssertExpectations(s.T())
}

func (s *WorkflowTestSuite) Test_Success_BotSetTradeWorkflow() {
	s.env.RegisterWorkflow(BotSetTradeWorkflow)
	s.env.RegisterActivity(BotSetTradeActivity)

	input := "https://www.set.or.th/th/home"
	s.env.ExecuteWorkflow(BotSetTradeWorkflow, input)

	s.Require().True(s.env.IsWorkflowCompleted())
	var output WorkflowResult
	s.Require().NoError(s.env.GetWorkflowResult(&output))

}

func (s *WorkflowTestSuite) Test_Success_BotGoldWorkflow() {
	s.env.RegisterWorkflow(BotGoldWorkflow)
	s.env.RegisterActivity(BotGoldActivity)

	input := "https://xn--42cah7d0cxcvbbb9x.com"
	s.env.ExecuteWorkflow(BotGoldWorkflow, input)

	s.Require().True(s.env.IsWorkflowCompleted())
	var output WorkflowResult
	s.Require().NoError(s.env.GetWorkflowResult(&output))

}
