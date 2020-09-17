// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts the process of verifying an email identity. An identity is an email
// address or domain that you use when you send email. Before you can use an
// identity to send email, you first have to verify it. By verifying an identity,
// you demonstrate that you're the owner of the identity, and that you've given
// Amazon SES API v2 permission to send email from the identity. When you verify an
// email address, Amazon SES sends an email to the address. Your email address is
// verified as soon as you follow the link in the verification email.  </p> <p>When
// you verify a domain without specifying the <code>DkimSigningAttributes</code>
// object, this operation provides a set of DKIM tokens. You can convert these
// tokens into CNAME records, which you then add to the DNS configuration for your
// domain. Your domain is verified when Amazon SES detects these records in the DNS
// configuration for your domain. This verification method is known as <a
// href="https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html">Easy
// DKIM</a>.</p> <p>Alternatively, you can perform the verification process by
// providing your own public-private key pair. This verification method is known as
// Bring Your Own DKIM (BYODKIM). To use BYODKIM, your call to the
// <code>CreateEmailIdentity</code> operation has to include the
// <code>DkimSigningAttributes</code> object. When you specify this object, you
// provide a selector (a component of the DNS record name that identifies the
// public key that you want to use for DKIM authentication) and a private key.</p>
func (c *Client) CreateEmailIdentity(ctx context.Context, params *CreateEmailIdentityInput, optFns ...func(*Options)) (*CreateEmailIdentityOutput, error) {
	stack := middleware.NewStack("CreateEmailIdentity", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateEmailIdentityMiddlewares(stack)
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
	addOpCreateEmailIdentityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEmailIdentity(options.Region), middleware.Before)
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
			OperationName: "CreateEmailIdentity",
			Err:           err,
		}
	}
	out := result.(*CreateEmailIdentityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to begin the verification process for an email identity (an email
// address or domain).
type CreateEmailIdentityInput struct {
	// If your request includes this object, Amazon SES configures the identity to use
	// Bring Your Own DKIM (BYODKIM) for DKIM authentication purposes, as opposed to
	// the default method, Easy DKIM
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html). You can
	// only specify this object if the email identity is a domain, as opposed to an
	// address.
	DkimSigningAttributes *types.DkimSigningAttributes
	// The email address or domain that you want to verify.
	EmailIdentity *string
	// An array of objects that define the tags (keys and values) that you want to
	// associate with the email identity.
	Tags []*types.Tag
}

// If the email identity is a domain, this object contains information about the
// DKIM verification status for the domain. If the email identity is an email
// address, this object is empty.
type CreateEmailIdentityOutput struct {
	// Specifies whether or not the identity is verified. You can only send email from
	// verified email addresses or domains. For more information about verifying
	// identities, see the Amazon Pinpoint User Guide
	// (https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-email-manage-verify.html).
	VerifiedForSendingStatus *bool
	// An object that contains information about the DKIM attributes for the identity.
	DkimAttributes *types.DkimAttributes
	// The email identity type.
	IdentityType types.IdentityType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateEmailIdentityMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateEmailIdentity{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateEmailIdentity{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateEmailIdentity(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "CreateEmailIdentity",
	}
}
