// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteGlobalReplicationGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the Global Datastore
	//
	// GlobalReplicationGroupId is a required field
	GlobalReplicationGroupId *string `type:"string" required:"true"`

	// The primary replication group is retained as a standalone replication group.
	//
	// RetainPrimaryReplicationGroup is a required field
	RetainPrimaryReplicationGroup *bool `type:"boolean" required:"true"`
}

// String returns the string representation
func (s DeleteGlobalReplicationGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteGlobalReplicationGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteGlobalReplicationGroupInput"}

	if s.GlobalReplicationGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GlobalReplicationGroupId"))
	}

	if s.RetainPrimaryReplicationGroup == nil {
		invalidParams.Add(aws.NewErrParamRequired("RetainPrimaryReplicationGroup"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteGlobalReplicationGroupOutput struct {
	_ struct{} `type:"structure"`

	// Consists of a primary cluster that accepts writes and an associated secondary
	// cluster that resides in a different AWS region. The secondary cluster accepts
	// only reads. The primary cluster automatically replicates updates to the secondary
	// cluster.
	//
	//    * The GlobalReplicationGroupIdSuffix represents the name of the Global
	//    Datastore, which is what you use to associate a secondary cluster.
	GlobalReplicationGroup *GlobalReplicationGroup `type:"structure"`
}

// String returns the string representation
func (s DeleteGlobalReplicationGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteGlobalReplicationGroup = "DeleteGlobalReplicationGroup"

// DeleteGlobalReplicationGroupRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Deleting a Global Datastore is a two-step process:
//
//    * First, you must DisassociateGlobalReplicationGroup to remove the secondary
//    clusters in the Global Datastore.
//
//    * Once the Global Datastore contains only the primary cluster, you can
//    use DeleteGlobalReplicationGroup API to delete the Global Datastore while
//    retainining the primary cluster using Retain…= true.
//
// Since the Global Datastore has only a primary cluster, you can delete the
// Global Datastore while retaining the primary by setting RetainPrimaryCluster=true.
//
// When you receive a successful response from this operation, Amazon ElastiCache
// immediately begins deleting the selected resources; you cannot cancel or
// revert this operation.
//
//    // Example sending a request using DeleteGlobalReplicationGroupRequest.
//    req := client.DeleteGlobalReplicationGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DeleteGlobalReplicationGroup
func (c *Client) DeleteGlobalReplicationGroupRequest(input *DeleteGlobalReplicationGroupInput) DeleteGlobalReplicationGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteGlobalReplicationGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteGlobalReplicationGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteGlobalReplicationGroupOutput{})

	return DeleteGlobalReplicationGroupRequest{Request: req, Input: input, Copy: c.DeleteGlobalReplicationGroupRequest}
}

// DeleteGlobalReplicationGroupRequest is the request type for the
// DeleteGlobalReplicationGroup API operation.
type DeleteGlobalReplicationGroupRequest struct {
	*aws.Request
	Input *DeleteGlobalReplicationGroupInput
	Copy  func(*DeleteGlobalReplicationGroupInput) DeleteGlobalReplicationGroupRequest
}

// Send marshals and sends the DeleteGlobalReplicationGroup API request.
func (r DeleteGlobalReplicationGroupRequest) Send(ctx context.Context) (*DeleteGlobalReplicationGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteGlobalReplicationGroupResponse{
		DeleteGlobalReplicationGroupOutput: r.Request.Data.(*DeleteGlobalReplicationGroupOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteGlobalReplicationGroupResponse is the response type for the
// DeleteGlobalReplicationGroup API operation.
type DeleteGlobalReplicationGroupResponse struct {
	*DeleteGlobalReplicationGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteGlobalReplicationGroup request.
func (r *DeleteGlobalReplicationGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}