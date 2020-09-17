// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the email sending status of the Amazon SES account for the current
// region. You can execute this operation no more than once per second.
func (c *Client) GetAccountSendingEnabled(ctx context.Context, params *GetAccountSendingEnabledInput, optFns ...func(*Options)) (*GetAccountSendingEnabledOutput, error) {
	stack := middleware.NewStack("GetAccountSendingEnabled", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpGetAccountSendingEnabledMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccountSendingEnabled(options.Region), middleware.Before)
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
			OperationName: "GetAccountSendingEnabled",
			Err:           err,
		}
	}
	out := result.(*GetAccountSendingEnabledOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccountSendingEnabledInput struct {
}

// Represents a request to return the email sending status for your Amazon SES
// account in the current AWS Region.
type GetAccountSendingEnabledOutput struct {
	// Describes whether email sending is enabled or disabled for your Amazon SES
	// account in the current AWS Region.
	Enabled *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpGetAccountSendingEnabledMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpGetAccountSendingEnabled{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpGetAccountSendingEnabled{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetAccountSendingEnabled(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetAccountSendingEnabled",
	}
}
