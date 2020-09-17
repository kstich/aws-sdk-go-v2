// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the principals associated with the specified thing. A principal can be
// X.509 certificates, IAM users, groups, and roles, Amazon Cognito identities or
// federated identities.
func (c *Client) ListThingPrincipals(ctx context.Context, params *ListThingPrincipalsInput, optFns ...func(*Options)) (*ListThingPrincipalsOutput, error) {
	stack := middleware.NewStack("ListThingPrincipals", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListThingPrincipalsMiddlewares(stack)
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
	addOpListThingPrincipalsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListThingPrincipals(options.Region), middleware.Before)
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
			OperationName: "ListThingPrincipals",
			Err:           err,
		}
	}
	out := result.(*ListThingPrincipalsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the ListThingPrincipal operation.
type ListThingPrincipalsInput struct {
	// The name of the thing.
	ThingName *string
}

// The output from the ListThingPrincipals operation.
type ListThingPrincipalsOutput struct {
	// The principals associated with the thing.
	Principals []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListThingPrincipalsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListThingPrincipals{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListThingPrincipals{}, middleware.After)
}

func newServiceMetadataMiddleware_opListThingPrincipals(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListThingPrincipals",
	}
}
