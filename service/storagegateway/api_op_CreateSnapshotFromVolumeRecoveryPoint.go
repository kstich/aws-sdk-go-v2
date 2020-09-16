// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Initiates a snapshot of a gateway from a volume recovery point. This operation
// is only supported in the cached volume gateway type.  <p>A volume recovery point
// is a point in time at which all data of the volume is consistent and from which
// you can create a snapshot. To get a list of volume recovery point for cached
// volume gateway, use <a>ListVolumeRecoveryPoints</a>.</p> <p>In the
// <code>CreateSnapshotFromVolumeRecoveryPoint</code> request, you identify the
// volume by providing its Amazon Resource Name (ARN). You must also provide a
// description for the snapshot. When the gateway takes a snapshot of the specified
// volume, the snapshot and its description appear in the AWS Storage Gateway
// console. In response, the gateway returns you a snapshot ID. You can use this
// snapshot ID to check the snapshot progress or later use it when you want to
// create a volume from a snapshot.</p> <note> <p>To list or delete a snapshot, you
// must use the Amazon EC2 API. For more information, see <a
// href="https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSnapshots.html">DescribeSnapshots</a>
// or <a
// href="https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DeleteSnapshot.html">DeleteSnapshot</a>
// in the <i>Amazon Elastic Compute Cloud API Reference</i>.</p> </note>
func (c *Client) CreateSnapshotFromVolumeRecoveryPoint(ctx context.Context, params *CreateSnapshotFromVolumeRecoveryPointInput, optFns ...func(*Options)) (*CreateSnapshotFromVolumeRecoveryPointOutput, error) {
	stack := middleware.NewStack("CreateSnapshotFromVolumeRecoveryPoint", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateSnapshotFromVolumeRecoveryPointMiddlewares(stack)
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
	addOpCreateSnapshotFromVolumeRecoveryPointValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSnapshotFromVolumeRecoveryPoint(options.Region), middleware.Before)

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
			OperationName: "CreateSnapshotFromVolumeRecoveryPoint",
			Err:           err,
		}
	}
	out := result.(*CreateSnapshotFromVolumeRecoveryPointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSnapshotFromVolumeRecoveryPointInput struct {
	// A list of up to 50 tags that can be assigned to a snapshot. Each tag is a
	// key-value pair.  <note> <p>Valid characters for key and value are letters,
	// spaces, and numbers representable in UTF-8 format, and the following special
	// characters: + - = . _ : / @. The maximum length of a tag's key is 128
	// characters, and the maximum length for a tag's value is 256.</p> </note>
	Tags []*types.Tag
	// Textual description of the snapshot that appears in the Amazon EC2 console,
	// Elastic Block Store snapshots panel in the Description field, and in the AWS
	// Storage Gateway snapshot Details pane, Description field.
	SnapshotDescription *string
	// The Amazon Resource Name (ARN) of the iSCSI volume target. Use the
	// DescribeStorediSCSIVolumes () operation to return to retrieve the TargetARN for
	// specified VolumeARN.
	VolumeARN *string
}

type CreateSnapshotFromVolumeRecoveryPointOutput struct {
	// The Amazon Resource Name (ARN) of the iSCSI volume target. Use the
	// DescribeStorediSCSIVolumes () operation to return to retrieve the TargetARN for
	// specified VolumeARN.
	VolumeARN *string
	// The ID of the snapshot.
	SnapshotId *string
	// The time the volume was created from the recovery point.
	VolumeRecoveryPointTime *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateSnapshotFromVolumeRecoveryPointMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateSnapshotFromVolumeRecoveryPoint{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateSnapshotFromVolumeRecoveryPoint{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateSnapshotFromVolumeRecoveryPoint(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "CreateSnapshotFromVolumeRecoveryPoint",
	}
}