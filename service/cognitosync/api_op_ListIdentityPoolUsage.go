// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of identity pools registered with Cognito. ListIdentityPoolUsage can
// only be called with developer credentials. You cannot make this API call with
// the temporary user credentials provided by Cognito Identity.
func (c *Client) ListIdentityPoolUsage(ctx context.Context, params *ListIdentityPoolUsageInput, optFns ...func(*Options)) (*ListIdentityPoolUsageOutput, error) {
	stack := middleware.NewStack("ListIdentityPoolUsage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListIdentityPoolUsageMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListIdentityPoolUsage(options.Region), middleware.Before)
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
			OperationName: "ListIdentityPoolUsage",
			Err:           err,
		}
	}
	out := result.(*ListIdentityPoolUsageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request for usage information on an identity pool.
type ListIdentityPoolUsageInput struct {
	// A pagination token for obtaining the next page of results.
	NextToken *string
	// The maximum number of results to be returned.
	MaxResults *int32
}

// Returned for a successful ListIdentityPoolUsage request.
type ListIdentityPoolUsageOutput struct {
	// Usage information for the identity pools.
	IdentityPoolUsages []*types.IdentityPoolUsage
	// Total number of identities for the identity pool.
	Count *int32
	// A pagination token for obtaining the next page of results.
	NextToken *string
	// The maximum number of results to be returned.
	MaxResults *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListIdentityPoolUsageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListIdentityPoolUsage{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListIdentityPoolUsage{}, middleware.After)
}

func newServiceMetadataMiddleware_opListIdentityPoolUsage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-sync",
		OperationName: "ListIdentityPoolUsage",
	}
}
