// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Terminates the ML compute instance. Before terminating the instance, Amazon
// SageMaker disconnects the ML storage volume from it. Amazon SageMaker preserves
// the ML storage volume. Amazon SageMaker stops charging you for the ML compute
// instance when you call StopNotebookInstance. To access data on the ML storage
// volume for a notebook instance that has been terminated, call the
// StartNotebookInstance API. StartNotebookInstance launches another ML compute
// instance, configures it, and attaches the preserved ML storage volume so you can
// continue your work.
func (c *Client) StopNotebookInstance(ctx context.Context, params *StopNotebookInstanceInput, optFns ...func(*Options)) (*StopNotebookInstanceOutput, error) {
	stack := middleware.NewStack("StopNotebookInstance", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStopNotebookInstanceMiddlewares(stack)
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
	addOpStopNotebookInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopNotebookInstance(options.Region), middleware.Before)
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
			OperationName: "StopNotebookInstance",
			Err:           err,
		}
	}
	out := result.(*StopNotebookInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopNotebookInstanceInput struct {
	// The name of the notebook instance to terminate.
	NotebookInstanceName *string
}

type StopNotebookInstanceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStopNotebookInstanceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStopNotebookInstance{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopNotebookInstance{}, middleware.After)
}

func newServiceMetadataMiddleware_opStopNotebookInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "StopNotebookInstance",
	}
}
