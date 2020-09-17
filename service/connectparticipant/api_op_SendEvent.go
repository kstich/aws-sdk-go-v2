// Code generated by smithy-go-codegen DO NOT EDIT.

package connectparticipant

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sends an event. Note that ConnectionToken is used for invoking this API instead
// of ParticipantToken.
func (c *Client) SendEvent(ctx context.Context, params *SendEventInput, optFns ...func(*Options)) (*SendEventOutput, error) {
	stack := middleware.NewStack("SendEvent", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpSendEventMiddlewares(stack)
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
	addIdempotencyToken_opSendEventMiddleware(stack, options)
	addOpSendEventValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSendEvent(options.Region), middleware.Before)
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
			OperationName: "SendEvent",
			Err:           err,
		}
	}
	out := result.(*SendEventOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendEventInput struct {
	// The authentication token associated with the participant's connection.
	ConnectionToken *string
	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string
	// The content type of the request. Supported types are:  <ul> <li>
	// <p>application/vnd.amazonaws.connect.event.typing</p> </li> <li>
	// <p>application/vnd.amazonaws.connect.event.connection.acknowledged</p> </li>
	// </ul>
	ContentType *string
	// The content of the event to be sent (for example, message text). This is not yet
	// supported.
	Content *string
}

type SendEventOutput struct {
	// The ID of the response.
	Id *string
	// The time when the event was sent. It's specified in ISO 8601 format:
	// yyyy-MM-ddThh:mm:ss.SSSZ. For example, 2019-11-08T02:41:28.172Z.
	AbsoluteTime *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpSendEventMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpSendEvent{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpSendEvent{}, middleware.After)
}

type idempotencyToken_initializeOpSendEvent struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpSendEvent) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpSendEvent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*SendEventInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *SendEventInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opSendEventMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpSendEvent{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opSendEvent(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "SendEvent",
	}
}
