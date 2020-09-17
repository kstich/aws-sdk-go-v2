// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of existing virtual nodes.
func (c *Client) ListVirtualNodes(ctx context.Context, params *ListVirtualNodesInput, optFns ...func(*Options)) (*ListVirtualNodesOutput, error) {
	stack := middleware.NewStack("ListVirtualNodes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListVirtualNodesMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpListVirtualNodesValidationMiddleware(stack)
	addResponseErrorWrapper(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "ListVirtualNodes",
			Err:           err,
		}
	}
	out := result.(*ListVirtualNodesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ListVirtualNodesInput struct {
	// The name of the service mesh to list virtual nodes in.
	MeshName *string
	// The nextToken value returned from a previous paginated ListVirtualNodes request
	// where limit was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value.
	NextToken *string
	// The maximum number of results returned by ListVirtualNodes in paginated output.
	// When you use this parameter, ListVirtualNodes returns only limit results in a
	// single page along with a nextToken response element. You can see the remaining
	// results of the initial request by sending another ListVirtualNodes request with
	// the returned nextToken value. This value can be between 1 and 100. If you don't
	// use this parameter, ListVirtualNodes returns up to 100 results and a nextToken
	// value if applicable.
	Limit *int32
	// The AWS IAM account ID of the service mesh owner. If the account ID is not your
	// own, then it's the ID of the account that shared the mesh with your account. For
	// more information about mesh sharing, see Working with shared meshes
	// (https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).
	MeshOwner *string
}

//
type ListVirtualNodesOutput struct {
	// The list of existing virtual nodes for the specified service mesh.
	VirtualNodes []*types.VirtualNodeRef
	// The nextToken value to include in a future ListVirtualNodes request. When the
	// results of a ListVirtualNodes request exceed limit, you can use this value to
	// retrieve the next page of results. This value is null when there are no more
	// results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListVirtualNodesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListVirtualNodes{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListVirtualNodes{}, middleware.After)
}
