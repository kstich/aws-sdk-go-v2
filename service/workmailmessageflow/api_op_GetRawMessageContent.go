// Code generated by smithy-go-codegen DO NOT EDIT.

package workmailmessageflow

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"io"
)

// Retrieves the raw content of an in-transit email message, in MIME format.
func (c *Client) GetRawMessageContent(ctx context.Context, params *GetRawMessageContentInput, optFns ...func(*Options)) (*GetRawMessageContentOutput, error) {
	stack := middleware.NewStack("GetRawMessageContent", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetRawMessageContentMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	addOpGetRawMessageContentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRawMessageContent(options.Region), middleware.Before)
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
			OperationName: "GetRawMessageContent",
			Err:           err,
		}
	}
	out := result.(*GetRawMessageContentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRawMessageContentInput struct {
	// The identifier of the email message to retrieve.
	MessageId *string
}

type GetRawMessageContentOutput struct {
	// The raw content of the email message, in MIME format.
	MessageContent io.ReadCloser

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetRawMessageContentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetRawMessageContent{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRawMessageContent{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetRawMessageContent(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmailmessageflow",
		OperationName: "GetRawMessageContent",
	}
}
