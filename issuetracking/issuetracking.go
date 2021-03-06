package issuetracking

import (
	"strings"

	"github.com/cyberark/dev-flow/common"
)

type IssueTrackingClient interface {
	GetCurrentUser() string
	GetUserRealName(string) string
	Issues() []common.Issue
	Issue(string) common.Issue
	AssignIssue(common.Issue, string)
	AddIssueLabel(common.Issue, string) error
	RemoveIssueLabel(common.Issue, string)
}

func GetClient() IssueTrackingClient {
	return GitHub{}
}

func GetIssueKeyFromBranchName(branchName string) string {
	return strings.Split(branchName, "--")[0]
}
