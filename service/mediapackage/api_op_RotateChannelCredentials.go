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

// Changes the Channel's first IngestEndpoint's username and password. WARNING -
// This API is deprecated. Please use RotateIngestEndpointCredentials instead
func (c *Client) RotateChannelCredentials(ctx context.Context, params *RotateChannelCredentialsInput, optFns ...func(*Options)) (*RotateChannelCredentialsOutput, error) {
	stack := middleware.NewStack("RotateChannelCredentials", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpRotateChannelCredentialsMiddlewares(stack)
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
	addOpRotateChannelCredentialsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRotateChannelCredentials(options.Region), middleware.Before)
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
			OperationName: "RotateChannelCredentials",
			Err:           err,
		}
	}
	out := result.(*RotateChannelCredentialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RotateChannelCredentialsInput struct {
	// The ID of the channel to update.
	Id *string
}

type RotateChannelCredentialsOutput struct {
	// A short text description of the Channel.
	Description *string
	// The ID of the Channel.
	Id *string
	// A collection of tags associated with a resource
	Tags map[string]*string
	// An HTTP Live Streaming (HLS) ingest resource configuration.
	HlsIngest *types.HlsIngest
	// The Amazon Resource Name (ARN) assigned to the Channel.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpRotateChannelCredentialsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpRotateChannelCredentials{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpRotateChannelCredentials{}, middleware.After)
}

func newServiceMetadataMiddleware_opRotateChannelCredentials(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackage",
		OperationName: "RotateChannelCredentials",
	}
}
