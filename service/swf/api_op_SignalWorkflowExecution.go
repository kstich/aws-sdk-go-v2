// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Records a WorkflowExecutionSignaled event in the workflow execution history and
// creates a decision task for the workflow execution identified by the given
// domain, workflowId and runId. The event is recorded with the specified user
// defined signalName and input (if provided).  <note> <p>If a runId isn't
// specified, then the <code>WorkflowExecutionSignaled</code> event is recorded in
// the history of the current open workflow with the matching workflowId in the
// domain.</p> </note> <note> <p>If the specified workflow execution isn't open,
// this method fails with <code>UnknownResource</code>.</p> </note> <p> <b>Access
// Control</b> </p> <p>You can use IAM policies to control this action's access to
// Amazon SWF resources as follows:</p> <ul> <li> <p>Use a <code>Resource</code>
// element with the domain name to limit the action to only specified domains.</p>
// </li> <li> <p>Use an <code>Action</code> element to allow or deny permission to
// call this action.</p> </li> <li> <p>You cannot use an IAM policy to constrain
// this action's parameters.</p> </li> </ul> <p>If the caller doesn't have
// sufficient permissions to invoke the action, or the parameter values fall
// outside the specified constraints, the action fails. The associated event
// attribute's <code>cause</code> parameter is set to
// <code>OPERATION_NOT_PERMITTED</code>. For details and example IAM policies, see
// <a
// href="https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html">Using
// IAM to Manage Access to Amazon SWF Workflows</a> in the <i>Amazon SWF Developer
// Guide</i>.</p>
func (c *Client) SignalWorkflowExecution(ctx context.Context, params *SignalWorkflowExecutionInput, optFns ...func(*Options)) (*SignalWorkflowExecutionOutput, error) {
	stack := middleware.NewStack("SignalWorkflowExecution", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpSignalWorkflowExecutionMiddlewares(stack)
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
	addOpSignalWorkflowExecutionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSignalWorkflowExecution(options.Region), middleware.Before)

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
			OperationName: "SignalWorkflowExecution",
			Err:           err,
		}
	}
	out := result.(*SignalWorkflowExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SignalWorkflowExecutionInput struct {
	// The workflowId of the workflow execution to signal.
	WorkflowId *string
	// The name of the domain containing the workflow execution to signal.
	Domain *string
	// The name of the signal. This name must be meaningful to the target workflow.
	SignalName *string
	// Data to attach to the WorkflowExecutionSignaled event in the target workflow
	// execution's history.
	Input *string
	// The runId of the workflow execution to signal.
	RunId *string
}

type SignalWorkflowExecutionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpSignalWorkflowExecutionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpSignalWorkflowExecution{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpSignalWorkflowExecution{}, middleware.After)
}

func newServiceMetadataMiddleware_opSignalWorkflowExecution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "SignalWorkflowExecution",
	}
}
