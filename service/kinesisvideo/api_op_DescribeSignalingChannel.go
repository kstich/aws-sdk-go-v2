// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the most current information about the signaling channel. You must
// specify either the name or the Amazon Resource Name (ARN) of the channel that
// you want to describe.
func (c *Client) DescribeSignalingChannel(ctx context.Context, params *DescribeSignalingChannelInput, optFns ...func(*Options)) (*DescribeSignalingChannelOutput, error) {
	stack := middleware.NewStack("DescribeSignalingChannel", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeSignalingChannelMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSignalingChannel(options.Region), middleware.Before)

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
			OperationName: "DescribeSignalingChannel",
			Err:           err,
		}
	}
	out := result.(*DescribeSignalingChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSignalingChannelInput struct {
	// The ARN of the signaling channel that you want to describe.
	ChannelARN *string
	// The name of the signaling channel that you want to describe.
	ChannelName *string
}

type DescribeSignalingChannelOutput struct {
	// A structure that encapsulates the specified signaling channel's metadata and
	// properties.
	ChannelInfo *types.ChannelInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeSignalingChannelMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeSignalingChannel{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeSignalingChannel{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeSignalingChannel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "DescribeSignalingChannel",
	}
}
