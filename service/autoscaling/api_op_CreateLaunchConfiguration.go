// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a launch configuration. If you exceed your maximum limit of launch
// configurations, the call fails. To query this limit, call the
// DescribeAccountLimits () API. For information about updating this limit, see
// Amazon EC2 Auto Scaling Service Quotas
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-account-limits.html)
// in the Amazon EC2 Auto Scaling User Guide. For more information, see Launch
// Configurations
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/LaunchConfiguration.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) CreateLaunchConfiguration(ctx context.Context, params *CreateLaunchConfigurationInput, optFns ...func(*Options)) (*CreateLaunchConfigurationOutput, error) {
	stack := middleware.NewStack("CreateLaunchConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateLaunchConfigurationMiddlewares(stack)
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
	addOpCreateLaunchConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLaunchConfiguration(options.Region), middleware.Before)
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
			OperationName: "CreateLaunchConfiguration",
			Err:           err,
		}
	}
	out := result.(*CreateLaunchConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLaunchConfigurationInput struct {
	// The name or the Amazon Resource Name (ARN) of the instance profile associated
	// with the IAM role for the instance. The instance profile contains the IAM role.
	// For more information, see IAM Role for Applications That Run on Amazon EC2
	// Instances
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/us-iam-role.html) in the
	// Amazon EC2 Auto Scaling User Guide.
	IamInstanceProfile *string
	// The ID of the RAM disk to select.
	RamdiskId *string
	// A block device mapping, which specifies the block devices for the instance. You
	// can specify virtual devices and EBS volumes. For more information, see Block
	// Device Mapping
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html)
	// in the Amazon EC2 User Guide for Linux Instances.
	BlockDeviceMappings []*types.BlockDeviceMapping
	// The ID of the kernel associated with the AMI.
	KernelId *string
	// The metadata options for the instances. For more information, see Instance
	// Metadata and User Data
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html)
	// in the Amazon EC2 User Guide for Linux Instances.
	MetadataOptions *types.InstanceMetadataOptions
	// Specifies whether the launch configuration is optimized for EBS I/O (true) or
	// not (false). The optimization provides dedicated throughput to Amazon EBS and an
	// optimized configuration stack to provide optimal I/O performance. This
	// optimization is not available with all instance types. Additional fees are
	// incurred when you enable EBS optimization for an instance type that is not
	// EBS-optimized by default. For more information, see Amazon EBS-Optimized
	// Instances
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSOptimized.html) in the
	// Amazon EC2 User Guide for Linux Instances. The default value is false.
	EbsOptimized *bool
	// The ID of a ClassicLink-enabled VPC to link your EC2-Classic instances to. For
	// more information, see ClassicLink
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) in
	// the Amazon EC2 User Guide for Linux Instances and Linking EC2-Classic Instances
	// to a VPC
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-in-vpc.html#as-ClassicLink)
	// in the Amazon EC2 Auto Scaling User Guide. This parameter can only be used if
	// you are launching EC2-Classic instances.
	ClassicLinkVPCId *string
	// The ID of the instance to use to create the launch configuration. The new launch
	// configuration derives attributes from the instance, except for the block device
	// mapping. To create a launch configuration with a block device mapping or
	// override any other instance attributes, specify them as part of the same
	// request. For more information, see Create a Launch Configuration Using an EC2
	// Instance
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-lc-with-instanceID.html)
	// in the Amazon EC2 Auto Scaling User Guide. If you do not specify InstanceId, you
	// must specify both ImageId and InstanceType.
	InstanceId *string
	// For Auto Scaling groups that are running in a virtual private cloud (VPC),
	// specifies whether to assign a public IP address to the group's instances. If you
	// specify true, each instance in the Auto Scaling group receives a unique public
	// IP address. For more information, see Launching Auto Scaling Instances in a VPC
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-in-vpc.html) in the
	// Amazon EC2 Auto Scaling User Guide. If you specify this parameter, you must
	// specify at least one subnet for VPCZoneIdentifier when you create your group. If
	// the instance is launched into a default subnet, the default is to assign a
	// public IP address, unless you disabled the option to assign a public IP address
	// on the subnet. If the instance is launched into a nondefault subnet, the default
	// is not to assign a public IP address, unless you enabled the option to assign a
	// public IP address on the subnet.
	AssociatePublicIpAddress *bool
	// The IDs of one or more security groups for the specified ClassicLink-enabled
	// VPC. For more information, see ClassicLink
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) in
	// the Amazon EC2 User Guide for Linux Instances and Linking EC2-Classic Instances
	// to a VPC
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-in-vpc.html#as-ClassicLink)
	// in the Amazon EC2 Auto Scaling User Guide. If you specify the ClassicLinkVPCId
	// parameter, you must specify this parameter.
	ClassicLinkVPCSecurityGroups []*string
	// The tenancy of the instance. An instance with dedicated tenancy runs on
	// isolated, single-tenant hardware and can only be launched into a VPC. To launch
	// dedicated instances into a shared tenancy VPC (a VPC with the instance placement
	// tenancy attribute set to default), you must set the value of this parameter to
	// dedicated. If you specify PlacementTenancy, you must specify at least one subnet
	// for VPCZoneIdentifier when you create your group. For more information, see
	// Instance Placement Tenancy
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-in-vpc.html#as-vpc-tenancy)
	// in the Amazon EC2 Auto Scaling User Guide. Valid Values: default | dedicated
	PlacementTenancy *string
	// A list that contains the security groups to assign to the instances in the Auto
	// Scaling group. [EC2-VPC] Specify the security group IDs. For more information,
	// see Security Groups for Your VPC
	// (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_SecurityGroups.html)
	// in the Amazon Virtual Private Cloud User Guide. [EC2-Classic] Specify either the
	// security group names or the security group IDs. For more information, see Amazon
	// EC2 Security Groups
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-network-security.html)
	// in the Amazon EC2 User Guide for Linux Instances.
	SecurityGroups []*string
	// The name of the launch configuration. This name must be unique per Region per
	// account.
	LaunchConfigurationName *string
	// The maximum hourly price to be paid for any Spot Instance launched to fulfill
	// the request. Spot Instances are launched when the price you specify exceeds the
	// current Spot price. For more information, see Launching Spot Instances in Your
	// Auto Scaling Group
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-launch-spot-instances.html)
	// in the Amazon EC2 Auto Scaling User Guide. When you change your maximum price by
	// creating a new launch configuration, running instances will continue to run as
	// long as the maximum price for those running instances is higher than the current
	// Spot price.
	SpotPrice *string
	// The Base64-encoded user data to make available to the launched EC2 instances.
	// For more information, see Instance Metadata and User Data
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html)
	// in the Amazon EC2 User Guide for Linux Instances.
	UserData *string
	// The name of the key pair. For more information, see Amazon EC2 Key Pairs
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html) in the
	// Amazon EC2 User Guide for Linux Instances.
	KeyName *string
	// The ID of the Amazon Machine Image (AMI) that was assigned during registration.
	// For more information, see Finding an AMI
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/finding-an-ami.html) in the
	// Amazon EC2 User Guide for Linux Instances. If you do not specify InstanceId, you
	// must specify ImageId.
	ImageId *string
	// Controls whether instances in this group are launched with detailed (true) or
	// basic (false) monitoring. The default value is true (enabled). When detailed
	// monitoring is enabled, Amazon CloudWatch generates metrics every minute and your
	// account is charged a fee. When you disable detailed monitoring, CloudWatch
	// generates metrics every 5 minutes. For more information, see Configure
	// Monitoring for Auto Scaling Instances
	// (https://docs.aws.amazon.com/autoscaling/latest/userguide/as-instance-monitoring.html#enable-as-instance-metrics)
	// in the Amazon EC2 Auto Scaling User Guide.
	InstanceMonitoring *types.InstanceMonitoring
	// Specifies the instance type of the EC2 instance. For information about available
	// instance types, see Available Instance Types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#AvailableInstanceTypes)
	// in the Amazon EC2 User Guide for Linux Instances. If you do not specify
	// InstanceId, you must specify InstanceType.
	InstanceType *string
}

type CreateLaunchConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateLaunchConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateLaunchConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateLaunchConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateLaunchConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "CreateLaunchConfiguration",
	}
}
