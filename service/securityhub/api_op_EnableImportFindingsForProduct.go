// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type EnableImportFindingsForProductInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the product to enable the integration for.
	//
	// ProductArn is a required field
	ProductArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s EnableImportFindingsForProductInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnableImportFindingsForProductInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EnableImportFindingsForProductInput"}

	if s.ProductArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProductArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s EnableImportFindingsForProductInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ProductArn != nil {
		v := *s.ProductArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ProductArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type EnableImportFindingsForProductOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of your subscription to the product to enable integrations for.
	ProductSubscriptionArn *string `type:"string"`
}

// String returns the string representation
func (s EnableImportFindingsForProductOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s EnableImportFindingsForProductOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ProductSubscriptionArn != nil {
		v := *s.ProductSubscriptionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ProductSubscriptionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opEnableImportFindingsForProduct = "EnableImportFindingsForProduct"

// EnableImportFindingsForProductRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Enables the integration of a partner product with Security Hub. Integrated
// products send findings to Security Hub.
//
// When you enable a product integration, a permissions policy that grants permission
// for the product to send findings to Security Hub is applied.
//
//    // Example sending a request using EnableImportFindingsForProductRequest.
//    req := client.EnableImportFindingsForProductRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/EnableImportFindingsForProduct
func (c *Client) EnableImportFindingsForProductRequest(input *EnableImportFindingsForProductInput) EnableImportFindingsForProductRequest {
	op := &aws.Operation{
		Name:       opEnableImportFindingsForProduct,
		HTTPMethod: "POST",
		HTTPPath:   "/productSubscriptions",
	}

	if input == nil {
		input = &EnableImportFindingsForProductInput{}
	}

	req := c.newRequest(op, input, &EnableImportFindingsForProductOutput{})

	return EnableImportFindingsForProductRequest{Request: req, Input: input, Copy: c.EnableImportFindingsForProductRequest}
}

// EnableImportFindingsForProductRequest is the request type for the
// EnableImportFindingsForProduct API operation.
type EnableImportFindingsForProductRequest struct {
	*aws.Request
	Input *EnableImportFindingsForProductInput
	Copy  func(*EnableImportFindingsForProductInput) EnableImportFindingsForProductRequest
}

// Send marshals and sends the EnableImportFindingsForProduct API request.
func (r EnableImportFindingsForProductRequest) Send(ctx context.Context) (*EnableImportFindingsForProductResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnableImportFindingsForProductResponse{
		EnableImportFindingsForProductOutput: r.Request.Data.(*EnableImportFindingsForProductOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnableImportFindingsForProductResponse is the response type for the
// EnableImportFindingsForProduct API operation.
type EnableImportFindingsForProductResponse struct {
	*EnableImportFindingsForProductOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnableImportFindingsForProduct request.
func (r *EnableImportFindingsForProductResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}