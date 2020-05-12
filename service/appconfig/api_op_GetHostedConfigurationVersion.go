// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appconfig

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetHostedConfigurationVersionInput struct {
	_ struct{} `type:"structure"`

	// The application ID.
	//
	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"ApplicationId" type:"string" required:"true"`

	// The configuration profile ID.
	//
	// ConfigurationProfileId is a required field
	ConfigurationProfileId *string `location:"uri" locationName:"ConfigurationProfileId" type:"string" required:"true"`

	// The version.
	//
	// VersionNumber is a required field
	VersionNumber *int64 `location:"uri" locationName:"VersionNumber" type:"integer" required:"true"`
}

// String returns the string representation
func (s GetHostedConfigurationVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetHostedConfigurationVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetHostedConfigurationVersionInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.ConfigurationProfileId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationProfileId"))
	}

	if s.VersionNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("VersionNumber"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetHostedConfigurationVersionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ApplicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ConfigurationProfileId != nil {
		v := *s.ConfigurationProfileId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ConfigurationProfileId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VersionNumber != nil {
		v := *s.VersionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "VersionNumber", protocol.Int64Value(v), metadata)
	}
	return nil
}

type GetHostedConfigurationVersionOutput struct {
	_ struct{} `type:"structure" payload:"Content"`

	// The application ID.
	ApplicationId *string `location:"header" locationName:"Application-Id" type:"string"`

	// The configuration profile ID.
	ConfigurationProfileId *string `location:"header" locationName:"Configuration-Profile-Id" type:"string"`

	// The content of the configuration or the configuration data.
	Content []byte `type:"blob" sensitive:"true"`

	// A standard MIME type describing the format of the configuration content.
	// For more information, see Content-Type (https://docs.aws.amazon.com/https:/www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17).
	ContentType *string `location:"header" locationName:"Content-Type" min:"1" type:"string"`

	// A description of the configuration.
	Description *string `location:"header" locationName:"Description" type:"string"`

	// The configuration version.
	VersionNumber *int64 `location:"header" locationName:"Version-Number" type:"integer"`
}

// String returns the string representation
func (s GetHostedConfigurationVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetHostedConfigurationVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Application-Id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ConfigurationProfileId != nil {
		v := *s.ConfigurationProfileId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Configuration-Profile-Id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VersionNumber != nil {
		v := *s.VersionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Version-Number", protocol.Int64Value(v), metadata)
	}
	if s.Content != nil {
		v := s.Content

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "Content", protocol.BytesStream(v), metadata)
	}
	return nil
}

const opGetHostedConfigurationVersion = "GetHostedConfigurationVersion"

// GetHostedConfigurationVersionRequest returns a request value for making API operation for
// Amazon AppConfig.
//
// Get information about a specific configuration version.
//
//    // Example sending a request using GetHostedConfigurationVersionRequest.
//    req := client.GetHostedConfigurationVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appconfig-2019-10-09/GetHostedConfigurationVersion
func (c *Client) GetHostedConfigurationVersionRequest(input *GetHostedConfigurationVersionInput) GetHostedConfigurationVersionRequest {
	op := &aws.Operation{
		Name:       opGetHostedConfigurationVersion,
		HTTPMethod: "GET",
		HTTPPath:   "/applications/{ApplicationId}/configurationprofiles/{ConfigurationProfileId}/hostedconfigurationversions/{VersionNumber}",
	}

	if input == nil {
		input = &GetHostedConfigurationVersionInput{}
	}

	req := c.newRequest(op, input, &GetHostedConfigurationVersionOutput{})

	return GetHostedConfigurationVersionRequest{Request: req, Input: input, Copy: c.GetHostedConfigurationVersionRequest}
}

// GetHostedConfigurationVersionRequest is the request type for the
// GetHostedConfigurationVersion API operation.
type GetHostedConfigurationVersionRequest struct {
	*aws.Request
	Input *GetHostedConfigurationVersionInput
	Copy  func(*GetHostedConfigurationVersionInput) GetHostedConfigurationVersionRequest
}

// Send marshals and sends the GetHostedConfigurationVersion API request.
func (r GetHostedConfigurationVersionRequest) Send(ctx context.Context) (*GetHostedConfigurationVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetHostedConfigurationVersionResponse{
		GetHostedConfigurationVersionOutput: r.Request.Data.(*GetHostedConfigurationVersionOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetHostedConfigurationVersionResponse is the response type for the
// GetHostedConfigurationVersion API operation.
type GetHostedConfigurationVersionResponse struct {
	*GetHostedConfigurationVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetHostedConfigurationVersion request.
func (r *GetHostedConfigurationVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}