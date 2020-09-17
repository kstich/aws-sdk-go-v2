// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an InputProcessingConfiguration () from an input.
func (c *Client) DeleteApplicationInputProcessingConfiguration(ctx context.Context, params *DeleteApplicationInputProcessingConfigurationInput, optFns ...func(*Options)) (*DeleteApplicationInputProcessingConfigurationOutput, error) {
	stack := middleware.NewStack("DeleteApplicationInputProcessingConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteApplicationInputProcessingConfigurationMiddlewares(stack)
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
	addOpDeleteApplicationInputProcessingConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteApplicationInputProcessingConfiguration(options.Region), middleware.Before)
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
			OperationName: "DeleteApplicationInputProcessingConfiguration",
			Err:           err,
		}
	}
	out := result.(*DeleteApplicationInputProcessingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteApplicationInputProcessingConfigurationInput struct {
	// The ID of the input configuration from which to delete the input processing
	// configuration. You can get a list of the input IDs for an application by using
	// the DescribeApplication () operation.
	InputId *string
	// The application version. You can use the DescribeApplication () operation to get
	// the current application version. If the version specified is not the current
	// version, the ConcurrentModificationException is returned.
	CurrentApplicationVersionId *int64
	// The name of the application.
	ApplicationName *string
}

type DeleteApplicationInputProcessingConfigurationOutput struct {
	// The current application version ID.
	ApplicationVersionId *int64
	// The Amazon Resource Name (ARN) of the application.
	ApplicationARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteApplicationInputProcessingConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteApplicationInputProcessingConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteApplicationInputProcessingConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteApplicationInputProcessingConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "DeleteApplicationInputProcessingConfiguration",
	}
}
