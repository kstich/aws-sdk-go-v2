// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the existing custom verification email templates for your account in the
// current AWS Region. For more information about custom verification email
// templates, see Using Custom Verification Email Templates
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/custom-verification-emails.html)
// in the Amazon SES Developer Guide. You can execute this operation no more than
// once per second.
func (c *Client) ListCustomVerificationEmailTemplates(ctx context.Context, params *ListCustomVerificationEmailTemplatesInput, optFns ...func(*Options)) (*ListCustomVerificationEmailTemplatesOutput, error) {
	stack := middleware.NewStack("ListCustomVerificationEmailTemplates", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpListCustomVerificationEmailTemplatesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListCustomVerificationEmailTemplates(options.Region), middleware.Before)
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
			OperationName: "ListCustomVerificationEmailTemplates",
			Err:           err,
		}
	}
	out := result.(*ListCustomVerificationEmailTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to list the existing custom verification email templates
// for your account. For more information about custom verification email
// templates, see Using Custom Verification Email Templates
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/custom-verification-emails.html)
// in the Amazon SES Developer Guide.
type ListCustomVerificationEmailTemplatesInput struct {
	// An array the contains the name and creation time stamp for each template in your
	// Amazon SES account.
	NextToken *string
	// The maximum number of custom verification email templates to return. This value
	// must be at least 1 and less than or equal to 50. If you do not specify a value,
	// or if you specify a value less than 1 or greater than 50, the operation will
	// return up to 50 results.
	MaxResults *int32
}

// A paginated list of custom verification email templates.
type ListCustomVerificationEmailTemplatesOutput struct {
	// A list of the custom verification email templates that exist in your account.
	CustomVerificationEmailTemplates []*types.CustomVerificationEmailTemplate
	// A token indicating that there are additional custom verification email templates
	// available to be listed. Pass this token to a subsequent call to ListTemplates to
	// retrieve the next 50 custom verification email templates.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpListCustomVerificationEmailTemplatesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpListCustomVerificationEmailTemplates{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpListCustomVerificationEmailTemplates{}, middleware.After)
}

func newServiceMetadataMiddleware_opListCustomVerificationEmailTemplates(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListCustomVerificationEmailTemplates",
	}
}
