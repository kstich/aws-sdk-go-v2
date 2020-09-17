// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a cache for the GraphQL API.
func (c *Client) CreateApiCache(ctx context.Context, params *CreateApiCacheInput, optFns ...func(*Options)) (*CreateApiCacheOutput, error) {
	stack := middleware.NewStack("CreateApiCache", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateApiCacheMiddlewares(stack)
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
	addOpCreateApiCacheValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateApiCache(options.Region), middleware.Before)
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
			OperationName: "CreateApiCache",
			Err:           err,
		}
	}
	out := result.(*CreateApiCacheOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CreateApiCache operation.
type CreateApiCacheInput struct {
	// Caching behavior.
	//
	//     * FULL_REQUEST_CACHING: All requests are fully cached.
	//
	//
	// * PER_RESOLVER_CACHING: Individual resovlers that you specify are cached.
	ApiCachingBehavior types.ApiCachingBehavior
	// The cache instance type.
	//
	//     * T2_SMALL: A t2.small instance type.
	//
	//     *
	// T2_MEDIUM: A t2.medium instance type.
	//
	//     * R4_LARGE: A r4.large instance
	// type.
	//
	//     * R4_XLARGE: A r4.xlarge instance type.
	//
	//     * R4_2XLARGE: A
	// r4.2xlarge instance type.
	//
	//     * R4_4XLARGE: A r4.4xlarge instance type.
	//
	//     *
	// R4_8XLARGE: A r4.8xlarge instance type.
	Type types.ApiCacheType
	// At rest encryption flag for cache. This setting cannot be updated after
	// creation.
	AtRestEncryptionEnabled *bool
	// The GraphQL API Id.
	ApiId *string
	// Transit encryption flag when connecting to cache. This setting cannot be updated
	// after creation.
	TransitEncryptionEnabled *bool
	// TTL in seconds for cache entries. Valid values are between 1 and 3600 seconds.
	Ttl *int64
}

// Represents the output of a CreateApiCache operation.
type CreateApiCacheOutput struct {
	// The ApiCache object.
	ApiCache *types.ApiCache

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateApiCacheMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateApiCache{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateApiCache{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateApiCache(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appsync",
		OperationName: "CreateApiCache",
	}
}
