// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the queued purchases for the specified Reserved Instances.
func (c *Client) DeleteQueuedReservedInstances(ctx context.Context, params *DeleteQueuedReservedInstancesInput, optFns ...func(*Options)) (*DeleteQueuedReservedInstancesOutput, error) {
	stack := middleware.NewStack("DeleteQueuedReservedInstances", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDeleteQueuedReservedInstancesMiddlewares(stack)
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
	addOpDeleteQueuedReservedInstancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteQueuedReservedInstances(options.Region), middleware.Before)
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
			OperationName: "DeleteQueuedReservedInstances",
			Err:           err,
		}
	}
	out := result.(*DeleteQueuedReservedInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteQueuedReservedInstancesInput struct {
	// The IDs of the Reserved Instances.
	ReservedInstancesIds []*string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DeleteQueuedReservedInstancesOutput struct {
	// Information about the queued purchases that were successfully deleted.
	SuccessfulQueuedPurchaseDeletions []*types.SuccessfulQueuedPurchaseDeletion
	// Information about the queued purchases that could not be deleted.
	FailedQueuedPurchaseDeletions []*types.FailedQueuedPurchaseDeletion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDeleteQueuedReservedInstancesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDeleteQueuedReservedInstances{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDeleteQueuedReservedInstances{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteQueuedReservedInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeleteQueuedReservedInstances",
	}
}
