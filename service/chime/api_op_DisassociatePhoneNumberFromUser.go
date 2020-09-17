// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates the primary provisioned phone number from the specified Amazon
// Chime user.
func (c *Client) DisassociatePhoneNumberFromUser(ctx context.Context, params *DisassociatePhoneNumberFromUserInput, optFns ...func(*Options)) (*DisassociatePhoneNumberFromUserOutput, error) {
	stack := middleware.NewStack("DisassociatePhoneNumberFromUser", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDisassociatePhoneNumberFromUserMiddlewares(stack)
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
	addOpDisassociatePhoneNumberFromUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociatePhoneNumberFromUser(options.Region), middleware.Before)
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
			OperationName: "DisassociatePhoneNumberFromUser",
			Err:           err,
		}
	}
	out := result.(*DisassociatePhoneNumberFromUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociatePhoneNumberFromUserInput struct {
	// The user ID.
	UserId *string
	// The Amazon Chime account ID.
	AccountId *string
}

type DisassociatePhoneNumberFromUserOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDisassociatePhoneNumberFromUserMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDisassociatePhoneNumberFromUser{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociatePhoneNumberFromUser{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisassociatePhoneNumberFromUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "DisassociatePhoneNumberFromUser",
	}
}
