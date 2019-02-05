// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotjobsdataplane

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

const opDescribeJobExecution = "DescribeJobExecution"

// DescribeJobExecutionRequest is a API request type for the DescribeJobExecution API operation.
type DescribeJobExecutionRequest struct {
	*aws.Request
	Input *DescribeJobExecutionInput
	Copy  func(*DescribeJobExecutionInput) DescribeJobExecutionRequest
}

// Send marshals and sends the DescribeJobExecution API request.
func (r DescribeJobExecutionRequest) Send() (*DescribeJobExecutionOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*DescribeJobExecutionOutput), nil
}

// DescribeJobExecutionRequest returns a request value for making API operation for
// AWS IoT Jobs Data Plane.
//
// Gets details of a job execution.
//
//    // Example sending a request using the DescribeJobExecutionRequest method.
//    req := client.DescribeJobExecutionRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-jobs-data-2017-09-29/DescribeJobExecution
func (c *IoTJobsDataPlane) DescribeJobExecutionRequest(input *DescribeJobExecutionInput) DescribeJobExecutionRequest {
	op := &aws.Operation{
		Name:       opDescribeJobExecution,
		HTTPMethod: "GET",
		HTTPPath:   "/things/{thingName}/jobs/{jobId}",
	}

	if input == nil {
		input = &DescribeJobExecutionInput{}
	}

	output := &DescribeJobExecutionOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return DescribeJobExecutionRequest{Request: req, Input: input, Copy: c.DescribeJobExecutionRequest}
}

const opGetPendingJobExecutions = "GetPendingJobExecutions"

// GetPendingJobExecutionsRequest is a API request type for the GetPendingJobExecutions API operation.
type GetPendingJobExecutionsRequest struct {
	*aws.Request
	Input *GetPendingJobExecutionsInput
	Copy  func(*GetPendingJobExecutionsInput) GetPendingJobExecutionsRequest
}

// Send marshals and sends the GetPendingJobExecutions API request.
func (r GetPendingJobExecutionsRequest) Send() (*GetPendingJobExecutionsOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*GetPendingJobExecutionsOutput), nil
}

// GetPendingJobExecutionsRequest returns a request value for making API operation for
// AWS IoT Jobs Data Plane.
//
// Gets the list of all jobs for a thing that are not in a terminal status.
//
//    // Example sending a request using the GetPendingJobExecutionsRequest method.
//    req := client.GetPendingJobExecutionsRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-jobs-data-2017-09-29/GetPendingJobExecutions
func (c *IoTJobsDataPlane) GetPendingJobExecutionsRequest(input *GetPendingJobExecutionsInput) GetPendingJobExecutionsRequest {
	op := &aws.Operation{
		Name:       opGetPendingJobExecutions,
		HTTPMethod: "GET",
		HTTPPath:   "/things/{thingName}/jobs",
	}

	if input == nil {
		input = &GetPendingJobExecutionsInput{}
	}

	output := &GetPendingJobExecutionsOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return GetPendingJobExecutionsRequest{Request: req, Input: input, Copy: c.GetPendingJobExecutionsRequest}
}

const opStartNextPendingJobExecution = "StartNextPendingJobExecution"

// StartNextPendingJobExecutionRequest is a API request type for the StartNextPendingJobExecution API operation.
type StartNextPendingJobExecutionRequest struct {
	*aws.Request
	Input *StartNextPendingJobExecutionInput
	Copy  func(*StartNextPendingJobExecutionInput) StartNextPendingJobExecutionRequest
}

// Send marshals and sends the StartNextPendingJobExecution API request.
func (r StartNextPendingJobExecutionRequest) Send() (*StartNextPendingJobExecutionOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*StartNextPendingJobExecutionOutput), nil
}

// StartNextPendingJobExecutionRequest returns a request value for making API operation for
// AWS IoT Jobs Data Plane.
//
// Gets and starts the next pending (status IN_PROGRESS or QUEUED) job execution
// for a thing.
//
//    // Example sending a request using the StartNextPendingJobExecutionRequest method.
//    req := client.StartNextPendingJobExecutionRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-jobs-data-2017-09-29/StartNextPendingJobExecution
func (c *IoTJobsDataPlane) StartNextPendingJobExecutionRequest(input *StartNextPendingJobExecutionInput) StartNextPendingJobExecutionRequest {
	op := &aws.Operation{
		Name:       opStartNextPendingJobExecution,
		HTTPMethod: "PUT",
		HTTPPath:   "/things/{thingName}/jobs/$next",
	}

	if input == nil {
		input = &StartNextPendingJobExecutionInput{}
	}

	output := &StartNextPendingJobExecutionOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return StartNextPendingJobExecutionRequest{Request: req, Input: input, Copy: c.StartNextPendingJobExecutionRequest}
}

const opUpdateJobExecution = "UpdateJobExecution"

// UpdateJobExecutionRequest is a API request type for the UpdateJobExecution API operation.
type UpdateJobExecutionRequest struct {
	*aws.Request
	Input *UpdateJobExecutionInput
	Copy  func(*UpdateJobExecutionInput) UpdateJobExecutionRequest
}

// Send marshals and sends the UpdateJobExecution API request.
func (r UpdateJobExecutionRequest) Send() (*UpdateJobExecutionOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*UpdateJobExecutionOutput), nil
}

// UpdateJobExecutionRequest returns a request value for making API operation for
// AWS IoT Jobs Data Plane.
//
// Updates the status of a job execution.
//
//    // Example sending a request using the UpdateJobExecutionRequest method.
//    req := client.UpdateJobExecutionRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-jobs-data-2017-09-29/UpdateJobExecution
func (c *IoTJobsDataPlane) UpdateJobExecutionRequest(input *UpdateJobExecutionInput) UpdateJobExecutionRequest {
	op := &aws.Operation{
		Name:       opUpdateJobExecution,
		HTTPMethod: "POST",
		HTTPPath:   "/things/{thingName}/jobs/{jobId}",
	}

	if input == nil {
		input = &UpdateJobExecutionInput{}
	}

	output := &UpdateJobExecutionOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return UpdateJobExecutionRequest{Request: req, Input: input, Copy: c.UpdateJobExecutionRequest}
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-jobs-data-2017-09-29/DescribeJobExecutionRequest
type DescribeJobExecutionInput struct {
	_ struct{} `type:"structure"`

	// Optional. A number that identifies a particular job execution on a particular
	// device. If not specified, the latest job execution is returned.
	ExecutionNumber *int64 `location:"querystring" locationName:"executionNumber" type:"long"`

	// Optional. When set to true, the response contains the job document. The default
	// is false.
	IncludeJobDocument *bool `location:"querystring" locationName:"includeJobDocument" type:"boolean"`

	// The unique identifier assigned to this job when it was created.
	//
	// JobId is a required field
	JobId *string `location:"uri" locationName:"jobId" type:"string" required:"true"`

	// The thing name associated with the device the job execution is running on.
	//
	// ThingName is a required field
	ThingName *string `location:"uri" locationName:"thingName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeJobExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeJobExecutionInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeJobExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeJobExecutionInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}

	if s.ThingName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingName"))
	}
	if s.ThingName != nil && len(*s.ThingName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeJobExecutionInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingName != nil {
		v := *s.ThingName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "thingName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ExecutionNumber != nil {
		v := *s.ExecutionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "executionNumber", protocol.Int64Value(v), metadata)
	}
	if s.IncludeJobDocument != nil {
		v := *s.IncludeJobDocument

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "includeJobDocument", protocol.BoolValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-jobs-data-2017-09-29/DescribeJobExecutionResponse
type DescribeJobExecutionOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// Contains data about a job execution.
	Execution *JobExecution `locationName:"execution" type:"structure"`
}

// String returns the string representation
func (s DescribeJobExecutionOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeJobExecutionOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s DescribeJobExecutionOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeJobExecutionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Execution != nil {
		v := s.Execution

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "execution", v, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-jobs-data-2017-09-29/GetPendingJobExecutionsRequest
type GetPendingJobExecutionsInput struct {
	_ struct{} `type:"structure"`

	// The name of the thing that is executing the job.
	//
	// ThingName is a required field
	ThingName *string `location:"uri" locationName:"thingName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPendingJobExecutionsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetPendingJobExecutionsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPendingJobExecutionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPendingJobExecutionsInput"}

	if s.ThingName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingName"))
	}
	if s.ThingName != nil && len(*s.ThingName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPendingJobExecutionsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ThingName != nil {
		v := *s.ThingName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "thingName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-jobs-data-2017-09-29/GetPendingJobExecutionsResponse
type GetPendingJobExecutionsOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// A list of JobExecutionSummary objects with status IN_PROGRESS.
	InProgressJobs []JobExecutionSummary `locationName:"inProgressJobs" type:"list"`

	// A list of JobExecutionSummary objects with status QUEUED.
	QueuedJobs []JobExecutionSummary `locationName:"queuedJobs" type:"list"`
}

// String returns the string representation
func (s GetPendingJobExecutionsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetPendingJobExecutionsOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s GetPendingJobExecutionsOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPendingJobExecutionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.InProgressJobs) > 0 {
		v := s.InProgressJobs

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "inProgressJobs", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if len(s.QueuedJobs) > 0 {
		v := s.QueuedJobs

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "queuedJobs", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

// Contains data about a job execution.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-jobs-data-2017-09-29/JobExecution
type JobExecution struct {
	_ struct{} `type:"structure"`

	// The estimated number of seconds that remain before the job execution status
	// will be changed to TIMED_OUT.
	ApproximateSecondsBeforeTimedOut *int64 `locationName:"approximateSecondsBeforeTimedOut" type:"long"`

	// A number that identifies a particular job execution on a particular device.
	// It can be used later in commands that return or update job execution information.
	ExecutionNumber *int64 `locationName:"executionNumber" type:"long"`

	// The content of the job document.
	JobDocument *string `locationName:"jobDocument" type:"string"`

	// The unique identifier you assigned to this job when it was created.
	JobId *string `locationName:"jobId" min:"1" type:"string"`

	// The time, in milliseconds since the epoch, when the job execution was last
	// updated.
	LastUpdatedAt *int64 `locationName:"lastUpdatedAt" type:"long"`

	// The time, in milliseconds since the epoch, when the job execution was enqueued.
	QueuedAt *int64 `locationName:"queuedAt" type:"long"`

	// The time, in milliseconds since the epoch, when the job execution was started.
	StartedAt *int64 `locationName:"startedAt" type:"long"`

	// The status of the job execution. Can be one of: "QUEUED", "IN_PROGRESS",
	// "FAILED", "SUCCESS", "CANCELED", "REJECTED", or "REMOVED".
	Status JobExecutionStatus `locationName:"status" type:"string" enum:"true"`

	// A collection of name/value pairs that describe the status of the job execution.
	StatusDetails map[string]string `locationName:"statusDetails" type:"map"`

	// The name of the thing that is executing the job.
	ThingName *string `locationName:"thingName" min:"1" type:"string"`

	// The version of the job execution. Job execution versions are incremented
	// each time they are updated by a device.
	VersionNumber *int64 `locationName:"versionNumber" type:"long"`
}

// String returns the string representation
func (s JobExecution) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s JobExecution) GoString() string {
	return s.String()
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s JobExecution) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApproximateSecondsBeforeTimedOut != nil {
		v := *s.ApproximateSecondsBeforeTimedOut

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "approximateSecondsBeforeTimedOut", protocol.Int64Value(v), metadata)
	}
	if s.ExecutionNumber != nil {
		v := *s.ExecutionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "executionNumber", protocol.Int64Value(v), metadata)
	}
	if s.JobDocument != nil {
		v := *s.JobDocument

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobDocument", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastUpdatedAt != nil {
		v := *s.LastUpdatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastUpdatedAt", protocol.Int64Value(v), metadata)
	}
	if s.QueuedAt != nil {
		v := *s.QueuedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "queuedAt", protocol.Int64Value(v), metadata)
	}
	if s.StartedAt != nil {
		v := *s.StartedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "startedAt", protocol.Int64Value(v), metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.StatusDetails) > 0 {
		v := s.StatusDetails

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "statusDetails", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.ThingName != nil {
		v := *s.ThingName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thingName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VersionNumber != nil {
		v := *s.VersionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "versionNumber", protocol.Int64Value(v), metadata)
	}
	return nil
}

// Contains data about the state of a job execution.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-jobs-data-2017-09-29/JobExecutionState
type JobExecutionState struct {
	_ struct{} `type:"structure"`

	// The status of the job execution. Can be one of: "QUEUED", "IN_PROGRESS",
	// "FAILED", "SUCCESS", "CANCELED", "REJECTED", or "REMOVED".
	Status JobExecutionStatus `locationName:"status" type:"string" enum:"true"`

	// A collection of name/value pairs that describe the status of the job execution.
	StatusDetails map[string]string `locationName:"statusDetails" type:"map"`

	// The version of the job execution. Job execution versions are incremented
	// each time they are updated by a device.
	VersionNumber *int64 `locationName:"versionNumber" type:"long"`
}

// String returns the string representation
func (s JobExecutionState) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s JobExecutionState) GoString() string {
	return s.String()
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s JobExecutionState) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.StatusDetails) > 0 {
		v := s.StatusDetails

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "statusDetails", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.VersionNumber != nil {
		v := *s.VersionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "versionNumber", protocol.Int64Value(v), metadata)
	}
	return nil
}

// Contains a subset of information about a job execution.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-jobs-data-2017-09-29/JobExecutionSummary
type JobExecutionSummary struct {
	_ struct{} `type:"structure"`

	// A number that identifies a particular job execution on a particular device.
	ExecutionNumber *int64 `locationName:"executionNumber" type:"long"`

	// The unique identifier you assigned to this job when it was created.
	JobId *string `locationName:"jobId" min:"1" type:"string"`

	// The time, in milliseconds since the epoch, when the job execution was last
	// updated.
	LastUpdatedAt *int64 `locationName:"lastUpdatedAt" type:"long"`

	// The time, in milliseconds since the epoch, when the job execution was enqueued.
	QueuedAt *int64 `locationName:"queuedAt" type:"long"`

	// The time, in milliseconds since the epoch, when the job execution started.
	StartedAt *int64 `locationName:"startedAt" type:"long"`

	// The version of the job execution. Job execution versions are incremented
	// each time AWS IoT Jobs receives an update from a device.
	VersionNumber *int64 `locationName:"versionNumber" type:"long"`
}

// String returns the string representation
func (s JobExecutionSummary) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s JobExecutionSummary) GoString() string {
	return s.String()
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s JobExecutionSummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.ExecutionNumber != nil {
		v := *s.ExecutionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "executionNumber", protocol.Int64Value(v), metadata)
	}
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastUpdatedAt != nil {
		v := *s.LastUpdatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastUpdatedAt", protocol.Int64Value(v), metadata)
	}
	if s.QueuedAt != nil {
		v := *s.QueuedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "queuedAt", protocol.Int64Value(v), metadata)
	}
	if s.StartedAt != nil {
		v := *s.StartedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "startedAt", protocol.Int64Value(v), metadata)
	}
	if s.VersionNumber != nil {
		v := *s.VersionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "versionNumber", protocol.Int64Value(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-jobs-data-2017-09-29/StartNextPendingJobExecutionRequest
type StartNextPendingJobExecutionInput struct {
	_ struct{} `type:"structure"`

	// A collection of name/value pairs that describe the status of the job execution.
	// If not specified, the statusDetails are unchanged.
	StatusDetails map[string]string `locationName:"statusDetails" type:"map"`

	// Specifies the amount of time this device has to finish execution of this
	// job. If the job execution status is not set to a terminal state before this
	// timer expires, or before the timer is reset (by calling UpdateJobExecution,
	// setting the status to IN_PROGRESS and specifying a new timeout value in field
	// stepTimeoutInMinutes) the job execution status will be automatically set
	// to TIMED_OUT. Note that setting this timeout has no effect on that job execution
	// timeout which may have been specified when the job was created (CreateJob
	// using field timeoutConfig).
	StepTimeoutInMinutes *int64 `locationName:"stepTimeoutInMinutes" type:"long"`

	// The name of the thing associated with the device.
	//
	// ThingName is a required field
	ThingName *string `location:"uri" locationName:"thingName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartNextPendingJobExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StartNextPendingJobExecutionInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartNextPendingJobExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartNextPendingJobExecutionInput"}

	if s.ThingName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingName"))
	}
	if s.ThingName != nil && len(*s.ThingName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartNextPendingJobExecutionInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.StatusDetails) > 0 {
		v := s.StatusDetails

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "statusDetails", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.StepTimeoutInMinutes != nil {
		v := *s.StepTimeoutInMinutes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "stepTimeoutInMinutes", protocol.Int64Value(v), metadata)
	}
	if s.ThingName != nil {
		v := *s.ThingName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "thingName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-jobs-data-2017-09-29/StartNextPendingJobExecutionResponse
type StartNextPendingJobExecutionOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// A JobExecution object.
	Execution *JobExecution `locationName:"execution" type:"structure"`
}

// String returns the string representation
func (s StartNextPendingJobExecutionOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StartNextPendingJobExecutionOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s StartNextPendingJobExecutionOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartNextPendingJobExecutionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Execution != nil {
		v := s.Execution

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "execution", v, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-jobs-data-2017-09-29/UpdateJobExecutionRequest
type UpdateJobExecutionInput struct {
	_ struct{} `type:"structure"`

	// Optional. A number that identifies a particular job execution on a particular
	// device.
	ExecutionNumber *int64 `locationName:"executionNumber" type:"long"`

	// Optional. The expected current version of the job execution. Each time you
	// update the job execution, its version is incremented. If the version of the
	// job execution stored in Jobs does not match, the update is rejected with
	// a VersionMismatch error, and an ErrorResponse that contains the current job
	// execution status data is returned. (This makes it unnecessary to perform
	// a separate DescribeJobExecution request in order to obtain the job execution
	// status data.)
	ExpectedVersion *int64 `locationName:"expectedVersion" type:"long"`

	// Optional. When set to true, the response contains the job document. The default
	// is false.
	IncludeJobDocument *bool `locationName:"includeJobDocument" type:"boolean"`

	// Optional. When included and set to true, the response contains the JobExecutionState
	// data. The default is false.
	IncludeJobExecutionState *bool `locationName:"includeJobExecutionState" type:"boolean"`

	// The unique identifier assigned to this job when it was created.
	//
	// JobId is a required field
	JobId *string `location:"uri" locationName:"jobId" min:"1" type:"string" required:"true"`

	// The new status for the job execution (IN_PROGRESS, FAILED, SUCCESS, or REJECTED).
	// This must be specified on every update.
	//
	// Status is a required field
	Status JobExecutionStatus `locationName:"status" type:"string" required:"true" enum:"true"`

	// Optional. A collection of name/value pairs that describe the status of the
	// job execution. If not specified, the statusDetails are unchanged.
	StatusDetails map[string]string `locationName:"statusDetails" type:"map"`

	// Specifies the amount of time this device has to finish execution of this
	// job. If the job execution status is not set to a terminal state before this
	// timer expires, or before the timer is reset (by again calling UpdateJobExecution,
	// setting the status to IN_PROGRESS and specifying a new timeout value in this
	// field) the job execution status will be automatically set to TIMED_OUT. Note
	// that setting or resetting this timeout has no effect on that job execution
	// timeout which may have been specified when the job was created (CreateJob
	// using field timeoutConfig).
	StepTimeoutInMinutes *int64 `locationName:"stepTimeoutInMinutes" type:"long"`

	// The name of the thing associated with the device.
	//
	// ThingName is a required field
	ThingName *string `location:"uri" locationName:"thingName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateJobExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateJobExecutionInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateJobExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateJobExecutionInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 1))
	}
	if len(s.Status) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Status"))
	}

	if s.ThingName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingName"))
	}
	if s.ThingName != nil && len(*s.ThingName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateJobExecutionInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ExecutionNumber != nil {
		v := *s.ExecutionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "executionNumber", protocol.Int64Value(v), metadata)
	}
	if s.ExpectedVersion != nil {
		v := *s.ExpectedVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "expectedVersion", protocol.Int64Value(v), metadata)
	}
	if s.IncludeJobDocument != nil {
		v := *s.IncludeJobDocument

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "includeJobDocument", protocol.BoolValue(v), metadata)
	}
	if s.IncludeJobExecutionState != nil {
		v := *s.IncludeJobExecutionState

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "includeJobExecutionState", protocol.BoolValue(v), metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.StatusDetails) > 0 {
		v := s.StatusDetails

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "statusDetails", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.StepTimeoutInMinutes != nil {
		v := *s.StepTimeoutInMinutes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "stepTimeoutInMinutes", protocol.Int64Value(v), metadata)
	}
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingName != nil {
		v := *s.ThingName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "thingName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-jobs-data-2017-09-29/UpdateJobExecutionResponse
type UpdateJobExecutionOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// A JobExecutionState object.
	ExecutionState *JobExecutionState `locationName:"executionState" type:"structure"`

	// The contents of the Job Documents.
	JobDocument *string `locationName:"jobDocument" type:"string"`
}

// String returns the string representation
func (s UpdateJobExecutionOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateJobExecutionOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s UpdateJobExecutionOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateJobExecutionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ExecutionState != nil {
		v := s.ExecutionState

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "executionState", v, metadata)
	}
	if s.JobDocument != nil {
		v := *s.JobDocument

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobDocument", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type JobExecutionStatus string

// Enum values for JobExecutionStatus
const (
	JobExecutionStatusQueued     JobExecutionStatus = "QUEUED"
	JobExecutionStatusInProgress JobExecutionStatus = "IN_PROGRESS"
	JobExecutionStatusSucceeded  JobExecutionStatus = "SUCCEEDED"
	JobExecutionStatusFailed     JobExecutionStatus = "FAILED"
	JobExecutionStatusTimedOut   JobExecutionStatus = "TIMED_OUT"
	JobExecutionStatusRejected   JobExecutionStatus = "REJECTED"
	JobExecutionStatusRemoved    JobExecutionStatus = "REMOVED"
	JobExecutionStatusCanceled   JobExecutionStatus = "CANCELED"
)

func (enum JobExecutionStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum JobExecutionStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
