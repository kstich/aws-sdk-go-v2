// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates or updates one or more scheduled scaling actions for an Auto Scaling
// group. If you leave a parameter unspecified when updating a scheduled scaling
// action, the corresponding value remains unchanged.
func (c *Client) BatchPutScheduledUpdateGroupAction(ctx context.Context, params *BatchPutScheduledUpdateGroupActionInput, optFns ...func(*Options)) (*BatchPutScheduledUpdateGroupActionOutput, error) {
	stack := middleware.NewStack("BatchPutScheduledUpdateGroupAction", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpBatchPutScheduledUpdateGroupActionMiddlewares(stack)
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
	addOpBatchPutScheduledUpdateGroupActionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchPutScheduledUpdateGroupAction(options.Region), middleware.Before)
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
			OperationName: "BatchPutScheduledUpdateGroupAction",
			Err:           err,
		}
	}
	out := result.(*BatchPutScheduledUpdateGroupActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchPutScheduledUpdateGroupActionInput struct {
	// The name of the Auto Scaling group.
	AutoScalingGroupName *string
	// One or more scheduled actions. The maximum number allowed is 50.
	ScheduledUpdateGroupActions []*types.ScheduledUpdateGroupActionRequest
}

type BatchPutScheduledUpdateGroupActionOutput struct {
	// The names of the scheduled actions that could not be created or updated,
	// including an error message.
	FailedScheduledUpdateGroupActions []*types.FailedScheduledUpdateGroupActionRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpBatchPutScheduledUpdateGroupActionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpBatchPutScheduledUpdateGroupAction{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpBatchPutScheduledUpdateGroupAction{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchPutScheduledUpdateGroupAction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "BatchPutScheduledUpdateGroupAction",
	}
}
