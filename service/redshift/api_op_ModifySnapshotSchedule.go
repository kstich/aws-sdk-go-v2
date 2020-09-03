// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Modifies a snapshot schedule. Any schedule associated with a cluster is modified
// asynchronously.
func (c *Client) ModifySnapshotSchedule(ctx context.Context, params *ModifySnapshotScheduleInput, optFns ...func(*Options)) (*ModifySnapshotScheduleOutput, error) {
	stack := middleware.NewStack("ModifySnapshotSchedule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpModifySnapshotScheduleMiddlewares(stack)
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
	addOpModifySnapshotScheduleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifySnapshotSchedule(options.Region), middleware.Before)

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
			OperationName: "ModifySnapshotSchedule",
			Err:           err,
		}
	}
	out := result.(*ModifySnapshotScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifySnapshotScheduleInput struct {
	// A unique alphanumeric identifier of the schedule to modify.
	ScheduleIdentifier *string
	// An updated list of schedule definitions. A schedule definition is made up of
	// schedule expressions, for example, "cron(30 12 *)" or "rate(12 hours)".
	ScheduleDefinitions []*string
}

// Describes a snapshot schedule. You can set a regular interval for creating
// snapshots of a cluster. You can also schedule snapshots for specific dates.
type ModifySnapshotScheduleOutput struct {
	//
	NextInvocations []*time.Time
	// A list of clusters associated with the schedule. A maximum of 100 clusters is
	// returned.
	AssociatedClusters []*types.ClusterAssociatedToSchedule
	// The number of clusters associated with the schedule.
	AssociatedClusterCount *int32
	// A unique identifier for the schedule.
	ScheduleIdentifier *string
	// An optional set of tags describing the schedule.
	Tags []*types.Tag
	// The description of the schedule.
	ScheduleDescription *string
	// A list of ScheduleDefinitions.
	ScheduleDefinitions []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpModifySnapshotScheduleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpModifySnapshotSchedule{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpModifySnapshotSchedule{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifySnapshotSchedule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "ModifySnapshotSchedule",
	}
}
