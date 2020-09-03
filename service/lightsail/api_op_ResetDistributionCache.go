// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Deletes currently cached content from your Amazon Lightsail content delivery
// network (CDN) distribution.  <p>After resetting the cache, the next time a
// content request is made, your distribution pulls, serves, and caches it from the
// origin.</p>
func (c *Client) ResetDistributionCache(ctx context.Context, params *ResetDistributionCacheInput, optFns ...func(*Options)) (*ResetDistributionCacheOutput, error) {
	stack := middleware.NewStack("ResetDistributionCache", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpResetDistributionCacheMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opResetDistributionCache(options.Region), middleware.Before)

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
			OperationName: "ResetDistributionCache",
			Err:           err,
		}
	}
	out := result.(*ResetDistributionCacheOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResetDistributionCacheInput struct {
	// The name of the distribution for which to reset cache.  <p>Use the
	// <code>GetDistributions</code> action to get a list of distribution names that
	// you can specify.</p>
	DistributionName *string
}

type ResetDistributionCacheOutput struct {
	// The timestamp of the reset cache request (e.g., 1479734909.17) in Unix time
	// format.
	CreateTime *time.Time
	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operation *types.Operation
	// The status of the reset cache request.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpResetDistributionCacheMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpResetDistributionCache{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpResetDistributionCache{}, middleware.After)
}

func newServiceMetadataMiddleware_opResetDistributionCache(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "ResetDistributionCache",
	}
}
