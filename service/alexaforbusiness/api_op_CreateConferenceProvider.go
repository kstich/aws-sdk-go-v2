// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds a new conference provider under the user's AWS account.
func (c *Client) CreateConferenceProvider(ctx context.Context, params *CreateConferenceProviderInput, optFns ...func(*Options)) (*CreateConferenceProviderOutput, error) {
	stack := middleware.NewStack("CreateConferenceProvider", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateConferenceProviderMiddlewares(stack)
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
	addIdempotencyToken_opCreateConferenceProviderMiddleware(stack, options)
	addOpCreateConferenceProviderValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConferenceProvider(options.Region), middleware.Before)
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
			OperationName: "CreateConferenceProvider",
			Err:           err,
		}
	}
	out := result.(*CreateConferenceProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateConferenceProviderInput struct {
	// The IP endpoint and protocol for calling.
	IPDialIn *types.IPDialIn
	// The request token of the client.
	ClientRequestToken *string
	// The information for PSTN conferencing.
	PSTNDialIn *types.PSTNDialIn
	// The meeting settings for the conference provider.
	MeetingSetting *types.MeetingSetting
	// The name of the conference provider.
	ConferenceProviderName *string
	// Represents a type within a list of predefined types.
	ConferenceProviderType types.ConferenceProviderType
}

type CreateConferenceProviderOutput struct {
	// The ARN of the newly-created conference provider.
	ConferenceProviderArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateConferenceProviderMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateConferenceProvider{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateConferenceProvider{}, middleware.After)
}

type idempotencyToken_initializeOpCreateConferenceProvider struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateConferenceProvider) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateConferenceProvider) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateConferenceProviderInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateConferenceProviderInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateConferenceProviderMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateConferenceProvider{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateConferenceProvider(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "CreateConferenceProvider",
	}
}
