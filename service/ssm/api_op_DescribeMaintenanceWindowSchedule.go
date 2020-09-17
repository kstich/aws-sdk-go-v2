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

// Retrieves information about upcoming executions of a maintenance window.
func (c *Client) DescribeMaintenanceWindowSchedule(ctx context.Context, params *DescribeMaintenanceWindowScheduleInput, optFns ...func(*Options)) (*DescribeMaintenanceWindowScheduleOutput, error) {
	stack := middleware.NewStack("DescribeMaintenanceWindowSchedule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeMaintenanceWindowScheduleMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMaintenanceWindowSchedule(options.Region), middleware.Before)
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
			OperationName: "DescribeMaintenanceWindowSchedule",
			Err:           err,
		}
	}
	out := result.(*DescribeMaintenanceWindowScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMaintenanceWindowScheduleInput struct {
	// Filters used to limit the range of results. For example, you can limit
	// maintenance window executions to only those scheduled before or after a certain
	// date and time.
	Filters []*types.PatchOrchestratorFilter
	// The instance ID or key/value pair to retrieve information about.
	Targets []*types.Target
	// The type of resource you want to retrieve information about. For example,
	// "INSTANCE".
	ResourceType types.MaintenanceWindowResourceType
	// The ID of the maintenance window to retrieve information about.
	WindowId *string
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32
	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
}

type DescribeMaintenanceWindowScheduleOutput struct {
	// Information about maintenance window executions scheduled for the specified time
	// range.
	ScheduledWindowExecutions []*types.ScheduledWindowExecution
	// The token for the next set of items to return. (You use this token in the next
	// call.)
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeMaintenanceWindowScheduleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMaintenanceWindowSchedule{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMaintenanceWindowSchedule{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeMaintenanceWindowSchedule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeMaintenanceWindowSchedule",
	}
}
