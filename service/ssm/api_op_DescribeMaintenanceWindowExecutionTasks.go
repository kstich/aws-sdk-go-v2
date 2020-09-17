// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// For a given maintenance window execution, lists the tasks that were run.
func (c *Client) DescribeMaintenanceWindowExecutionTasks(ctx context.Context, params *DescribeMaintenanceWindowExecutionTasksInput, optFns ...func(*Options)) (*DescribeMaintenanceWindowExecutionTasksOutput, error) {
	stack := middleware.NewStack("DescribeMaintenanceWindowExecutionTasks", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeMaintenanceWindowExecutionTasksMiddlewares(stack)
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
	addOpDescribeMaintenanceWindowExecutionTasksValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMaintenanceWindowExecutionTasks(options.Region), middleware.Before)
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
			OperationName: "DescribeMaintenanceWindowExecutionTasks",
			Err:           err,
		}
	}
	out := result.(*DescribeMaintenanceWindowExecutionTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMaintenanceWindowExecutionTasksInput struct {
	// Optional filters used to scope down the returned tasks. The supported filter key
	// is STATUS with the corresponding values PENDING, IN_PROGRESS, SUCCESS, FAILED,
	// TIMED_OUT, CANCELLING, and CANCELLED.
	Filters []*types.MaintenanceWindowFilter
	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32
	// The ID of the maintenance window execution whose task executions should be
	// retrieved.
	WindowExecutionId *string
}

type DescribeMaintenanceWindowExecutionTasksOutput struct {
	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string
	// Information about the task executions.
	WindowExecutionTaskIdentities []*types.MaintenanceWindowExecutionTaskIdentity

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeMaintenanceWindowExecutionTasksMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMaintenanceWindowExecutionTasks{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMaintenanceWindowExecutionTasks{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeMaintenanceWindowExecutionTasks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeMaintenanceWindowExecutionTasks",
	}
}
