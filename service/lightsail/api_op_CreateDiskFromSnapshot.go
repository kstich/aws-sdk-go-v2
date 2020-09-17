// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a block storage disk from a manual or automatic snapshot of a disk. The
// resulting disk can be attached to an Amazon Lightsail instance in the same
// Availability Zone (e.g., us-east-2a). The create disk from snapshot operation
// supports tag-based access control via request tags and resource tags applied to
// the resource identified by disk snapshot name. For more information, see the
// Lightsail Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) CreateDiskFromSnapshot(ctx context.Context, params *CreateDiskFromSnapshotInput, optFns ...func(*Options)) (*CreateDiskFromSnapshotOutput, error) {
	stack := middleware.NewStack("CreateDiskFromSnapshot", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateDiskFromSnapshotMiddlewares(stack)
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
	addOpCreateDiskFromSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDiskFromSnapshot(options.Region), middleware.Before)
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
			OperationName: "CreateDiskFromSnapshot",
			Err:           err,
		}
	}
	out := result.(*CreateDiskFromSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDiskFromSnapshotInput struct {
	// The name of the source disk from which the source automatic snapshot was
	// created. Constraints:
	//
	//     * This parameter cannot be defined together with the
	// disk snapshot name parameter. The source disk name and disk snapshot name
	// parameters are mutually exclusive.
	//
	//     * Define this parameter only when
	// creating a new disk from an automatic snapshot. For more information, see the
	// Lightsail Dev Guide
	// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-configuring-automatic-snapshots).
	SourceDiskName *string
	// The unique Lightsail disk name (e.g., my-disk).
	DiskName *string
	// The size of the disk in GB (e.g., 32).
	SizeInGb *int32
	// An array of objects that represent the add-ons to enable for the new disk.
	AddOns []*types.AddOnRequest
	// The name of the disk snapshot (e.g., my-snapshot) from which to create the new
	// storage disk. Constraint:
	//
	//     * This parameter cannot be defined together with
	// the source disk name parameter. The disk snapshot name and source disk name
	// parameters are mutually exclusive.
	DiskSnapshotName *string
	// A Boolean value to indicate whether to use the latest available automatic
	// snapshot. Constraints:
	//
	//     * This parameter cannot be defined together with the
	// restore date parameter. The use latest restorable auto snapshot and restore date
	// parameters are mutually exclusive.
	//
	//     * Define this parameter only when
	// creating a new disk from an automatic snapshot. For more information, see the
	// Lightsail Dev Guide
	// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-configuring-automatic-snapshots).
	UseLatestRestorableAutoSnapshot *bool
	// The tag keys and optional values to add to the resource during create. Use the
	// TagResource action to tag a resource after it's created.
	Tags []*types.Tag
	// The Availability Zone where you want to create the disk (e.g., us-east-2a).
	// Choose the same Availability Zone as the Lightsail instance where you want to
	// create the disk. Use the GetRegions operation to list the Availability Zones
	// where Lightsail is currently available.
	AvailabilityZone *string
	// The date of the automatic snapshot to use for the new disk. Use the get auto
	// snapshots operation to identify the dates of the available automatic snapshots.
	// Constraints:
	//
	//     * Must be specified in YYYY-MM-DD format.
	//
	//     * This
	// parameter cannot be defined together with the use latest restorable auto
	// snapshot parameter. The restore date and use latest restorable auto snapshot
	// parameters are mutually exclusive.
	//
	//     * Define this parameter only when
	// creating a new disk from an automatic snapshot. For more information, see the
	// Lightsail Dev Guide
	// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-configuring-automatic-snapshots).
	RestoreDate *string
}

type CreateDiskFromSnapshotOutput struct {
	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateDiskFromSnapshotMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDiskFromSnapshot{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDiskFromSnapshot{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDiskFromSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "CreateDiskFromSnapshot",
	}
}
