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

// This operation returns the AuthCode for the domain. To transfer a domain to
// another registrar, you provide this value to the new registrar.
func (c *Client) RetrieveDomainAuthCode(ctx context.Context, params *RetrieveDomainAuthCodeInput, optFns ...func(*Options)) (*RetrieveDomainAuthCodeOutput, error) {
	stack := middleware.NewStack("RetrieveDomainAuthCode", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRetrieveDomainAuthCodeMiddlewares(stack)
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
	addOpRetrieveDomainAuthCodeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRetrieveDomainAuthCode(options.Region), middleware.Before)
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
			OperationName: "RetrieveDomainAuthCode",
			Err:           err,
		}
	}
	out := result.(*RetrieveDomainAuthCodeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request for the authorization code for the specified domain. To transfer a
// domain to another registrar, you provide this value to the new registrar.
type RetrieveDomainAuthCodeInput struct {
	// The name of the domain that you want to get an authorization code for.
	DomainName *string
}

// The RetrieveDomainAuthCode response includes the following element.
type RetrieveDomainAuthCodeOutput struct {
	// The authorization code for the domain.
	AuthCode *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRetrieveDomainAuthCodeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRetrieveDomainAuthCode{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRetrieveDomainAuthCode{}, middleware.After)
}

func newServiceMetadataMiddleware_opRetrieveDomainAuthCode(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "RetrieveDomainAuthCode",
	}
}
