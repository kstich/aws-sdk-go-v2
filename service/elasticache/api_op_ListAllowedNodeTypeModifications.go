// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all available node types that you can scale your Redis cluster's or
// replication group's current node type.  <p>When you use the
// <code>ModifyCacheCluster</code> or <code>ModifyReplicationGroup</code>
// operations to scale your cluster or replication group, the value of the
// <code>CacheNodeType</code> parameter must be one of the node types returned by
// this operation.</p>
func (c *Client) ListAllowedNodeTypeModifications(ctx context.Context, params *ListAllowedNodeTypeModificationsInput, optFns ...func(*Options)) (*ListAllowedNodeTypeModificationsOutput, error) {
	stack := middleware.NewStack("ListAllowedNodeTypeModifications", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpListAllowedNodeTypeModificationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAllowedNodeTypeModifications(options.Region), middleware.Before)
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
			OperationName: "ListAllowedNodeTypeModifications",
			Err:           err,
		}
	}
	out := result.(*ListAllowedNodeTypeModificationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input parameters for the ListAllowedNodeTypeModifications operation.
type ListAllowedNodeTypeModificationsInput struct {
	// The name of the replication group want to scale up to a larger node type.
	// ElastiCache uses the replication group id to identify the current node type
	// being used by this replication group, and from that to create a list of node
	// types you can scale up to.  <important> <p>You must provide a value for either
	// the <code>CacheClusterId</code> or the <code>ReplicationGroupId</code>.</p>
	// </important>
	ReplicationGroupId *string
	// The name of the cluster you want to scale up to a larger node instanced type.
	// ElastiCache uses the cluster id to identify the current node type of this
	// cluster and from that to create a list of node types you can scale up to.
	// <important> <p>You must provide a value for either the
	// <code>CacheClusterId</code> or the <code>ReplicationGroupId</code>.</p>
	// </important>
	CacheClusterId *string
}

// Represents the allowed node types you can use to modify your cluster or
// replication group.
type ListAllowedNodeTypeModificationsOutput struct {
	// A string list, each element of which specifies a cache node type which you can
	// use to scale your cluster or replication group. When scaling up a Redis cluster
	// or replication group using ModifyCacheCluster or ModifyReplicationGroup, use a
	// value from this list for the CacheNodeType parameter.
	ScaleUpModifications []*string
	// A string list, each element of which specifies a cache node type which you can
	// use to scale your cluster or replication group.  When scaling down a Redis
	// cluster or replication group using ModifyCacheCluster or ModifyReplicationGroup,
	// use a value from this list for the CacheNodeType parameter. </p>
	ScaleDownModifications []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpListAllowedNodeTypeModificationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpListAllowedNodeTypeModifications{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpListAllowedNodeTypeModifications{}, middleware.After)
}

func newServiceMetadataMiddleware_opListAllowedNodeTypeModifications(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "ListAllowedNodeTypeModifications",
	}
}
