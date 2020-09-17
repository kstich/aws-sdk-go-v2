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

// Sets the tags for a bucket. Use tags to organize your AWS bill to reflect your
// own cost structure. To do this, sign up to get your AWS account bill with tag
// key values included. Then, to see the cost of combined resources, organize your
// billing information according to resources with the same tag key values. For
// example, you can tag several resources with a specific application name, and
// then organize your billing information to see the total cost of that application
// across several services. For more information, see Cost Allocation and Tagging
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html).
// <note> <p>Within a bucket, if you add a tag that has the same key as an existing
// tag, the new value overwrites the old value. For more information, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/CostAllocTagging.html">Using
// Cost Allocation in Amazon S3 Bucket Tags</a>.</p> </note> <p>To use this
// operation, you must have permissions to perform the
// <code>s3:PutBucketTagging</code> action. The bucket owner has this permission by
// default and can grant this permission to others. For more information about
// permissions, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources">Permissions
// Related to Bucket Subresource Operations</a> and <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html">Managing
// Access Permissions to Your Amazon S3 Resources</a>.</p> <p>
// <code>PutBucketTagging</code> has the following special errors:</p> <ul> <li>
// <p>Error code: <code>InvalidTagError</code> </p> <ul> <li> <p>Description: The
// tag provided was not a valid tag. This error can occur if the tag did not pass
// input validation. For information about tag restrictions, see <a
// href="https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/allocation-tag-restrictions.html">User-Defined
// Tag Restrictions</a> and <a
// href="https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/aws-tag-restrictions.html">AWS-Generated
// Cost Allocation Tag Restrictions</a>.</p> </li> </ul> </li> <li> <p>Error code:
// <code>MalformedXMLError</code> </p> <ul> <li> <p>Description: The XML provided
// does not match the schema.</p> </li> </ul> </li> <li> <p>Error code:
// <code>OperationAbortedError </code> </p> <ul> <li> <p>Description: A conflicting
// conditional operation is currently in progress against this resource. Please try
// again.</p> </li> </ul> </li> <li> <p>Error code: <code>InternalError</code> </p>
// <ul> <li> <p>Description: The service was unable to apply the provided tag to
// the bucket.</p> </li> </ul> </li> </ul> <p>The following operations are related
// to <code>PutBucketTagging</code>:</p> <ul> <li> <p> <a>GetBucketTagging</a> </p>
// </li> <li> <p> <a>DeleteBucketTagging</a> </p> </li> </ul>
func (c *Client) PutBucketTagging(ctx context.Context, params *PutBucketTaggingInput, optFns ...func(*Options)) (*PutBucketTaggingOutput, error) {
	stack := middleware.NewStack("PutBucketTagging", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpPutBucketTaggingMiddlewares(stack)
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
	addOpPutBucketTaggingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketTagging(options.Region), middleware.Before)
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
			OperationName: "PutBucketTagging",
			Err:           err,
		}
	}
	out := result.(*PutBucketTaggingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketTaggingInput struct {
	// The base64-encoded 128-bit MD5 digest of the data. You must use this header as a
	// message integrity check to verify that the request body was not corrupted in
	// transit. For more information, see RFC 1864
	// (http://www.ietf.org/rfc/rfc1864.txt).
	ContentMD5 *string
	// The bucket name.
	Bucket *string
	// Container for the TagSet and Tag elements.
	Tagging *types.Tagging
}

type PutBucketTaggingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpPutBucketTaggingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpPutBucketTagging{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketTagging{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutBucketTagging(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketTagging",
	}
}
