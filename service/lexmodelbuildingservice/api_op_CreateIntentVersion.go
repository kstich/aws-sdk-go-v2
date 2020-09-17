// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a new version of an intent based on the $LATEST version of the intent.
// If the $LATEST version of this intent hasn't changed since you last updated it,
// Amazon Lex doesn't create a new version. It returns the last version you
// created. You can update only the $LATEST version of the intent. You can't update
// the numbered versions that you create with the CreateIntentVersion operation.
// When you create a version of an intent, Amazon Lex sets the version to 1.
// Subsequent versions increment by 1. For more information, see versioning-intro
// (). This operation requires permissions to perform the lex:CreateIntentVersion
// action.
func (c *Client) CreateIntentVersion(ctx context.Context, params *CreateIntentVersionInput, optFns ...func(*Options)) (*CreateIntentVersionOutput, error) {
	stack := middleware.NewStack("CreateIntentVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateIntentVersionMiddlewares(stack)
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
	addOpCreateIntentVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIntentVersion(options.Region), middleware.Before)
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
			OperationName: "CreateIntentVersion",
			Err:           err,
		}
	}
	out := result.(*CreateIntentVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIntentVersionInput struct {
	// Checksum of the $LATEST version of the intent that should be used to create the
	// new version. If you specify a checksum and the $LATEST version of the intent has
	// a different checksum, Amazon Lex returns a PreconditionFailedException exception
	// and doesn't publish a new version. If you don't specify a checksum, Amazon Lex
	// publishes the $LATEST version.
	Checksum *string
	// The name of the intent that you want to create a new version of. The name is
	// case sensitive.
	Name *string
}

type CreateIntentVersionOutput struct {
	// Configuration information, if any, for connectin an Amazon Kendra index with the
	// AMAZON.KendraSearchIntent intent.
	KendraConfiguration *types.KendraConfiguration
	// Checksum of the intent version created.
	Checksum *string
	// A description of the intent.
	Description *string
	// Describes how the intent is fulfilled.
	FulfillmentActivity *types.FulfillmentActivity
	// If defined, the prompt that Amazon Lex uses to confirm the user's intent before
	// fulfilling it.
	ConfirmationPrompt *types.Prompt
	// The date that the intent was created.
	CreatedDate *time.Time
	// If the user answers "no" to the question defined in confirmationPrompt, Amazon
	// Lex responds with this statement to acknowledge that the intent was canceled.
	RejectionStatement *types.Statement
	// A unique identifier for a built-in intent.
	ParentIntentSignature *string
	// An array of sample utterances configured for the intent.
	SampleUtterances []*string
	// If defined, Amazon Lex invokes this Lambda function for each user input.
	DialogCodeHook *types.CodeHook
	// An array of slot types that defines the information required to fulfill the
	// intent.
	Slots []*types.Slot
	// The name of the intent.
	Name *string
	// After the Lambda function specified in the fulfillmentActivity field fulfills
	// the intent, Amazon Lex conveys this statement to the user.
	ConclusionStatement *types.Statement
	// The date that the intent was updated.
	LastUpdatedDate *time.Time
	// If defined, Amazon Lex uses this prompt to solicit additional user activity
	// after the intent is fulfilled.
	FollowUpPrompt *types.FollowUpPrompt
	// The version number assigned to the new version of the intent.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateIntentVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateIntentVersion{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateIntentVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateIntentVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "CreateIntentVersion",
	}
}
