// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicediscovery

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// ServiceDiscovery provides the API operation methods for making requests to
// AWS Cloud Map. See this package's package overview docs
// for details on the service.
//
// ServiceDiscovery methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type ServiceDiscovery struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*ServiceDiscovery)

// Used for custom request initialization logic
var initRequest func(*ServiceDiscovery, *aws.Request)

// Service information constants
const (
	ServiceName = "servicediscovery" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName        // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the ServiceDiscovery client with a config.
//
// Example:
//     // Create a ServiceDiscovery client from just a config.
//     svc := servicediscovery.New(myConfig)
func New(config aws.Config) *ServiceDiscovery {
	var signingName string
	signingRegion := config.Region

	svc := &ServiceDiscovery{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2017-03-14",
				JSONVersion:   "1.1",
				TargetPrefix:  "Route53AutoNaming_v20170314",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a ServiceDiscovery operation and runs any
// custom request initialization.
func (c *ServiceDiscovery) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
