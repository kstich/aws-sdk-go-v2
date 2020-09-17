// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops an in-progress operation on a stack set and its associated stack
// instances.
func (c *Client) StopStackSetOperation(ctx context.Context, params *StopStackSetOperationInput, optFns ...func(*Options)) (*StopStackSetOperationOutput, error) {
	stack := middleware.NewStack("StopStackSetOperation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpStopStackSetOperationMiddlewares(stack)
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
	addOpStopStackSetOperationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopStackSetOperation(options.Region), middleware.Before)
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
			OperationName: "StopStackSetOperation",
			Err:           err,
		}
	}
	out := result.(*StopStackSetOperationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopStackSetOperationInput struct {
	// The ID of the stack operation.
	OperationId *string
	// The name or unique ID of the stack set that you want to stop the operation for.
	StackSetName *string
}

type StopStackSetOperationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpStopStackSetOperationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpStopStackSetOperation{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpStopStackSetOperation{}, middleware.After)
}

func newServiceMetadataMiddleware_opStopStackSetOperation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "StopStackSetOperation",
	}
}
