// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateAssociationBatchInput struct {
	_ struct{} `type:"structure"`

	// One or more associations.
	//
	// Entries is a required field
	Entries []CreateAssociationBatchRequestEntry `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s CreateAssociationBatchInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAssociationBatchInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateAssociationBatchInput"}

	if s.Entries == nil {
		invalidParams.Add(aws.NewErrParamRequired("Entries"))
	}
	if s.Entries != nil && len(s.Entries) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Entries", 1))
	}
	if s.Entries != nil {
		for i, v := range s.Entries {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Entries", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateAssociationBatchOutput struct {
	_ struct{} `type:"structure"`

	// Information about the associations that failed.
	Failed []FailedCreateAssociation `type:"list"`

	// Information about the associations that succeeded.
	Successful []AssociationDescription `type:"list"`
}

// String returns the string representation
func (s CreateAssociationBatchOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateAssociationBatch = "CreateAssociationBatch"

// CreateAssociationBatchRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Associates the specified Systems Manager document with the specified instances
// or targets.
//
// When you associate a document with one or more instances using instance IDs
// or tags, SSM Agent running on the instance processes the document and configures
// the instance as specified.
//
// If you associate a document with an instance that already has an associated
// document, the system returns the AssociationAlreadyExists exception.
//
//    // Example sending a request using CreateAssociationBatchRequest.
//    req := client.CreateAssociationBatchRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/CreateAssociationBatch
func (c *Client) CreateAssociationBatchRequest(input *CreateAssociationBatchInput) CreateAssociationBatchRequest {
	op := &aws.Operation{
		Name:       opCreateAssociationBatch,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateAssociationBatchInput{}
	}

	req := c.newRequest(op, input, &CreateAssociationBatchOutput{})

	return CreateAssociationBatchRequest{Request: req, Input: input, Copy: c.CreateAssociationBatchRequest}
}

// CreateAssociationBatchRequest is the request type for the
// CreateAssociationBatch API operation.
type CreateAssociationBatchRequest struct {
	*aws.Request
	Input *CreateAssociationBatchInput
	Copy  func(*CreateAssociationBatchInput) CreateAssociationBatchRequest
}

// Send marshals and sends the CreateAssociationBatch API request.
func (r CreateAssociationBatchRequest) Send(ctx context.Context) (*CreateAssociationBatchResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateAssociationBatchResponse{
		CreateAssociationBatchOutput: r.Request.Data.(*CreateAssociationBatchOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateAssociationBatchResponse is the response type for the
// CreateAssociationBatch API operation.
type CreateAssociationBatchResponse struct {
	*CreateAssociationBatchOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateAssociationBatch request.
func (r *CreateAssociationBatchResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
