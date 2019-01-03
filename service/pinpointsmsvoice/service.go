// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointsmsvoice

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// PinpointSMSVoice provides the API operation methods for making requests to
// Amazon Pinpoint SMS and Voice Service. See this package's package overview docs
// for details on the service.
//
// PinpointSMSVoice methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type PinpointSMSVoice struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*PinpointSMSVoice)

// Used for custom request initialization logic
var initRequest func(*PinpointSMSVoice, *aws.Request)

// Service information constants
const (
	ServiceName = "sms-voice.pinpoint" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName          // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the PinpointSMSVoice client with a config.
//
// Example:
//     // Create a PinpointSMSVoice client from just a config.
//     svc := pinpointsmsvoice.New(myConfig)
func New(config aws.Config) *PinpointSMSVoice {
	var signingName string
	signingName = "sms-voice"
	signingRegion := config.Region

	svc := &PinpointSMSVoice{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2018-09-05",
				JSONVersion:   "1.1",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restjson.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restjson.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restjson.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restjson.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a PinpointSMSVoice operation and runs any
// custom request initialization.
func (c *PinpointSMSVoice) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}