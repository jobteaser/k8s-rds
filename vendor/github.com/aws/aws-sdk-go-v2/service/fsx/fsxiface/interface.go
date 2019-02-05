// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package fsxiface provides an interface to enable mocking the Amazon FSx service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package fsxiface

import (
	"github.com/aws/aws-sdk-go-v2/service/fsx"
)

// FSxAPI provides an interface to enable mocking the
// fsx.FSx service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon FSx.
//    func myFunc(svc fsxiface.FSxAPI) bool {
//        // Make svc.CreateBackup request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := fsx.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockFSxClient struct {
//        fsxiface.FSxAPI
//    }
//    func (m *mockFSxClient) CreateBackup(input *fsx.CreateBackupInput) (*fsx.CreateBackupOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockFSxClient{}
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
type FSxAPI interface {
	CreateBackupRequest(*fsx.CreateBackupInput) fsx.CreateBackupRequest

	CreateFileSystemRequest(*fsx.CreateFileSystemInput) fsx.CreateFileSystemRequest

	CreateFileSystemFromBackupRequest(*fsx.CreateFileSystemFromBackupInput) fsx.CreateFileSystemFromBackupRequest

	DeleteBackupRequest(*fsx.DeleteBackupInput) fsx.DeleteBackupRequest

	DeleteFileSystemRequest(*fsx.DeleteFileSystemInput) fsx.DeleteFileSystemRequest

	DescribeBackupsRequest(*fsx.DescribeBackupsInput) fsx.DescribeBackupsRequest

	DescribeFileSystemsRequest(*fsx.DescribeFileSystemsInput) fsx.DescribeFileSystemsRequest

	ListTagsForResourceRequest(*fsx.ListTagsForResourceInput) fsx.ListTagsForResourceRequest

	TagResourceRequest(*fsx.TagResourceInput) fsx.TagResourceRequest

	UntagResourceRequest(*fsx.UntagResourceInput) fsx.UntagResourceRequest

	UpdateFileSystemRequest(*fsx.UpdateFileSystemInput) fsx.UpdateFileSystemRequest
}

var _ FSxAPI = (*fsx.FSx)(nil)
