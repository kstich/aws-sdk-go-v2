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

// Returns the custom email verification template for the template name you
// specify. For more information about custom verification email templates, see
// Using Custom Verification Email Templates
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/custom-verification-emails.html)
// in the Amazon SES Developer Guide. You can execute this operation no more than
// once per second.
func (c *Client) GetCustomVerificationEmailTemplate(ctx context.Context, params *GetCustomVerificationEmailTemplateInput, optFns ...func(*Options)) (*GetCustomVerificationEmailTemplateOutput, error) {
	stack := middleware.NewStack("GetCustomVerificationEmailTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpGetCustomVerificationEmailTemplateMiddlewares(stack)
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
	addOpGetCustomVerificationEmailTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCustomVerificationEmailTemplate(options.Region), middleware.Before)
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
			OperationName: "GetCustomVerificationEmailTemplate",
			Err:           err,
		}
	}
	out := result.(*GetCustomVerificationEmailTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to retrieve an existing custom verification email template.
type GetCustomVerificationEmailTemplateInput struct {
	// The name of the custom verification email template that you want to retrieve.
	TemplateName *string
}

// The content of the custom verification email template.
type GetCustomVerificationEmailTemplateOutput struct {
	// The email address that the custom verification email is sent from.
	FromEmailAddress *string
	// The URL that the recipient of the verification email is sent to if his or her
	// address is not successfully verified.
	FailureRedirectionURL *string
	// The URL that the recipient of the verification email is sent to if his or her
	// address is successfully verified.
	SuccessRedirectionURL *string
	// The content of the custom verification email.
	TemplateContent *string
	// The subject line of the custom verification email.
	TemplateSubject *string
	// The name of the custom verification email template.
	TemplateName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpGetCustomVerificationEmailTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpGetCustomVerificationEmailTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpGetCustomVerificationEmailTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetCustomVerificationEmailTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetCustomVerificationEmailTemplate",
	}
}
