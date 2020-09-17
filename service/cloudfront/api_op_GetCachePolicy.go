// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a cache policy, including the following metadata:
//
//     * The policy’s
// identifier.
//
//     * The date and time when the policy was last modified.
//
// To get
// a cache policy, you must provide the policy’s identifier. If the cache policy is
// attached to a distribution’s cache behavior, you can get the policy’s identifier
// using ListDistributions or GetDistribution. If the cache policy is not attached
// to a cache behavior, you can get the identifier using ListCachePolicies.
func (c *Client) GetCachePolicy(ctx context.Context, params *GetCachePolicyInput, optFns ...func(*Options)) (*GetCachePolicyOutput, error) {
	stack := middleware.NewStack("GetCachePolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetCachePolicyMiddlewares(stack)
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
	addOpGetCachePolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCachePolicy(options.Region), middleware.Before)
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
			OperationName: "GetCachePolicy",
			Err:           err,
		}
	}
	out := result.(*GetCachePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCachePolicyInput struct {
	// The unique identifier for the cache policy. If the cache policy is attached to a
	// distribution’s cache behavior, you can get the policy’s identifier using
	// ListDistributions or GetDistribution. If the cache policy is not attached to a
	// cache behavior, you can get the identifier using ListCachePolicies.
	Id *string
}

type GetCachePolicyOutput struct {
	// The current version of the cache policy.
	ETag *string
	// The cache policy.
	CachePolicy *types.CachePolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetCachePolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetCachePolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetCachePolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetCachePolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "GetCachePolicy",
	}
}
