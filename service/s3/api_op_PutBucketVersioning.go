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

// Sets the versioning state of an existing bucket. To set the versioning state,
// you must be the bucket owner. You can set the versioning state with one of the
// following values:  <p> <b>Enabled</b>—Enables versioning for the objects in the
// bucket. All objects added to the bucket receive a unique version ID.</p> <p>
// <b>Suspended</b>—Disables versioning for the objects in the bucket. All objects
// added to the bucket receive the version ID null.</p> <p>If the versioning state
// has never been set on a bucket, it has no versioning state; a
// <a>GetBucketVersioning</a> request does not return a versioning state value.</p>
// <p>If the bucket owner enables MFA Delete in the bucket versioning
// configuration, the bucket owner must include the <code>x-amz-mfa request</code>
// header and the <code>Status</code> and the <code>MfaDelete</code> request
// elements in a request to set the versioning state of the bucket.</p> <important>
// <p>If you have an object expiration lifecycle policy in your non-versioned
// bucket and you want to maintain the same permanent delete behavior when you
// enable versioning, you must add a noncurrent expiration policy. The noncurrent
// expiration lifecycle policy will manage the deletes of the noncurrent object
// versions in the version-enabled bucket. (A version-enabled bucket maintains one
// current and zero or more noncurrent object versions.) For more information, see
// <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html#lifecycle-and-other-bucket-config">Lifecycle
// and Versioning</a>.</p> </important> <p class="title"> <b>Related Resources</b>
// </p> <ul> <li> <p> <a>CreateBucket</a> </p> </li> <li> <p> <a>DeleteBucket</a>
// </p> </li> <li> <p> <a>GetBucketVersioning</a> </p> </li> </ul>
func (c *Client) PutBucketVersioning(ctx context.Context, params *PutBucketVersioningInput, optFns ...func(*Options)) (*PutBucketVersioningOutput, error) {
	stack := middleware.NewStack("PutBucketVersioning", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpPutBucketVersioningMiddlewares(stack)
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
	addOpPutBucketVersioningValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketVersioning(options.Region), middleware.Before)
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
			OperationName: "PutBucketVersioning",
			Err:           err,
		}
	}
	out := result.(*PutBucketVersioningOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketVersioningInput struct {
	// The concatenation of the authentication device's serial number, a space, and the
	// value that is displayed on your authentication device.
	MFA *string
	// The bucket name.
	Bucket *string
	// >The base64-encoded 128-bit MD5 digest of the data. You must use this header as
	// a message integrity check to verify that the request body was not corrupted in
	// transit. For more information, see RFC 1864
	// (http://www.ietf.org/rfc/rfc1864.txt).
	ContentMD5 *string
	// Container for setting the versioning state.
	VersioningConfiguration *types.VersioningConfiguration
}

type PutBucketVersioningOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpPutBucketVersioningMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpPutBucketVersioning{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketVersioning{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutBucketVersioning(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketVersioning",
	}
}
