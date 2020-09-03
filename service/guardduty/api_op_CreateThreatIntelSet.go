// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new ThreatIntelSet. ThreatIntelSets consist of known malicious IP
// addresses. GuardDuty generates findings based on ThreatIntelSets. Only users of
// the master account can use this operation.
func (c *Client) CreateThreatIntelSet(ctx context.Context, params *CreateThreatIntelSetInput, optFns ...func(*Options)) (*CreateThreatIntelSetOutput, error) {
	stack := middleware.NewStack("CreateThreatIntelSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateThreatIntelSetMiddlewares(stack)
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
	addIdempotencyToken_opCreateThreatIntelSetMiddleware(stack, options)
	addOpCreateThreatIntelSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateThreatIntelSet(options.Region), middleware.Before)

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
			OperationName: "CreateThreatIntelSet",
			Err:           err,
		}
	}
	out := result.(*CreateThreatIntelSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateThreatIntelSetInput struct {
	// The idempotency token for the create request.
	ClientToken *string
	// The URI of the file that contains the ThreatIntelSet. For example:
	// https://s3.us-west-2.amazonaws.com/my-bucket/my-object-key.
	Location *string
	// The tags to be added to a new threat list resource.
	Tags map[string]*string
	// A user-friendly ThreatIntelSet name displayed in all findings that are generated
	// by activity that involves IP addresses included in this ThreatIntelSet.
	Name *string
	// The format of the file that contains the ThreatIntelSet.
	Format types.ThreatIntelSetFormat
	// The unique ID of the detector of the GuardDuty account that you want to create a
	// threatIntelSet for.
	DetectorId *string
	// A Boolean value that indicates whether GuardDuty is to start using the uploaded
	// ThreatIntelSet.
	Activate *bool
}

type CreateThreatIntelSetOutput struct {
	// The ID of the ThreatIntelSet resource.
	ThreatIntelSetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateThreatIntelSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateThreatIntelSet{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateThreatIntelSet{}, middleware.After)
}

type idempotencyToken_initializeOpCreateThreatIntelSet struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateThreatIntelSet) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateThreatIntelSet) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateThreatIntelSetInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateThreatIntelSetInput ")
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
func addIdempotencyToken_opCreateThreatIntelSetMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateThreatIntelSet{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateThreatIntelSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "CreateThreatIntelSet",
	}
}
