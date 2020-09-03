// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Initiates a contact flow to start a new chat for the customer. Response of this
// API provides a token required to obtain credentials from the
// CreateParticipantConnection
// (https://docs.aws.amazon.com/connect-participant/latest/APIReference/API_CreateParticipantConnection.html)
// API in the Amazon Connect Participant Service.  <p>When a new chat contact is
// successfully created, clients need to subscribe to the  participant’s connection
// for the created chat within 5 minutes. This is achieved by invoking
// CreateParticipantConnection
// (https://docs.aws.amazon.com/connect-participant/latest/APIReference/API_CreateParticipantConnection.html)
// with WEBSOCKET and CONNECTION_CREDENTIALS.
func (c *Client) StartChatContact(ctx context.Context, params *StartChatContactInput, optFns ...func(*Options)) (*StartChatContactOutput, error) {
	stack := middleware.NewStack("StartChatContact", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpStartChatContactMiddlewares(stack)
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
	addIdempotencyToken_opStartChatContactMiddleware(stack, options)
	addOpStartChatContactValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartChatContact(options.Region), middleware.Before)

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
			OperationName: "StartChatContact",
			Err:           err,
		}
	}
	out := result.(*StartChatContactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartChatContactInput struct {
	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string
	// The identifier of the Amazon Connect instance.
	InstanceId *string
	// The initial message to be sent to the newly created chat.
	InitialMessage *types.ChatMessage
	// The identifier of the contact flow for the chat.
	ContactFlowId *string
	// A custom key-value pair using an attribute map. The attributes are standard
	// Amazon Connect attributes, and can be accessed in contact flows just like any
	// other contact attributes. There can be up to 32,768 UTF-8 bytes across all
	// key-value pairs per contact. Attribute keys can include only alphanumeric, dash,
	// and underscore characters.
	Attributes map[string]*string
	// Information identifying the participant.
	ParticipantDetails *types.ParticipantDetails
}

type StartChatContactOutput struct {
	// The token used by the chat participant to call CreateParticipantConnection
	// (https://docs.aws.amazon.com/connect-participant/latest/APIReference/API_CreateParticipantConnection.html).
	// The participant token is valid for the lifetime of a chat participant.
	ParticipantToken *string
	// The identifier for a chat participant. The participantId for a chat participant
	// is the same throughout the chat lifecycle.
	ParticipantId *string
	// The identifier of this contact within the Amazon Connect instance.
	ContactId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpStartChatContactMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpStartChatContact{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpStartChatContact{}, middleware.After)
}

type idempotencyToken_initializeOpStartChatContact struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartChatContact) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartChatContact) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartChatContactInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartChatContactInput ")
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
func addIdempotencyToken_opStartChatContactMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpStartChatContact{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartChatContact(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "StartChatContact",
	}
}