// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds a new client ID (also known as audience) to the list of client IDs already
// registered for the specified IAM OpenID Connect (OIDC) provider resource. This
// operation is idempotent; it does not fail or return an error if you add an
// existing client ID to the provider.
func (c *Client) AddClientIDToOpenIDConnectProvider(ctx context.Context, params *AddClientIDToOpenIDConnectProviderInput, optFns ...func(*Options)) (*AddClientIDToOpenIDConnectProviderOutput, error) {
	stack := middleware.NewStack("AddClientIDToOpenIDConnectProvider", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpAddClientIDToOpenIDConnectProviderMiddlewares(stack)
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
	addOpAddClientIDToOpenIDConnectProviderValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddClientIDToOpenIDConnectProvider(options.Region), middleware.Before)
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
			OperationName: "AddClientIDToOpenIDConnectProvider",
			Err:           err,
		}
	}
	out := result.(*AddClientIDToOpenIDConnectProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddClientIDToOpenIDConnectProviderInput struct {
	// The client ID (also known as audience) to add to the IAM OpenID Connect provider
	// resource.
	ClientID *string
	// The Amazon Resource Name (ARN) of the IAM OpenID Connect (OIDC) provider
	// resource to add the client ID to. You can get a list of OIDC provider ARNs by
	// using the ListOpenIDConnectProviders () operation.
	OpenIDConnectProviderArn *string
}

type AddClientIDToOpenIDConnectProviderOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpAddClientIDToOpenIDConnectProviderMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpAddClientIDToOpenIDConnectProvider{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpAddClientIDToOpenIDConnectProvider{}, middleware.After)
}

func newServiceMetadataMiddleware_opAddClientIDToOpenIDConnectProvider(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "AddClientIDToOpenIDConnectProvider",
	}
}
