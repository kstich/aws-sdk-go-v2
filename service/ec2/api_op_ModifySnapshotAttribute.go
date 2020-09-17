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

// Adds or removes permission settings for the specified snapshot. You may add or
// remove specified AWS account IDs from a snapshot's list of create volume
// permissions, but you cannot do both in a single operation. If you need to both
// add and remove account IDs for a snapshot, you must use multiple operations. You
// can make up to 500 modifications to a snapshot in a single operation. Encrypted
// snapshots and snapshots with AWS Marketplace product codes cannot be made
// public. Snapshots encrypted with your default CMK cannot be shared with other
// accounts. For more information about modifying snapshot permissions, see Sharing
// Snapshots
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-modifying-snapshot-permissions.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) ModifySnapshotAttribute(ctx context.Context, params *ModifySnapshotAttributeInput, optFns ...func(*Options)) (*ModifySnapshotAttributeOutput, error) {
	stack := middleware.NewStack("ModifySnapshotAttribute", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpModifySnapshotAttributeMiddlewares(stack)
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
	addOpModifySnapshotAttributeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifySnapshotAttribute(options.Region), middleware.Before)
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
			OperationName: "ModifySnapshotAttribute",
			Err:           err,
		}
	}
	out := result.(*ModifySnapshotAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifySnapshotAttributeInput struct {
	// The account ID to modify for the snapshot.
	UserIds []*string
	// The ID of the snapshot.
	SnapshotId *string
	// The type of operation to perform to the attribute.
	OperationType types.OperationType
	// The snapshot attribute to modify. Only volume creation permissions can be
	// modified.
	Attribute types.SnapshotAttributeName
	// A JSON representation of the snapshot attribute modification.
	CreateVolumePermission *types.CreateVolumePermissionModifications
	// The group to modify for the snapshot.
	GroupNames []*string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type ModifySnapshotAttributeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpModifySnapshotAttributeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpModifySnapshotAttribute{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpModifySnapshotAttribute{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifySnapshotAttribute(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifySnapshotAttribute",
	}
}
