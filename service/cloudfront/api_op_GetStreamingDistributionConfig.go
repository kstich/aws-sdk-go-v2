// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Get the configuration information about a streaming distribution.
func (c *Client) GetStreamingDistributionConfig(ctx context.Context, params *GetStreamingDistributionConfigInput, optFns ...func(*Options)) (*GetStreamingDistributionConfigOutput, error) {
	stack := middleware.NewStack("GetStreamingDistributionConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetStreamingDistributionConfigMiddlewares(stack)
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
	addOpGetStreamingDistributionConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetStreamingDistributionConfig(options.Region), middleware.Before)
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
			OperationName: "GetStreamingDistributionConfig",
			Err:           err,
		}
	}
	out := result.(*GetStreamingDistributionConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// To request to get a streaming distribution configuration.
type GetStreamingDistributionConfigInput struct {
	// The streaming distribution's ID.
	Id *string
}

// The returned result of the corresponding request.
type GetStreamingDistributionConfigOutput struct {
	// The streaming distribution's configuration information.
	StreamingDistributionConfig *types.StreamingDistributionConfig
	// The current version of the configuration. For example: E2QWRUHAPOMQZL.
	ETag *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetStreamingDistributionConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetStreamingDistributionConfig{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetStreamingDistributionConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetStreamingDistributionConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "GetStreamingDistributionConfig",
	}
}
