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

// Creates a replication configuration or replaces an existing one. For more
// information, see Replication
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/replication.html) in the Amazon
// S3 Developer Guide. To perform this operation, the user or role performing the
// operation must have the iam:PassRole
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_passrole.html)
// permission. Specify the replication configuration in the request body. In the
// replication configuration, you provide the name of the destination bucket where
// you want Amazon S3 to replicate objects, the IAM role that Amazon S3 can assume
// to replicate objects on your behalf, and other relevant information.  <p>A
// replication configuration must include at least one rule, and can contain a
// maximum of 1,000. Each rule identifies a subset of objects to replicate by
// filtering the objects in the source bucket. To choose additional subsets of
// objects to replicate, add a rule for each subset. All rules must specify the
// same destination bucket.</p> <p>To specify a subset of the objects in the source
// bucket to apply a replication rule to, add the Filter element as a child of the
// Rule element. You can filter objects based on an object key prefix, one or more
// object tags, or both. When you add the Filter element in the configuration, you
// must also add the following elements: <code>DeleteMarkerReplication</code>,
// <code>Status</code>, and <code>Priority</code>.</p> <p>For information about
// enabling versioning on a bucket, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html">Using
// Versioning</a>.</p> <p>By default, a resource owner, in this case the AWS
// account that created the bucket, can perform this operation. The resource owner
// can also grant others permissions to perform the operation. For more information
// about permissions, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html">Specifying
// Permissions in a Policy</a> and <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html">Managing
// Access Permissions to Your Amazon S3 Resources</a>.</p> <p> <b>Handling
// Replication of Encrypted Objects</b> </p> <p>By default, Amazon S3 doesn't
// replicate objects that are stored at rest using server-side encryption with CMKs
// stored in AWS KMS. To replicate AWS KMS-encrypted objects, add the following:
// <code>SourceSelectionCriteria</code>, <code>SseKmsEncryptedObjects</code>,
// <code>Status</code>, <code>EncryptionConfiguration</code>, and
// <code>ReplicaKmsKeyID</code>. For information about replication configuration,
// see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-config-for-kms-objects.html">Replicating
// Objects Created with SSE Using CMKs stored in AWS KMS</a>.</p> <p>For
// information on <code>PutBucketReplication</code> errors, see
// <a>ReplicationErrorCodeList</a> </p> <p>The following operations are related to
// <code>PutBucketReplication</code>:</p> <ul> <li> <p> <a>GetBucketReplication</a>
// </p> </li> <li> <p> <a>DeleteBucketReplication</a> </p> </li> </ul>
func (c *Client) PutBucketReplication(ctx context.Context, params *PutBucketReplicationInput, optFns ...func(*Options)) (*PutBucketReplicationOutput, error) {
	stack := middleware.NewStack("PutBucketReplication", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpPutBucketReplicationMiddlewares(stack)
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
	addOpPutBucketReplicationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketReplication(options.Region), middleware.Before)
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
			OperationName: "PutBucketReplication",
			Err:           err,
		}
	}
	out := result.(*PutBucketReplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketReplicationInput struct {
	// The name of the bucket
	Bucket *string
	// A container for replication rules. You can add up to 1,000 rules. The maximum
	// size of a replication configuration is 2 MB.
	ReplicationConfiguration *types.ReplicationConfiguration
	// The base64-encoded 128-bit MD5 digest of the data. You must use this header as a
	// message integrity check to verify that the request body was not corrupted in
	// transit. For more information, see RFC 1864
	// (http://www.ietf.org/rfc/rfc1864.txt).
	ContentMD5 *string
	//
	Token *string
}

type PutBucketReplicationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpPutBucketReplicationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpPutBucketReplication{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketReplication{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutBucketReplication(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketReplication",
	}
}
