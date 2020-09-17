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

// Global Datastore for Redis offers fully managed, fast, reliable and secure
// cross-region replication. Using Global Datastore for Redis, you can create
// cross-region read replica clusters for ElastiCache for Redis to enable
// low-latency reads and disaster recovery across regions. For more information,
// see Replication Across Regions Using Global Datastore ().
//
//     * The
// GlobalReplicationGroupIdSuffix is the name of the Global Datastore.
//
//     * The
// PrimaryReplicationGroupId represents the name of the primary cluster that
// accepts writes and will replicate updates to the secondary cluster.
func (c *Client) CreateGlobalReplicationGroup(ctx context.Context, params *CreateGlobalReplicationGroupInput, optFns ...func(*Options)) (*CreateGlobalReplicationGroupOutput, error) {
	stack := middleware.NewStack("CreateGlobalReplicationGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateGlobalReplicationGroupMiddlewares(stack)
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
	addOpCreateGlobalReplicationGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGlobalReplicationGroup(options.Region), middleware.Before)
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
			OperationName: "CreateGlobalReplicationGroup",
			Err:           err,
		}
	}
	out := result.(*CreateGlobalReplicationGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGlobalReplicationGroupInput struct {
	// The suffix name of a Global Datastore. The suffix guarantees uniqueness of the
	// Global Datastore name across multiple regions.
	GlobalReplicationGroupIdSuffix *string
	// The name of the primary cluster that accepts writes and will replicate updates
	// to the secondary cluster.
	PrimaryReplicationGroupId *string
	// Provides details of the Global Datastore
	GlobalReplicationGroupDescription *string
}

type CreateGlobalReplicationGroupOutput struct {
	// Consists of a primary cluster that accepts writes and an associated secondary
	// cluster that resides in a different AWS region. The secondary cluster accepts
	// only reads. The primary cluster automatically replicates updates to the
	// secondary cluster.  <ul> <li> <p>The <b>GlobalReplicationGroupIdSuffix</b>
	// represents the name of the Global Datastore, which is what you use to associate
	// a secondary cluster.</p> </li> </ul>
	GlobalReplicationGroup *types.GlobalReplicationGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateGlobalReplicationGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateGlobalReplicationGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateGlobalReplicationGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateGlobalReplicationGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "CreateGlobalReplicationGroup",
	}
}
