// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation configures Amazon Route 53 to automatically renew the specified
// domain before the domain registration expires. The cost of renewing your domain
// registration is billed to your AWS account. The period during which you can
// renew a domain name varies by TLD. For a list of TLDs and their renewal
// policies, see Domains That You Can Register with Amazon Route 53
// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/registrar-tld-list.html)
// in the Amazon Route 53 Developer Guide. Route 53 requires that you renew before
// the end of the renewal period so we can complete processing before the deadline.
func (c *Client) EnableDomainAutoRenew(ctx context.Context, params *EnableDomainAutoRenewInput, optFns ...func(*Options)) (*EnableDomainAutoRenewOutput, error) {
	stack := middleware.NewStack("EnableDomainAutoRenew", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpEnableDomainAutoRenewMiddlewares(stack)
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
	addOpEnableDomainAutoRenewValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableDomainAutoRenew(options.Region), middleware.Before)
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
			OperationName: "EnableDomainAutoRenew",
			Err:           err,
		}
	}
	out := result.(*EnableDomainAutoRenewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableDomainAutoRenewInput struct {
	// The name of the domain that you want to enable automatic renewal for.
	DomainName *string
}

type EnableDomainAutoRenewOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpEnableDomainAutoRenewMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpEnableDomainAutoRenew{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableDomainAutoRenew{}, middleware.After)
}

func newServiceMetadataMiddleware_opEnableDomainAutoRenew(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "EnableDomainAutoRenew",
	}
}
