// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconnect

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Creates a new flow. The request must include one source. The request optionally
// can include outputs (up to 50) and one entitlement.
type CreateFlowInput struct {
	_ struct{} `type:"structure"`

	// The Availability Zone that you want to create the flow in. These options
	// are limited to the Availability Zones within the current AWS Region.
	AvailabilityZone *string `locationName:"availabilityZone" type:"string"`

	// The entitlements that you want to grant on a flow.
	Entitlements []GrantEntitlementRequest `locationName:"entitlements" type:"list"`

	// The name of the flow.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// The outputs that you want to add to this flow.
	Outputs []AddOutputRequest `locationName:"outputs" type:"list"`

	// The settings for the source of the flow.
	Source *SetSourceRequest `locationName:"source" type:"structure"`

	// The settings for source failover
	SourceFailoverConfig *FailoverConfig `locationName:"sourceFailoverConfig" type:"structure"`

	Sources []SetSourceRequest `locationName:"sources" type:"list"`

	// The VPC interfaces you want on the flow.
	VpcInterfaces []VpcInterfaceRequest `locationName:"vpcInterfaces" type:"list"`
}

// String returns the string representation
func (s CreateFlowInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateFlowInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateFlowInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Entitlements != nil {
		for i, v := range s.Entitlements {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Entitlements", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Outputs != nil {
		for i, v := range s.Outputs {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Outputs", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			invalidParams.AddNested("Source", err.(aws.ErrInvalidParams))
		}
	}
	if s.Sources != nil {
		for i, v := range s.Sources {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Sources", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.VpcInterfaces != nil {
		for i, v := range s.VpcInterfaces {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "VpcInterfaces", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateFlowInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AvailabilityZone != nil {
		v := *s.AvailabilityZone

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "availabilityZone", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Entitlements != nil {
		v := s.Entitlements

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "entitlements", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Outputs != nil {
		v := s.Outputs

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "outputs", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Source != nil {
		v := s.Source

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "source", v, metadata)
	}
	if s.SourceFailoverConfig != nil {
		v := s.SourceFailoverConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "sourceFailoverConfig", v, metadata)
	}
	if s.Sources != nil {
		v := s.Sources

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "sources", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.VpcInterfaces != nil {
		v := s.VpcInterfaces

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "vpcInterfaces", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

// The result of a successful CreateFlow request.
type CreateFlowOutput struct {
	_ struct{} `type:"structure"`

	// The settings for a flow, including its source, outputs, and entitlements.
	Flow *Flow `locationName:"flow" type:"structure"`
}

// String returns the string representation
func (s CreateFlowOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateFlowOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Flow != nil {
		v := s.Flow

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "flow", v, metadata)
	}
	return nil
}

const opCreateFlow = "CreateFlow"

// CreateFlowRequest returns a request value for making API operation for
// AWS MediaConnect.
//
// Creates a new flow. The request must include one source. The request optionally
// can include outputs (up to 50) and entitlements (up to 50).
//
//    // Example sending a request using CreateFlowRequest.
//    req := client.CreateFlowRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/CreateFlow
func (c *Client) CreateFlowRequest(input *CreateFlowInput) CreateFlowRequest {
	op := &aws.Operation{
		Name:       opCreateFlow,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/flows",
	}

	if input == nil {
		input = &CreateFlowInput{}
	}

	req := c.newRequest(op, input, &CreateFlowOutput{})

	return CreateFlowRequest{Request: req, Input: input, Copy: c.CreateFlowRequest}
}

// CreateFlowRequest is the request type for the
// CreateFlow API operation.
type CreateFlowRequest struct {
	*aws.Request
	Input *CreateFlowInput
	Copy  func(*CreateFlowInput) CreateFlowRequest
}

// Send marshals and sends the CreateFlow API request.
func (r CreateFlowRequest) Send(ctx context.Context) (*CreateFlowResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateFlowResponse{
		CreateFlowOutput: r.Request.Data.(*CreateFlowOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateFlowResponse is the response type for the
// CreateFlow API operation.
type CreateFlowResponse struct {
	*CreateFlowOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateFlow request.
func (r *CreateFlowResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}