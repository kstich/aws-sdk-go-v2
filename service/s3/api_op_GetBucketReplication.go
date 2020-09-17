// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the replication configuration of a bucket. It can take a while to
// propagate the put or delete a replication configuration to all Amazon S3
// systems. Therefore, a get request soon after put or delete can return a wrong
// result. For information about replication configuration, see Replication
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/replication.html) in the Amazon
// Simple Storage Service Developer Guide.  <p>This operation requires permissions
// for the <code>s3:GetReplicationConfiguration</code> action. For more information
// about permissions, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html">Using
// Bucket Policies and User Policies</a>.</p> <p>If you include the
// <code>Filter</code> element in a replication configuration, you must also
// include the <code>DeleteMarkerReplication</code> and <code>Priority</code>
// elements. The response also returns those elements.</p> <p>For information about
// <code>GetBucketReplication</code> errors, see <a>ReplicationErrorCodeList</a>
// </p> <p>The following operations are related to
// <code>GetBucketReplication</code>:</p> <ul> <li> <p> <a>PutBucketReplication</a>
// </p> </li> <li> <p> <a>DeleteBucketReplication</a> </p> </li> </ul>
func (c *Client) GetBucketReplication(ctx context.Context, params *GetBucketReplicationInput, optFns ...func(*Options)) (*GetBucketReplicationOutput, error) {
	stack := middleware.NewStack("GetBucketReplication", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetBucketReplicationMiddlewares(stack)
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
	addOpGetBucketReplicationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketReplication(options.Region), middleware.Before)
	addResponseErrorWrapper(stack)
	addUpdateEndpointMiddleware(stack, options)

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
			OperationName: "GetBucketReplication",
			Err:           err,
		}
	}
	out := result.(*GetBucketReplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketReplicationInput struct {
	// The bucket name for which to get the replication information.
	Bucket *string
}

type GetBucketReplicationOutput struct {
	// A container for replication rules. You can add up to 1,000 rules. The maximum
	// size of a replication configuration is 2 MB.
	ReplicationConfiguration *types.ReplicationConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetBucketReplicationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetBucketReplication{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketReplication{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetBucketReplication(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketReplication",
	}
}
