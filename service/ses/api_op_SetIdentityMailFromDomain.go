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

// Enables or disables the custom MAIL FROM domain setup for a verified identity
// (an email address or a domain). To send emails using the specified MAIL FROM
// domain, you must add an MX record to your MAIL FROM domain's DNS settings. If
// you want your emails to pass Sender Policy Framework (SPF) checks, you must also
// add or update an SPF record. For more information, see the Amazon SES Developer
// Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/mail-from-set.html). You
// can execute this operation no more than once per second.
func (c *Client) SetIdentityMailFromDomain(ctx context.Context, params *SetIdentityMailFromDomainInput, optFns ...func(*Options)) (*SetIdentityMailFromDomainOutput, error) {
	stack := middleware.NewStack("SetIdentityMailFromDomain", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpSetIdentityMailFromDomainMiddlewares(stack)
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
	addOpSetIdentityMailFromDomainValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetIdentityMailFromDomain(options.Region), middleware.Before)

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
			OperationName: "SetIdentityMailFromDomain",
			Err:           err,
		}
	}
	out := result.(*SetIdentityMailFromDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to enable or disable the Amazon SES custom MAIL FROM domain
// setup for a verified identity. For information about using a custom MAIL FROM
// domain, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/mail-from.html).
type SetIdentityMailFromDomainInput struct {
	// The verified identity for which you want to enable or disable the specified
	// custom MAIL FROM domain.
	Identity *string
	// The action that you want Amazon SES to take if it cannot successfully read the
	// required MX record when you send an email. If you choose UseDefaultValue, Amazon
	// SES will use amazonses.com (or a subdomain of that) as the MAIL FROM domain. If
	// you choose RejectMessage, Amazon SES will return a MailFromDomainNotVerified
	// error and not send the email. The action specified in BehaviorOnMXFailure is
	// taken when the custom MAIL FROM domain setup is in the Pending, Failed, and
	// TemporaryFailure states.
	BehaviorOnMXFailure types.BehaviorOnMXFailure
	// The custom MAIL FROM domain that you want the verified identity to use. The MAIL
	// FROM domain must 1) be a subdomain of the verified identity, 2) not be used in a
	// "From" address if the MAIL FROM domain is the destination of email feedback
	// forwarding (for more information, see the Amazon SES Developer Guide
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/mail-from.html)), and 3)
	// not be used to receive emails. A value of null disables the custom MAIL FROM
	// setting for the identity.
	MailFromDomain *string
}

// An empty element returned on a successful request.
type SetIdentityMailFromDomainOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpSetIdentityMailFromDomainMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpSetIdentityMailFromDomain{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpSetIdentityMailFromDomain{}, middleware.After)
}

func newServiceMetadataMiddleware_opSetIdentityMailFromDomain(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "SetIdentityMailFromDomain",
	}
}
