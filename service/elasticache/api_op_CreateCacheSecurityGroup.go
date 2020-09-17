// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new cache security group. Use a cache security group to control access
// to one or more clusters. Cache security groups are only used when you are
// creating a cluster outside of an Amazon Virtual Private Cloud (Amazon VPC). If
// you are creating a cluster inside of a VPC, use a cache subnet group instead.
// For more information, see CreateCacheSubnetGroup
// (https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CreateCacheSubnetGroup.html).
func (c *Client) CreateCacheSecurityGroup(ctx context.Context, params *CreateCacheSecurityGroupInput, optFns ...func(*Options)) (*CreateCacheSecurityGroupOutput, error) {
	stack := middleware.NewStack("CreateCacheSecurityGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateCacheSecurityGroupMiddlewares(stack)
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
	addOpCreateCacheSecurityGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCacheSecurityGroup(options.Region), middleware.Before)
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
			OperationName: "CreateCacheSecurityGroup",
			Err:           err,
		}
	}
	out := result.(*CreateCacheSecurityGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CreateCacheSecurityGroup operation.
type CreateCacheSecurityGroupInput struct {
	// A name for the cache security group. This value is stored as a lowercase string.
	// Constraints: Must contain no more than 255 alphanumeric characters. Cannot be
	// the word "Default". Example: mysecuritygroup
	CacheSecurityGroupName *string
	// A description for the cache security group.
	Description *string
}

type CreateCacheSecurityGroupOutput struct {
	// Represents the output of one of the following operations:
	//
	//     *
	// AuthorizeCacheSecurityGroupIngress
	//
	//     * CreateCacheSecurityGroup
	//
	//     *
	// RevokeCacheSecurityGroupIngress
	CacheSecurityGroup *types.CacheSecurityGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateCacheSecurityGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateCacheSecurityGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateCacheSecurityGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateCacheSecurityGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "CreateCacheSecurityGroup",
	}
}
