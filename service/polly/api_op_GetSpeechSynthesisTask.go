// Code generated by smithy-go-codegen DO NOT EDIT.

package polly

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/polly/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a specific SpeechSynthesisTask object based on its TaskID. This object
// contains information about the given speech synthesis task, including the status
// of the task, and a link to the S3 bucket containing the output of the task.
func (c *Client) GetSpeechSynthesisTask(ctx context.Context, params *GetSpeechSynthesisTaskInput, optFns ...func(*Options)) (*GetSpeechSynthesisTaskOutput, error) {
	stack := middleware.NewStack("GetSpeechSynthesisTask", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetSpeechSynthesisTaskMiddlewares(stack)
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
	addOpGetSpeechSynthesisTaskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSpeechSynthesisTask(options.Region), middleware.Before)
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
			OperationName: "GetSpeechSynthesisTask",
			Err:           err,
		}
	}
	out := result.(*GetSpeechSynthesisTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSpeechSynthesisTaskInput struct {
	// The Amazon Polly generated identifier for a speech synthesis task.
	TaskId *string
}

type GetSpeechSynthesisTaskOutput struct {
	// SynthesisTask object that provides information from the requested task,
	// including output format, creation time, task status, and so on.
	SynthesisTask *types.SynthesisTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetSpeechSynthesisTaskMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetSpeechSynthesisTask{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSpeechSynthesisTask{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetSpeechSynthesisTask(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "polly",
		OperationName: "GetSpeechSynthesisTask",
	}
}
