// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package securityhubiface provides an interface to enable mocking the AWS SecurityHub service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package securityhubiface

import (
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
)

// SecurityHubAPI provides an interface to enable mocking the
// securityhub.SecurityHub service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS SecurityHub.
//    func myFunc(svc securityhubiface.SecurityHubAPI) bool {
//        // Make svc.AcceptInvitation request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := securityhub.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockSecurityHubClient struct {
//        securityhubiface.SecurityHubAPI
//    }
//    func (m *mockSecurityHubClient) AcceptInvitation(input *securityhub.AcceptInvitationInput) (*securityhub.AcceptInvitationOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockSecurityHubClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type SecurityHubAPI interface {
	AcceptInvitationRequest(*securityhub.AcceptInvitationInput) securityhub.AcceptInvitationRequest

	BatchDisableStandardsRequest(*securityhub.BatchDisableStandardsInput) securityhub.BatchDisableStandardsRequest

	BatchEnableStandardsRequest(*securityhub.BatchEnableStandardsInput) securityhub.BatchEnableStandardsRequest

	BatchImportFindingsRequest(*securityhub.BatchImportFindingsInput) securityhub.BatchImportFindingsRequest

	CreateInsightRequest(*securityhub.CreateInsightInput) securityhub.CreateInsightRequest

	CreateMembersRequest(*securityhub.CreateMembersInput) securityhub.CreateMembersRequest

	DeclineInvitationsRequest(*securityhub.DeclineInvitationsInput) securityhub.DeclineInvitationsRequest

	DeleteInsightRequest(*securityhub.DeleteInsightInput) securityhub.DeleteInsightRequest

	DeleteInvitationsRequest(*securityhub.DeleteInvitationsInput) securityhub.DeleteInvitationsRequest

	DeleteMembersRequest(*securityhub.DeleteMembersInput) securityhub.DeleteMembersRequest

	DisableImportFindingsForProductRequest(*securityhub.DisableImportFindingsForProductInput) securityhub.DisableImportFindingsForProductRequest

	DisableSecurityHubRequest(*securityhub.DisableSecurityHubInput) securityhub.DisableSecurityHubRequest

	DisassociateFromMasterAccountRequest(*securityhub.DisassociateFromMasterAccountInput) securityhub.DisassociateFromMasterAccountRequest

	DisassociateMembersRequest(*securityhub.DisassociateMembersInput) securityhub.DisassociateMembersRequest

	EnableImportFindingsForProductRequest(*securityhub.EnableImportFindingsForProductInput) securityhub.EnableImportFindingsForProductRequest

	EnableSecurityHubRequest(*securityhub.EnableSecurityHubInput) securityhub.EnableSecurityHubRequest

	GetEnabledStandardsRequest(*securityhub.GetEnabledStandardsInput) securityhub.GetEnabledStandardsRequest

	GetFindingsRequest(*securityhub.GetFindingsInput) securityhub.GetFindingsRequest

	GetInsightResultsRequest(*securityhub.GetInsightResultsInput) securityhub.GetInsightResultsRequest

	GetInsightsRequest(*securityhub.GetInsightsInput) securityhub.GetInsightsRequest

	GetInvitationsCountRequest(*securityhub.GetInvitationsCountInput) securityhub.GetInvitationsCountRequest

	GetMasterAccountRequest(*securityhub.GetMasterAccountInput) securityhub.GetMasterAccountRequest

	GetMembersRequest(*securityhub.GetMembersInput) securityhub.GetMembersRequest

	InviteMembersRequest(*securityhub.InviteMembersInput) securityhub.InviteMembersRequest

	ListEnabledProductsForImportRequest(*securityhub.ListEnabledProductsForImportInput) securityhub.ListEnabledProductsForImportRequest

	ListInvitationsRequest(*securityhub.ListInvitationsInput) securityhub.ListInvitationsRequest

	ListMembersRequest(*securityhub.ListMembersInput) securityhub.ListMembersRequest

	UpdateFindingsRequest(*securityhub.UpdateFindingsInput) securityhub.UpdateFindingsRequest

	UpdateInsightRequest(*securityhub.UpdateInsightInput) securityhub.UpdateInsightRequest
}

var _ SecurityHubAPI = (*securityhub.SecurityHub)(nil)
