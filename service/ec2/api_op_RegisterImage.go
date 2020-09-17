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

// Registers an AMI. When you're creating an AMI, this is the final step you must
// complete before you can launch an instance from the AMI. For more information
// about creating AMIs, see Creating your own AMIs
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/creating-an-ami.html) in
// the Amazon Elastic Compute Cloud User Guide. For Amazon EBS-backed instances,
// CreateImage () creates and registers the AMI in a single request, so you don't
// have to register the AMI yourself.  <p>You can also use
// <code>RegisterImage</code> to create an Amazon EBS-backed Linux AMI from a
// snapshot of a root device volume. You specify the snapshot using the block
// device mapping. For more information, see <a
// href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-launch-snapshot.html">Launching
// a Linux instance from a backup</a> in the <i>Amazon Elastic Compute Cloud User
// Guide</i>.</p> <p>If any snapshots have AWS Marketplace product codes, they are
// copied to the new AMI.</p> <p>Windows and some Linux distributions, such as Red
// Hat Enterprise Linux (RHEL) and SUSE Linux Enterprise Server (SLES), use the EC2
// billing product code associated with an AMI to verify the subscription status
// for package updates. To create a new AMI for operating systems that require a
// billing product code, instead of registering the AMI, do the following to
// preserve the billing product code association:</p> <ol> <li> <p>Launch an
// instance from an existing AMI with that billing product code.</p> </li> <li>
// <p>Customize the instance.</p> </li> <li> <p>Create an AMI from the instance
// using <a>CreateImage</a>.</p> </li> </ol> <p>If you purchase a Reserved Instance
// to apply to an On-Demand Instance that was launched from an AMI with a billing
// product code, make sure that the Reserved Instance has the matching billing
// product code. If you purchase a Reserved Instance without the matching billing
// product code, the Reserved Instance will not be applied to the On-Demand
// Instance. For information about how to obtain the platform details and billing
// information of an AMI, see <a
// href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-billing-info.html">Obtaining
// billing information</a> in the <i>Amazon Elastic Compute Cloud User
// Guide</i>.</p> <p>If needed, you can deregister an AMI at any time. Any
// modifications you make to an AMI backed by an instance store volume invalidates
// its registration. If you make changes to an image, deregister the previous image
// and register the new image.</p>
func (c *Client) RegisterImage(ctx context.Context, params *RegisterImageInput, optFns ...func(*Options)) (*RegisterImageOutput, error) {
	stack := middleware.NewStack("RegisterImage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpRegisterImageMiddlewares(stack)
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
	addOpRegisterImageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterImage(options.Region), middleware.Before)
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
			OperationName: "RegisterImage",
			Err:           err,
		}
	}
	out := result.(*RegisterImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for RegisterImage.
type RegisterImageInput struct {
	// Set to simple to enable enhanced networking with the Intel 82599 Virtual
	// Function interface for the AMI and any instances that you launch from the AMI.
	// There is no way to disable sriovNetSupport at this time. This option is
	// supported only for HVM AMIs. Specifying this option with a PV AMI can make
	// instances launched from the AMI unreachable.
	SriovNetSupport *string
	// The full path to your AMI manifest in Amazon S3 storage. The specified bucket
	// must have the aws-exec-read canned access control list (ACL) to ensure that it
	// can be accessed by Amazon EC2. For more information, see Canned ACLs
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)
	// in the Amazon S3 Service Developer Guide.
	ImageLocation *string
	// The device name of the root device volume (for example, /dev/sda1).
	RootDeviceName *string
	// The billing product codes. Your account must be authorized to specify billing
	// product codes. Otherwise, you can use the AWS Marketplace to bill for the use of
	// an AMI.
	BillingProducts []*string
	// The architecture of the AMI. Default: For Amazon EBS-backed AMIs, i386. For
	// instance store-backed AMIs, the architecture specified in the manifest file.
	Architecture types.ArchitectureValues
	// The type of virtualization (hvm | paravirtual). Default: paravirtual
	VirtualizationType *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The ID of the kernel.
	KernelId *string
	// The ID of the RAM disk.
	RamdiskId *string
	// A name for your AMI. Constraints: 3-128 alphanumeric characters, parentheses
	// (()), square brackets ([]), spaces ( ), periods (.), slashes (/), dashes (-),
	// single quotes ('), at-signs (@), or underscores(_)
	Name *string
	// A description for your AMI.
	Description *string
	// The block device mapping entries.
	BlockDeviceMappings []*types.BlockDeviceMapping
	// Set to true to enable enhanced networking with ENA for the AMI and any instances
	// that you launch from the AMI. This option is supported only for HVM AMIs.
	// Specifying this option with a PV AMI can make instances launched from the AMI
	// unreachable.
	EnaSupport *bool
}

// Contains the output of RegisterImage.
type RegisterImageOutput struct {
	// The ID of the newly registered AMI.
	ImageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpRegisterImageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpRegisterImage{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpRegisterImage{}, middleware.After)
}

func newServiceMetadataMiddleware_opRegisterImage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "RegisterImage",
	}
}
