// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an Amazon CloudWatch log stream from an Amazon Kinesis Data Analytics
// application.
func (c *Client) DeleteApplicationCloudWatchLoggingOption(ctx context.Context, params *DeleteApplicationCloudWatchLoggingOptionInput, optFns ...func(*Options)) (*DeleteApplicationCloudWatchLoggingOptionOutput, error) {
	stack := middleware.NewStack("DeleteApplicationCloudWatchLoggingOption", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteApplicationCloudWatchLoggingOptionMiddlewares(stack)
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
	addOpDeleteApplicationCloudWatchLoggingOptionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteApplicationCloudWatchLoggingOption(options.Region), middleware.Before)
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
			OperationName: "DeleteApplicationCloudWatchLoggingOption",
			Err:           err,
		}
	}
	out := result.(*DeleteApplicationCloudWatchLoggingOptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteApplicationCloudWatchLoggingOptionInput struct {
	// The application name.
	ApplicationName *string
	// The CloudWatchLoggingOptionId of the Amazon CloudWatch logging option to delete.
	// You can get the CloudWatchLoggingOptionId by using the DescribeApplication ()
	// operation.
	CloudWatchLoggingOptionId *string
	// The version ID of the application. You can retrieve the application version ID
	// using DescribeApplication ().
	CurrentApplicationVersionId *int64
}

type DeleteApplicationCloudWatchLoggingOptionOutput struct {
	// The version ID of the application. Kinesis Data Analytics updates the
	// ApplicationVersionId each time you change the CloudWatch logging options.
	ApplicationVersionId *int64
	// The descriptions of the remaining CloudWatch logging options for the
	// application.
	CloudWatchLoggingOptionDescriptions []*types.CloudWatchLoggingOptionDescription
	// The application's Amazon Resource Name (ARN).
	ApplicationARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteApplicationCloudWatchLoggingOptionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteApplicationCloudWatchLoggingOption{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteApplicationCloudWatchLoggingOption{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteApplicationCloudWatchLoggingOption(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "DeleteApplicationCloudWatchLoggingOption",
	}
}
