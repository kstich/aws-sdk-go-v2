// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Copies a point-in-time snapshot of an EBS volume and stores it in Amazon S3. You
// can copy the snapshot within the same Region or from one Region to another. You
// can use the snapshot to create EBS volumes or Amazon Machine Images (AMIs).
// Copies of encrypted EBS snapshots remain encrypted. Copies of unencrypted
// snapshots remain unencrypted, unless you enable encryption for the snapshot copy
// operation. By default, encrypted snapshot copies use the default AWS Key
// Management Service (AWS KMS) customer master key (CMK); however, you can specify
// a different CMK. To copy an encrypted snapshot that has been shared from another
// account, you must have permissions for the CMK used to encrypt the snapshot.
// Snapshots created by copying another snapshot have an arbitrary volume ID that
// should not be used for any purpose. For more information, see Copying an Amazon
// EBS Snapshot
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-copy-snapshot.html) in
// the Amazon Elastic Compute Cloud User Guide.
func (c *Client) CopySnapshot(ctx context.Context, params *CopySnapshotInput, optFns ...func(*Options)) (*CopySnapshotOutput, error) {
	stack := middleware.NewStack("CopySnapshot", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpCopySnapshotMiddlewares(stack)
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
	addOpCopySnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCopySnapshot(options.Region), middleware.Before)
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
			OperationName: "CopySnapshot",
			Err:           err,
		}
	}
	out := result.(*CopySnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopySnapshotInput struct {
	// A description for the EBS snapshot.
	Description *string
	// When you copy an encrypted source snapshot using the Amazon EC2 Query API, you
	// must supply a pre-signed URL. This parameter is optional for unencrypted
	// snapshots. For more information, see Query Requests
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html).
	// The PresignedUrl should use the snapshot source endpoint, the CopySnapshot
	// action, and include the SourceRegion, SourceSnapshotId, and DestinationRegion
	// parameters. The PresignedUrl must be signed using AWS Signature Version 4.
	// Because EBS snapshots are stored in Amazon S3, the signing algorithm for this
	// parameter uses the same logic that is described in Authenticating Requests by
	// Using Query Parameters (AWS Signature Version 4)
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html)
	// in the Amazon Simple Storage Service API Reference. An invalid or improperly
	// signed PresignedUrl will cause the copy operation to fail asynchronously, and
	// the snapshot will move to an error state.
	PresignedUrl *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The ID of the Region that contains the snapshot to be copied.
	SourceRegion *string
	// The identifier of the AWS Key Management Service (AWS KMS) customer master key
	// (CMK) to use for Amazon EBS encryption. If this parameter is not specified, your
	// AWS managed CMK for EBS is used. If KmsKeyId is specified, the encrypted state
	// must be true. You can specify the CMK using any of the following:
	//
	//     * Key ID.
	// For example, key/1234abcd-12ab-34cd-56ef-1234567890ab.
	//
	//     * Key alias. For
	// example, alias/ExampleAlias.
	//
	//     * Key ARN. For example,
	// arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef.
	//
	//
	// * Alias ARN. For example,
	// arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	//
	// AWS authenticates the
	// CMK asynchronously. Therefore, if you specify an ID, alias, or ARN that is not
	// valid, the action can appear to complete, but eventually fails.
	KmsKeyId *string
	// The destination Region to use in the PresignedUrl parameter of a snapshot copy
	// operation. This parameter is only valid for specifying the destination Region in
	// a PresignedUrl parameter, where it is required.  <p>The snapshot copy is sent to
	// the regional endpoint that you sent the HTTP request to (for example,
	// <code>ec2.us-east-1.amazonaws.com</code>). With the AWS CLI, this is specified
	// using the <code>--region</code> parameter or the default Region in your AWS
	// configuration file.</p>
	DestinationRegion *string
	// To encrypt a copy of an unencrypted snapshot if encryption by default is not
	// enabled, enable encryption using this parameter. Otherwise, omit this parameter.
	// Encrypted snapshots are encrypted, even if you omit this parameter and
	// encryption by default is not enabled. You cannot set this parameter to false.
	// For more information, see Amazon EBS Encryption
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	Encrypted *bool
	// The ID of the EBS snapshot to copy.
	SourceSnapshotId *string
	// The tags to apply to the new snapshot.
	TagSpecifications []*types.TagSpecification
}

type CopySnapshotOutput struct {
	// The ID of the new snapshot.
	SnapshotId *string
	// Any tags applied to the new snapshot.
	Tags []*types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpCopySnapshotMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpCopySnapshot{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpCopySnapshot{}, middleware.After)
}

func newServiceMetadataMiddleware_opCopySnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CopySnapshot",
	}
}
