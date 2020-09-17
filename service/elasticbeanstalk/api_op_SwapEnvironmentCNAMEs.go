// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Swaps the CNAMEs of two environments.
func (c *Client) SwapEnvironmentCNAMEs(ctx context.Context, params *SwapEnvironmentCNAMEsInput, optFns ...func(*Options)) (*SwapEnvironmentCNAMEsOutput, error) {
	stack := middleware.NewStack("SwapEnvironmentCNAMEs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpSwapEnvironmentCNAMEsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opSwapEnvironmentCNAMEs(options.Region), middleware.Before)
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
			OperationName: "SwapEnvironmentCNAMEs",
			Err:           err,
		}
	}
	out := result.(*SwapEnvironmentCNAMEsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Swaps the CNAMEs of two environments.
type SwapEnvironmentCNAMEsInput struct {
	// The ID of the source environment. Condition: You must specify at least the
	// SourceEnvironmentID or the SourceEnvironmentName. You may also specify both. If
	// you specify the SourceEnvironmentId, you must specify the
	// DestinationEnvironmentId.
	SourceEnvironmentId *string
	// The name of the destination environment. Condition: You must specify at least
	// the DestinationEnvironmentID or the DestinationEnvironmentName. You may also
	// specify both. You must specify the SourceEnvironmentName with the
	// DestinationEnvironmentName.
	DestinationEnvironmentName *string
	// The ID of the destination environment. Condition: You must specify at least the
	// DestinationEnvironmentID or the DestinationEnvironmentName. You may also specify
	// both. You must specify the SourceEnvironmentId with the
	// DestinationEnvironmentId.
	DestinationEnvironmentId *string
	// The name of the source environment. Condition: You must specify at least the
	// SourceEnvironmentID or the SourceEnvironmentName. You may also specify both. If
	// you specify the SourceEnvironmentName, you must specify the
	// DestinationEnvironmentName.
	SourceEnvironmentName *string
}

type SwapEnvironmentCNAMEsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpSwapEnvironmentCNAMEsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpSwapEnvironmentCNAMEs{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpSwapEnvironmentCNAMEs{}, middleware.After)
}

func newServiceMetadataMiddleware_opSwapEnvironmentCNAMEs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "SwapEnvironmentCNAMEs",
	}
}
