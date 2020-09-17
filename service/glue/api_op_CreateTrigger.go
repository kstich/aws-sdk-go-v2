// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new trigger.
func (c *Client) CreateTrigger(ctx context.Context, params *CreateTriggerInput, optFns ...func(*Options)) (*CreateTriggerOutput, error) {
	stack := middleware.NewStack("CreateTrigger", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateTriggerMiddlewares(stack)
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
	addOpCreateTriggerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTrigger(options.Region), middleware.Before)
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
			OperationName: "CreateTrigger",
			Err:           err,
		}
	}
	out := result.(*CreateTriggerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTriggerInput struct {
	// A description of the new trigger.
	Description *string
	// Set to true to start SCHEDULED and CONDITIONAL triggers when created. True is
	// not supported for ON_DEMAND triggers.
	StartOnCreation *bool
	// A cron expression used to specify the schedule (see Time-Based Schedules for
	// Jobs and Crawlers
	// (https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html).
	// For example, to run something every day at 12:15 UTC, you would specify: cron(15
	// 12 * * ? *). This field is required when the trigger type is SCHEDULED.
	Schedule *string
	// The name of the workflow associated with the trigger.
	WorkflowName *string
	// The actions initiated by this trigger when it fires.
	Actions []*types.Action
	// The type of the new trigger.
	Type types.TriggerType
	// The tags to use with this trigger. You may use tags to limit access to the
	// trigger. For more information about tags in AWS Glue, see AWS Tags in AWS Glue
	// (https://docs.aws.amazon.com/glue/latest/dg/monitor-tags.html) in the developer
	// guide.
	Tags map[string]*string
	// The name of the trigger.
	Name *string
	// A predicate to specify when the new trigger should fire. This field is required
	// when the trigger type is CONDITIONAL.
	Predicate *types.Predicate
}

type CreateTriggerOutput struct {
	// The name of the trigger.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateTriggerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTrigger{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTrigger{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateTrigger(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "CreateTrigger",
	}
}
