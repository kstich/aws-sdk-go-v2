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

// Describes the most recent volume modification request for the specified EBS
// volumes. If a volume has never been modified, some information in the output
// will be null. If a volume has been modified more than once, the output includes
// only the most recent modification request. You can also use CloudWatch Events to
// check the status of a modification to an EBS volume. For information about
// CloudWatch Events, see the Amazon CloudWatch Events User Guide
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/). For more
// information, see Monitoring Volume Modifications
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-expand-volume.html#monitoring_mods)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) DescribeVolumesModifications(ctx context.Context, params *DescribeVolumesModificationsInput, optFns ...func(*Options)) (*DescribeVolumesModificationsOutput, error) {
	stack := middleware.NewStack("DescribeVolumesModifications", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeVolumesModificationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVolumesModifications(options.Region), middleware.Before)
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
			OperationName: "DescribeVolumesModifications",
			Err:           err,
		}
	}
	out := result.(*DescribeVolumesModificationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVolumesModificationsInput struct {
	// The filters.
	//
	//     * modification-state - The current modification state
	// (modifying | optimizing | completed | failed).
	//
	//     * original-iops - The
	// original IOPS rate of the volume.
	//
	//     * original-size - The original size of
	// the volume, in GiB.
	//
	//     * original-volume-type - The original volume type of
	// the volume (standard | io1 | gp2 | sc1 | st1).
	//
	//     * originalMultiAttachEnabled
	// - Indicates whether Multi-Attach support was enabled (true | false).
	//
	//     *
	// start-time - The modification start time.
	//
	//     * target-iops - The target IOPS
	// rate of the volume.
	//
	//     * target-size - The target size of the volume, in
	// GiB.
	//
	//     * target-volume-type - The target volume type of the volume (standard
	// | io1 | gp2 | sc1 | st1).
	//
	//     * targetMultiAttachEnabled - Indicates whether
	// Multi-Attach support is to be enabled (true | false).
	//
	//     * volume-id - The ID
	// of the volume.
	Filters []*types.Filter
	// The nextToken value returned by a previous paginated request.
	NextToken *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The maximum number of results (up to a limit of 500) to be returned in a
	// paginated request.
	MaxResults *int32
	// The IDs of the volumes.
	VolumeIds []*string
}

type DescribeVolumesModificationsOutput struct {
	// Information about the volume modifications.
	VolumesModifications []*types.VolumeModification
	// Token for pagination, null if there are no more results
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeVolumesModificationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeVolumesModifications{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVolumesModifications{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeVolumesModifications(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVolumesModifications",
	}
}
