// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchain

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns detailed information about a peer node.
func (c *Client) GetNode(ctx context.Context, params *GetNodeInput, optFns ...func(*Options)) (*GetNodeOutput, error) {
	stack := middleware.NewStack("GetNode", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetNodeMiddlewares(stack)
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
	addOpGetNodeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetNode(options.Region), middleware.Before)
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
			OperationName: "GetNode",
			Err:           err,
		}
	}
	out := result.(*GetNodeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetNodeInput struct {
	// The unique identifier of the node.
	NodeId *string
	// The unique identifier of the member that owns the node.
	MemberId *string
	// The unique identifier of the network to which the node belongs.
	NetworkId *string
}

type GetNodeOutput struct {
	// Properties of the node configuration.
	Node *types.Node

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetNodeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetNode{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetNode{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetNode(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "managedblockchain",
		OperationName: "GetNode",
	}
}
