// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input value for ListTagsForVaultInput.
type ListTagsForVaultInput struct {
	_ struct{} `type:"structure"`

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen),
	// in which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The name of the vault.
	//
	// VaultName is a required field
	VaultName *string `location:"uri" locationName:"vaultName" type:"string" required:"true"`
}

// String returns the string representation
func (s ListTagsForVaultInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagsForVaultInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTagsForVaultInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.VaultName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VaultName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListTagsForVaultInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VaultName != nil {
		v := *s.VaultName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "vaultName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Contains the Amazon S3 Glacier response to your request.
type ListTagsForVaultOutput struct {
	_ struct{} `type:"structure"`

	// The tags attached to the vault. Each tag is composed of a key and a value.
	Tags map[string]string `type:"map"`
}

// String returns the string representation
func (s ListTagsForVaultOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListTagsForVaultOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "Tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opListTagsForVault = "ListTagsForVault"

// ListTagsForVaultRequest returns a request value for making API operation for
// Amazon Glacier.
//
// This operation lists all the tags attached to a vault. The operation returns
// an empty map if there are no tags. For more information about tags, see Tagging
// Amazon S3 Glacier Resources (https://docs.aws.amazon.com/amazonglacier/latest/dev/tagging.html).
//
//    // Example sending a request using ListTagsForVaultRequest.
//    req := client.ListTagsForVaultRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListTagsForVaultRequest(input *ListTagsForVaultInput) ListTagsForVaultRequest {
	op := &aws.Operation{
		Name:       opListTagsForVault,
		HTTPMethod: "GET",
		HTTPPath:   "/{accountId}/vaults/{vaultName}/tags",
	}

	if input == nil {
		input = &ListTagsForVaultInput{}
	}

	req := c.newRequest(op, input, &ListTagsForVaultOutput{})

	return ListTagsForVaultRequest{Request: req, Input: input, Copy: c.ListTagsForVaultRequest}
}

// ListTagsForVaultRequest is the request type for the
// ListTagsForVault API operation.
type ListTagsForVaultRequest struct {
	*aws.Request
	Input *ListTagsForVaultInput
	Copy  func(*ListTagsForVaultInput) ListTagsForVaultRequest
}

// Send marshals and sends the ListTagsForVault API request.
func (r ListTagsForVaultRequest) Send(ctx context.Context) (*ListTagsForVaultResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTagsForVaultResponse{
		ListTagsForVaultOutput: r.Request.Data.(*ListTagsForVaultOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListTagsForVaultResponse is the response type for the
// ListTagsForVault API operation.
type ListTagsForVaultResponse struct {
	*ListTagsForVaultOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTagsForVault request.
func (r *ListTagsForVaultResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}