// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates or updates a scheduled scaling action for an Auto Scaling group. If you
// leave a parameter unspecified when updating a scheduled scaling action, the
// corresponding value remains unchanged. For more information, see Scheduled
// Scaling
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/schedule_time.html) in
// the Amazon EC2 Auto Scaling User Guide.
func (c *Client) PutScheduledUpdateGroupAction(ctx context.Context, params *PutScheduledUpdateGroupActionInput, optFns ...func(*Options)) (*PutScheduledUpdateGroupActionOutput, error) {
	stack := middleware.NewStack("PutScheduledUpdateGroupAction", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpPutScheduledUpdateGroupActionMiddlewares(stack)
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
	addOpPutScheduledUpdateGroupActionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutScheduledUpdateGroupAction(options.Region), middleware.Before)
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
			OperationName: "PutScheduledUpdateGroupAction",
			Err:           err,
		}
	}
	out := result.(*PutScheduledUpdateGroupActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutScheduledUpdateGroupActionInput struct {
	// The recurring schedule for this action, in Unix cron syntax format. This format
	// consists of five fields separated by white spaces: [Minute] [Hour]
	// [Day_of_Month] [Month_of_Year] [Day_of_Week]. The value must be in quotes (for
	// example, "30 0 1 1,6,12 *"). For more information about this format, see Crontab
	// (http://crontab.org). When StartTime and EndTime are specified with Recurrence,
	// they form the boundaries of when the recurring action starts and stops.
	Recurrence *string
	// The date and time for the recurring schedule to end. Amazon EC2 Auto Scaling
	// does not perform the action after this time.
	EndTime *time.Time
	// The maximum size of the Auto Scaling group.
	MaxSize *int32
	// The minimum size of the Auto Scaling group.
	MinSize *int32
	// The name of this scaling action.
	ScheduledActionName *string
	// The name of the Auto Scaling group.
	AutoScalingGroupName *string
	// The desired capacity is the initial capacity of the Auto Scaling group after the
	// scheduled action runs and the capacity it attempts to maintain. It can scale
	// beyond this capacity if you add more scaling conditions.
	DesiredCapacity *int32
	// The date and time for this action to start, in YYYY-MM-DDThh:mm:ssZ format in
	// UTC/GMT only and in quotes (for example, "2019-06-01T00:00:00Z"). If you specify
	// Recurrence and StartTime, Amazon EC2 Auto Scaling performs the action at this
	// time, and then performs the action based on the specified recurrence. If you try
	// to schedule your action in the past, Amazon EC2 Auto Scaling returns an error
	// message.
	StartTime *time.Time
	// This parameter is no longer used.
	Time *time.Time
}

type PutScheduledUpdateGroupActionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpPutScheduledUpdateGroupActionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpPutScheduledUpdateGroupAction{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpPutScheduledUpdateGroupAction{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutScheduledUpdateGroupAction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "PutScheduledUpdateGroupAction",
	}
}
