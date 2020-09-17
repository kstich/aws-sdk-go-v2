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

// Describes the specified images (AMIs, AKIs, and ARIs) available to you or all of
// the images available to you. The images available to you include public images,
// private images that you own, and private images owned by other AWS accounts for
// which you have explicit launch permissions. Recently deregistered images appear
// in the returned results for a short interval and then return empty results.
// After all instances that reference a deregistered AMI are terminated, specifying
// the ID of the image results in an error indicating that the AMI ID cannot be
// found.
func (c *Client) DescribeImages(ctx context.Context, params *DescribeImagesInput, optFns ...func(*Options)) (*DescribeImagesOutput, error) {
	stack := middleware.NewStack("DescribeImages", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeImagesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeImages(options.Region), middleware.Before)
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
			OperationName: "DescribeImages",
			Err:           err,
		}
	}
	out := result.(*DescribeImagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeImagesInput struct {
	// Scopes the results to images with the specified owners. You can specify a
	// combination of AWS account IDs, self, amazon, and aws-marketplace. If you omit
	// this parameter, the results include all images for which you have launch
	// permissions, regardless of ownership.
	Owners []*string
	// The image IDs. Default: Describes all images available to you.
	ImageIds []*string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The filters.
	//
	//     * architecture - The image architecture (i386 | x86_64 |
	// arm64).
	//
	//     * block-device-mapping.delete-on-termination - A Boolean value that
	// indicates whether the Amazon EBS volume is deleted on instance termination.
	//
	//
	// * block-device-mapping.device-name - The device name specified in the block
	// device mapping (for example, /dev/sdh or xvdh).
	//
	//     *
	// block-device-mapping.snapshot-id - The ID of the snapshot used for the EBS
	// volume.
	//
	//     * block-device-mapping.volume-size - The volume size of the EBS
	// volume, in GiB.
	//
	//     * block-device-mapping.volume-type - The volume type of the
	// EBS volume (gp2 | io1 | st1 | sc1 | standard).
	//
	//     *
	// block-device-mapping.encrypted - A Boolean that indicates whether the EBS volume
	// is encrypted.
	//
	//     * description - The description of the image (provided during
	// image creation).
	//
	//     * ena-support - A Boolean that indicates whether enhanced
	// networking with ENA is enabled.
	//
	//     * hypervisor - The hypervisor type (ovm |
	// xen).
	//
	//     * image-id - The ID of the image.
	//
	//     * image-type - The image type
	// (machine | kernel | ramdisk).
	//
	//     * is-public - A Boolean that indicates
	// whether the image is public.
	//
	//     * kernel-id - The kernel ID.
	//
	//     *
	// manifest-location - The location of the image manifest.
	//
	//     * name - The name
	// of the AMI (provided during image creation).
	//
	//     * owner-alias - The owner
	// alias, from an Amazon-maintained list (amazon | aws-marketplace). This is not
	// the user-configured AWS account alias set using the IAM console. We recommend
	// that you use the related parameter instead of this filter.
	//
	//     * owner-id - The
	// AWS account ID of the owner. We recommend that you use the related parameter
	// instead of this filter.
	//
	//     * platform - The platform. To only list
	// Windows-based AMIs, use windows.
	//
	//     * product-code - The product code.
	//
	//     *
	// product-code.type - The type of the product code (devpay | marketplace).
	//
	//     *
	// ramdisk-id - The RAM disk ID.
	//
	//     * root-device-name - The device name of the
	// root device volume (for example, /dev/sda1).
	//
	//     * root-device-type - The type
	// of the root device volume (ebs | instance-store).
	//
	//     * state - The state of
	// the image (available | pending | failed).
	//
	//     * state-reason-code - The reason
	// code for the state change.
	//
	//     * state-reason-message - The message for the
	// state change.
	//
	//     * sriov-net-support - A value of simple indicates that
	// enhanced networking with the Intel 82599 VF interface is enabled.
	//
	//     * tag: -
	// The key/value combination of a tag assigned to the resource. Use the tag key in
	// the filter name and the tag value as the filter value. For example, to find all
	// resources that have a tag with the key Owner and the value TeamA, specify
	// tag:Owner for the filter name and TeamA for the filter value.
	//
	//     * tag-key -
	// The key of a tag assigned to the resource. Use this filter to find all resources
	// assigned a tag with a specific key, regardless of the tag value.
	//
	//     *
	// virtualization-type - The virtualization type (paravirtual | hvm).
	Filters []*types.Filter
	// Scopes the images by users with explicit launch permissions. Specify an AWS
	// account ID, self (the sender of the request), or all (public AMIs).
	ExecutableUsers []*string
}

type DescribeImagesOutput struct {
	// Information about the images.
	Images []*types.Image

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeImagesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeImages{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeImages{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeImages(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeImages",
	}
}
