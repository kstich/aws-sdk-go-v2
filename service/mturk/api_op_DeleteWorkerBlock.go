// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The DeleteWorkerBlock operation allows you to reinstate a blocked Worker to work
// on your HITs. This operation reverses the effects of the CreateWorkerBlock
// operation. You need the Worker ID to use this operation. If the Worker ID is
// missing or invalid, this operation fails and returns the message “WorkerId is
// invalid.” If the specified Worker is not blocked, this operation returns
// successfully.
func (c *Client) DeleteWorkerBlock(ctx context.Context, params *DeleteWorkerBlockInput, optFns ...func(*Options)) (*DeleteWorkerBlockOutput, error) {
	stack := middleware.NewStack("DeleteWorkerBlock", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteWorkerBlockMiddlewares(stack)
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
	addOpDeleteWorkerBlockValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteWorkerBlock(options.Region), middleware.Before)
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
			OperationName: "DeleteWorkerBlock",
			Err:           err,
		}
	}
	out := result.(*DeleteWorkerBlockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteWorkerBlockInput struct {
	// A message that explains the reason for unblocking the Worker. The Worker does
	// not see this message.
	Reason *string
	// The ID of the Worker to unblock.
	WorkerId *string
}

type DeleteWorkerBlockOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteWorkerBlockMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteWorkerBlock{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteWorkerBlock{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteWorkerBlock(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "DeleteWorkerBlock",
	}
}
