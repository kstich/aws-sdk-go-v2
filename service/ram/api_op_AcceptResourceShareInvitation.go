// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Accepts an invitation to a resource share from another AWS account.
func (c *Client) AcceptResourceShareInvitation(ctx context.Context, params *AcceptResourceShareInvitationInput, optFns ...func(*Options)) (*AcceptResourceShareInvitationOutput, error) {
	stack := middleware.NewStack("AcceptResourceShareInvitation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpAcceptResourceShareInvitationMiddlewares(stack)
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
	addOpAcceptResourceShareInvitationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAcceptResourceShareInvitation(options.Region), middleware.Before)
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
			OperationName: "AcceptResourceShareInvitation",
			Err:           err,
		}
	}
	out := result.(*AcceptResourceShareInvitationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AcceptResourceShareInvitationInput struct {
	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string
	// The Amazon Resource Name (ARN) of the invitation.
	ResourceShareInvitationArn *string
}

type AcceptResourceShareInvitationOutput struct {
	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string
	// Information about the invitation.
	ResourceShareInvitation *types.ResourceShareInvitation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpAcceptResourceShareInvitationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpAcceptResourceShareInvitation{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpAcceptResourceShareInvitation{}, middleware.After)
}

func newServiceMetadataMiddleware_opAcceptResourceShareInvitation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "AcceptResourceShareInvitation",
	}
}
