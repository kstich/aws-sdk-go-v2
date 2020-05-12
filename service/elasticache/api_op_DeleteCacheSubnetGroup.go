// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Represents the input of a DeleteCacheSubnetGroup operation.
type DeleteCacheSubnetGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the cache subnet group to delete.
	//
	// Constraints: Must contain no more than 255 alphanumeric characters or hyphens.
	//
	// CacheSubnetGroupName is a required field
	CacheSubnetGroupName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteCacheSubnetGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCacheSubnetGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteCacheSubnetGroupInput"}

	if s.CacheSubnetGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("CacheSubnetGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteCacheSubnetGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteCacheSubnetGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteCacheSubnetGroup = "DeleteCacheSubnetGroup"

// DeleteCacheSubnetGroupRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Deletes a cache subnet group.
//
// You cannot delete a cache subnet group if it is associated with any clusters.
//
//    // Example sending a request using DeleteCacheSubnetGroupRequest.
//    req := client.DeleteCacheSubnetGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DeleteCacheSubnetGroup
func (c *Client) DeleteCacheSubnetGroupRequest(input *DeleteCacheSubnetGroupInput) DeleteCacheSubnetGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteCacheSubnetGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteCacheSubnetGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteCacheSubnetGroupOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteCacheSubnetGroupRequest{Request: req, Input: input, Copy: c.DeleteCacheSubnetGroupRequest}
}

// DeleteCacheSubnetGroupRequest is the request type for the
// DeleteCacheSubnetGroup API operation.
type DeleteCacheSubnetGroupRequest struct {
	*aws.Request
	Input *DeleteCacheSubnetGroupInput
	Copy  func(*DeleteCacheSubnetGroupInput) DeleteCacheSubnetGroupRequest
}

// Send marshals and sends the DeleteCacheSubnetGroup API request.
func (r DeleteCacheSubnetGroupRequest) Send(ctx context.Context) (*DeleteCacheSubnetGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteCacheSubnetGroupResponse{
		DeleteCacheSubnetGroupOutput: r.Request.Data.(*DeleteCacheSubnetGroupOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteCacheSubnetGroupResponse is the response type for the
// DeleteCacheSubnetGroup API operation.
type DeleteCacheSubnetGroupResponse struct {
	*DeleteCacheSubnetGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteCacheSubnetGroup request.
func (r *DeleteCacheSubnetGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}