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

// Provides the sending limits for the Amazon SES account. You can execute this
// operation no more than once per second.
func (c *Client) GetSendQuota(ctx context.Context, params *GetSendQuotaInput, optFns ...func(*Options)) (*GetSendQuotaOutput, error) {
	stack := middleware.NewStack("GetSendQuota", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpGetSendQuotaMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSendQuota(options.Region), middleware.Before)

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
			OperationName: "GetSendQuota",
			Err:           err,
		}
	}
	out := result.(*GetSendQuotaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSendQuotaInput struct {
}

// Represents your Amazon SES daily sending quota, maximum send rate, and the
// number of emails you have sent in the last 24 hours.
type GetSendQuotaOutput struct {
	// The number of emails sent during the previous 24 hours.
	SentLast24Hours *float64
	// The maximum number of emails the user is allowed to send in a 24-hour interval.
	// A value of -1 signifies an unlimited quota.
	Max24HourSend *float64
	// The maximum number of emails that Amazon SES can accept from the user's account
	// per second. The rate at which Amazon SES accepts the user's messages might be
	// less than the maximum send rate.
	MaxSendRate *float64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpGetSendQuotaMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpGetSendQuota{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpGetSendQuota{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetSendQuota(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetSendQuota",
	}
}
