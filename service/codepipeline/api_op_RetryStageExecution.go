// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Resumes the pipeline execution by retrying the last failed actions in a stage.
// You can retry a stage immediately if any of the actions in the stage fail. When
// you retry, all actions that are still in progress continue working, and failed
// actions are triggered again.
func (c *Client) RetryStageExecution(ctx context.Context, params *RetryStageExecutionInput, optFns ...func(*Options)) (*RetryStageExecutionOutput, error) {
	stack := middleware.NewStack("RetryStageExecution", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRetryStageExecutionMiddlewares(stack)
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
	addOpRetryStageExecutionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRetryStageExecution(options.Region), middleware.Before)
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
			OperationName: "RetryStageExecution",
			Err:           err,
		}
	}
	out := result.(*RetryStageExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a RetryStageExecution action.
type RetryStageExecutionInput struct {
	// The ID of the pipeline execution in the failed stage to be retried. Use the
	// GetPipelineState () action to retrieve the current pipelineExecutionId of the
	// failed stage
	PipelineExecutionId *string
	// The name of the failed stage to be retried.
	StageName *string
	// The scope of the retry attempt. Currently, the only supported value is
	// FAILED_ACTIONS.
	RetryMode types.StageRetryMode
	// The name of the pipeline that contains the failed stage.
	PipelineName *string
}

// Represents the output of a RetryStageExecution action.
type RetryStageExecutionOutput struct {
	// The ID of the current workflow execution in the failed stage.
	PipelineExecutionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRetryStageExecutionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRetryStageExecution{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRetryStageExecution{}, middleware.After)
}

func newServiceMetadataMiddleware_opRetryStageExecution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "RetryStageExecution",
	}
}
