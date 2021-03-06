// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package frauddetector

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutLabelInput struct {
	_ struct{} `type:"structure"`

	// The label description.
	Description *string `locationName:"description" min:"1" type:"string"`

	// The label name.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s PutLabelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutLabelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutLabelInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutLabelOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutLabelOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutLabel = "PutLabel"

// PutLabelRequest returns a request value for making API operation for
// Amazon Fraud Detector.
//
// Creates or updates label. A label classifies an event as fraudulent or legitimate.
// Labels are associated with event types and used to train supervised machine
// learning models in Amazon Fraud Detector.
//
//    // Example sending a request using PutLabelRequest.
//    req := client.PutLabelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/frauddetector-2019-11-15/PutLabel
func (c *Client) PutLabelRequest(input *PutLabelInput) PutLabelRequest {
	op := &aws.Operation{
		Name:       opPutLabel,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutLabelInput{}
	}

	req := c.newRequest(op, input, &PutLabelOutput{})

	return PutLabelRequest{Request: req, Input: input, Copy: c.PutLabelRequest}
}

// PutLabelRequest is the request type for the
// PutLabel API operation.
type PutLabelRequest struct {
	*aws.Request
	Input *PutLabelInput
	Copy  func(*PutLabelInput) PutLabelRequest
}

// Send marshals and sends the PutLabel API request.
func (r PutLabelRequest) Send(ctx context.Context) (*PutLabelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutLabelResponse{
		PutLabelOutput: r.Request.Data.(*PutLabelOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutLabelResponse is the response type for the
// PutLabel API operation.
type PutLabelResponse struct {
	*PutLabelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutLabel request.
func (r *PutLabelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
