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

// Represents the input of a TestFailover operation which test automatic failover
// on a specified node group (called shard in the console) in a replication group
// (called cluster in the console).  <p class="title"> <b>Note the following</b>
// </p> <ul> <li> <p>A customer can use this operation to test automatic failover
// on up to 5 shards (called node groups in the ElastiCache API and AWS CLI) in any
// rolling 24-hour period.</p> </li> <li> <p>If calling this operation on shards in
// different clusters (called replication groups in the API and CLI), the calls can
// be made concurrently.</p> <p> </p> </li> <li> <p>If calling this operation
// multiple times on different shards in the same Redis (cluster mode enabled)
// replication group, the first node replacement must complete before a subsequent
// call can be made.</p> </li> <li> <p>To determine whether the node replacement is
// complete you can check Events using the Amazon ElastiCache console, the AWS CLI,
// or the ElastiCache API. Look for the following automatic failover related
// events, listed here in order of occurrance:</p> <ol> <li> <p>Replication group
// message: <code>Test Failover API called for node group <node-group-id></code>
// </p> </li> <li> <p>Cache cluster message: <code>Failover from master node
// <primary-node-id> to replica node <node-id> completed</code> </p> </li> <li>
// <p>Replication group message: <code>Failover from master node <primary-node-id>
// to replica node <node-id> completed</code> </p> </li> <li> <p>Cache cluster
// message: <code>Recovering cache nodes <node-id></code> </p> </li> <li> <p>Cache
// cluster message: <code>Finished recovery for cache nodes <node-id></code> </p>
// </li> </ol> <p>For more information see:</p> <ul> <li> <p> <a
// href="https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/ECEvents.Viewing.html">Viewing
// ElastiCache Events</a> in the <i>ElastiCache User Guide</i> </p> </li> <li> <p>
// <a
// href="https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_DescribeEvents.html">DescribeEvents</a>
// in the ElastiCache API Reference</p> </li> </ul> </li> </ul> <p>Also see, <a
// href="https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/AutoFailover.html#auto-failover-test">Testing
// Multi-AZ </a> in the <i>ElastiCache User Guide</i>.</p>
func (c *Client) TestFailover(ctx context.Context, params *TestFailoverInput, optFns ...func(*Options)) (*TestFailoverOutput, error) {
	stack := middleware.NewStack("TestFailover", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpTestFailoverMiddlewares(stack)
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
	addOpTestFailoverValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTestFailover(options.Region), middleware.Before)
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
			OperationName: "TestFailover",
			Err:           err,
		}
	}
	out := result.(*TestFailoverOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TestFailoverInput struct {
	// The name of the node group (called shard in the console) in this replication
	// group on which automatic failover is to be tested. You may test automatic
	// failover on up to 5 node groups in any rolling 24-hour period.
	NodeGroupId *string
	// The name of the replication group (console: cluster) whose automatic failover is
	// being tested by this operation.
	ReplicationGroupId *string
}

type TestFailoverOutput struct {
	// Contains all of the attributes of a specific Redis replication group.
	ReplicationGroup *types.ReplicationGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpTestFailoverMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpTestFailover{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpTestFailover{}, middleware.After)
}

func newServiceMetadataMiddleware_opTestFailover(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "TestFailover",
	}
}
