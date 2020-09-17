// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackage

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets details about a Channel.
func (c *Client) DescribeChannel(ctx context.Context, params *DescribeChannelInput, optFns ...func(*Options)) (*DescribeChannelOutput, error) {
	stack := middleware.NewStack("DescribeChannel", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeChannelMiddlewares(stack)
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
	addOpDescribeChannelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeChannel(options.Region), middleware.Before)
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
			OperationName: "DescribeChannel",
			Err:           err,
		}
	}
	out := result.(*DescribeChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeChannelInput struct {
	// The ID of a Channel.
	Id *string
}

type DescribeChannelOutput struct {
	// A short text description of the Channel.
	Description *string
	// A collection of tags associated with a resource
	Tags map[string]*string
	// The ID of the Channel.
	Id *string
	// An HTTP Live Streaming (HLS) ingest resource configuration.
	HlsIngest *types.HlsIngest
	// The Amazon Resource Name (ARN) assigned to the Channel.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeChannelMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeChannel{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeChannel{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeChannel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackage",
		OperationName: "DescribeChannel",
	}
}
