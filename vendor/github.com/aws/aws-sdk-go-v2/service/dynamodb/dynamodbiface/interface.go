// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package dynamodbiface provides an interface to enable mocking the Amazon DynamoDB service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package dynamodbiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// DynamoDBAPI provides an interface to enable mocking the
// dynamodb.DynamoDB service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon DynamoDB.
//    func myFunc(svc dynamodbiface.DynamoDBAPI) bool {
//        // Make svc.BatchGetItem request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := dynamodb.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockDynamoDBClient struct {
//        dynamodbiface.DynamoDBAPI
//    }
//    func (m *mockDynamoDBClient) BatchGetItem(input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockDynamoDBClient{}
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
type DynamoDBAPI interface {
	BatchGetItemRequest(*dynamodb.BatchGetItemInput) dynamodb.BatchGetItemRequest

	BatchWriteItemRequest(*dynamodb.BatchWriteItemInput) dynamodb.BatchWriteItemRequest

	CreateBackupRequest(*dynamodb.CreateBackupInput) dynamodb.CreateBackupRequest

	CreateGlobalTableRequest(*dynamodb.CreateGlobalTableInput) dynamodb.CreateGlobalTableRequest

	CreateTableRequest(*dynamodb.CreateTableInput) dynamodb.CreateTableRequest

	DeleteBackupRequest(*dynamodb.DeleteBackupInput) dynamodb.DeleteBackupRequest

	DeleteItemRequest(*dynamodb.DeleteItemInput) dynamodb.DeleteItemRequest

	DeleteTableRequest(*dynamodb.DeleteTableInput) dynamodb.DeleteTableRequest

	DescribeBackupRequest(*dynamodb.DescribeBackupInput) dynamodb.DescribeBackupRequest

	DescribeContinuousBackupsRequest(*dynamodb.DescribeContinuousBackupsInput) dynamodb.DescribeContinuousBackupsRequest

	DescribeEndpointsRequest(*dynamodb.DescribeEndpointsInput) dynamodb.DescribeEndpointsRequest

	DescribeGlobalTableRequest(*dynamodb.DescribeGlobalTableInput) dynamodb.DescribeGlobalTableRequest

	DescribeGlobalTableSettingsRequest(*dynamodb.DescribeGlobalTableSettingsInput) dynamodb.DescribeGlobalTableSettingsRequest

	DescribeLimitsRequest(*dynamodb.DescribeLimitsInput) dynamodb.DescribeLimitsRequest

	DescribeTableRequest(*dynamodb.DescribeTableInput) dynamodb.DescribeTableRequest

	DescribeTimeToLiveRequest(*dynamodb.DescribeTimeToLiveInput) dynamodb.DescribeTimeToLiveRequest

	GetItemRequest(*dynamodb.GetItemInput) dynamodb.GetItemRequest

	ListBackupsRequest(*dynamodb.ListBackupsInput) dynamodb.ListBackupsRequest

	ListGlobalTablesRequest(*dynamodb.ListGlobalTablesInput) dynamodb.ListGlobalTablesRequest

	ListTablesRequest(*dynamodb.ListTablesInput) dynamodb.ListTablesRequest

	ListTagsOfResourceRequest(*dynamodb.ListTagsOfResourceInput) dynamodb.ListTagsOfResourceRequest

	PutItemRequest(*dynamodb.PutItemInput) dynamodb.PutItemRequest

	QueryRequest(*dynamodb.QueryInput) dynamodb.QueryRequest

	RestoreTableFromBackupRequest(*dynamodb.RestoreTableFromBackupInput) dynamodb.RestoreTableFromBackupRequest

	RestoreTableToPointInTimeRequest(*dynamodb.RestoreTableToPointInTimeInput) dynamodb.RestoreTableToPointInTimeRequest

	ScanRequest(*dynamodb.ScanInput) dynamodb.ScanRequest

	TagResourceRequest(*dynamodb.TagResourceInput) dynamodb.TagResourceRequest

	TransactGetItemsRequest(*dynamodb.TransactGetItemsInput) dynamodb.TransactGetItemsRequest

	TransactWriteItemsRequest(*dynamodb.TransactWriteItemsInput) dynamodb.TransactWriteItemsRequest

	UntagResourceRequest(*dynamodb.UntagResourceInput) dynamodb.UntagResourceRequest

	UpdateContinuousBackupsRequest(*dynamodb.UpdateContinuousBackupsInput) dynamodb.UpdateContinuousBackupsRequest

	UpdateGlobalTableRequest(*dynamodb.UpdateGlobalTableInput) dynamodb.UpdateGlobalTableRequest

	UpdateGlobalTableSettingsRequest(*dynamodb.UpdateGlobalTableSettingsInput) dynamodb.UpdateGlobalTableSettingsRequest

	UpdateItemRequest(*dynamodb.UpdateItemInput) dynamodb.UpdateItemRequest

	UpdateTableRequest(*dynamodb.UpdateTableInput) dynamodb.UpdateTableRequest

	UpdateTimeToLiveRequest(*dynamodb.UpdateTimeToLiveInput) dynamodb.UpdateTimeToLiveRequest

	WaitUntilTableExists(*dynamodb.DescribeTableInput) error
	WaitUntilTableExistsWithContext(aws.Context, *dynamodb.DescribeTableInput, ...aws.WaiterOption) error

	WaitUntilTableNotExists(*dynamodb.DescribeTableInput) error
	WaitUntilTableNotExistsWithContext(aws.Context, *dynamodb.DescribeTableInput, ...aws.WaiterOption) error
}

var _ DynamoDBAPI = (*dynamodb.DynamoDB)(nil)
