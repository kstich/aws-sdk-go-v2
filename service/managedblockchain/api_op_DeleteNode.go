// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchain

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a peer node from a member that your AWS account owns. All data on the
// node is lost and cannot be recovered.
func (c *Client) DeleteNode(ctx context.Context, params *DeleteNodeInput, optFns ...func(*Options)) (*DeleteNodeOutput, error) {
	stack := middleware.NewStack("DeleteNode", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteNodeMiddlewares(stack)
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
	addOpDeleteNodeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteNode(options.Region), middleware.Before)

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
			OperationName: "DeleteNode",
			Err:           err,
		}
	}
	out := result.(*DeleteNodeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteNodeInput struct {
	// The unique identifier of the network that the node belongs to.
	NetworkId *string
	// The unique identifier of the node.
	NodeId *string
	// The unique identifier of the member that owns this node.
	MemberId *string
}

type DeleteNodeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteNodeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteNode{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteNode{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteNode(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "managedblockchain",
		OperationName: "DeleteNode",
	}
}
