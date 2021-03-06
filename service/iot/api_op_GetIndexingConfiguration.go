// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetIndexingConfigurationInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetIndexingConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetIndexingConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	return nil
}

type GetIndexingConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The index configuration.
	ThingGroupIndexingConfiguration *ThingGroupIndexingConfiguration `locationName:"thingGroupIndexingConfiguration" type:"structure"`

	// Thing indexing configuration.
	ThingIndexingConfiguration *ThingIndexingConfiguration `locationName:"thingIndexingConfiguration" type:"structure"`
}

// String returns the string representation
func (s GetIndexingConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetIndexingConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ThingGroupIndexingConfiguration != nil {
		v := s.ThingGroupIndexingConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "thingGroupIndexingConfiguration", v, metadata)
	}
	if s.ThingIndexingConfiguration != nil {
		v := s.ThingIndexingConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "thingIndexingConfiguration", v, metadata)
	}
	return nil
}

const opGetIndexingConfiguration = "GetIndexingConfiguration"

// GetIndexingConfigurationRequest returns a request value for making API operation for
// AWS IoT.
//
// Gets the indexing configuration.
//
//    // Example sending a request using GetIndexingConfigurationRequest.
//    req := client.GetIndexingConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetIndexingConfigurationRequest(input *GetIndexingConfigurationInput) GetIndexingConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetIndexingConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/indexing/config",
	}

	if input == nil {
		input = &GetIndexingConfigurationInput{}
	}

	req := c.newRequest(op, input, &GetIndexingConfigurationOutput{})

	return GetIndexingConfigurationRequest{Request: req, Input: input, Copy: c.GetIndexingConfigurationRequest}
}

// GetIndexingConfigurationRequest is the request type for the
// GetIndexingConfiguration API operation.
type GetIndexingConfigurationRequest struct {
	*aws.Request
	Input *GetIndexingConfigurationInput
	Copy  func(*GetIndexingConfigurationInput) GetIndexingConfigurationRequest
}

// Send marshals and sends the GetIndexingConfiguration API request.
func (r GetIndexingConfigurationRequest) Send(ctx context.Context) (*GetIndexingConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetIndexingConfigurationResponse{
		GetIndexingConfigurationOutput: r.Request.Data.(*GetIndexingConfigurationOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetIndexingConfigurationResponse is the response type for the
// GetIndexingConfiguration API operation.
type GetIndexingConfigurationResponse struct {
	*GetIndexingConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetIndexingConfiguration request.
func (r *GetIndexingConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
