// Code generated by smithy-go-codegen DO NOT EDIT.

package iotjobsdataplane

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotjobsdataplane/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the status of a job execution.
func (c *Client) UpdateJobExecution(ctx context.Context, params *UpdateJobExecutionInput, optFns ...func(*Options)) (*UpdateJobExecutionOutput, error) {
	stack := middleware.NewStack("UpdateJobExecution", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateJobExecutionMiddlewares(stack)
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
	addOpUpdateJobExecutionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateJobExecution(options.Region), middleware.Before)

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
			OperationName: "UpdateJobExecution",
			Err:           err,
		}
	}
	out := result.(*UpdateJobExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateJobExecutionInput struct {
	// The unique identifier assigned to this job when it was created.
	JobId *string
	// The new status for the job execution (IN_PROGRESS, FAILED, SUCCESS, or
	// REJECTED). This must be specified on every update.
	Status types.JobExecutionStatus
	// Optional. A collection of name/value pairs that describe the status of the job
	// execution. If not specified, the statusDetails are unchanged.
	StatusDetails map[string]*string
	// Optional. The expected current version of the job execution. Each time you
	// update the job execution, its version is incremented. If the version of the job
	// execution stored in Jobs does not match, the update is rejected with a
	// VersionMismatch error, and an ErrorResponse that contains the current job
	// execution status data is returned. (This makes it unnecessary to perform a
	// separate DescribeJobExecution request in order to obtain the job execution
	// status data.)
	ExpectedVersion *int64
	// Optional. When included and set to true, the response contains the
	// JobExecutionState data. The default is false.
	IncludeJobExecutionState *bool
	// Optional. When set to true, the response contains the job document. The default
	// is false.
	IncludeJobDocument *bool
	// Specifies the amount of time this device has to finish execution of this job. If
	// the job execution status is not set to a terminal state before this timer
	// expires, or before the timer is reset (by again calling UpdateJobExecution,
	// setting the status to IN_PROGRESS and specifying a new timeout value in this
	// field) the job execution status will be automatically set to TIMED_OUT. Note
	// that setting or resetting this timeout has no effect on that job execution
	// timeout which may have been specified when the job was created (CreateJob using
	// field timeoutConfig).
	StepTimeoutInMinutes *int64
	// Optional. A number that identifies a particular job execution on a particular
	// device.
	ExecutionNumber *int64
	// The name of the thing associated with the device.
	ThingName *string
}

type UpdateJobExecutionOutput struct {
	// A JobExecutionState object.
	ExecutionState *types.JobExecutionState
	// The contents of the Job Documents.
	JobDocument *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateJobExecutionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateJobExecution{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateJobExecution{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateJobExecution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iot-jobs-data",
		OperationName: "UpdateJobExecution",
	}
}
