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

// Describes the details of the instance types that are offered in a location. The
// results can be filtered by the attributes of the instance types.
func (c *Client) DescribeInstanceTypes(ctx context.Context, params *DescribeInstanceTypesInput, optFns ...func(*Options)) (*DescribeInstanceTypesOutput, error) {
	stack := middleware.NewStack("DescribeInstanceTypes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeInstanceTypesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstanceTypes(options.Region), middleware.Before)
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
			OperationName: "DescribeInstanceTypes",
			Err:           err,
		}
	}
	out := result.(*DescribeInstanceTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstanceTypesInput struct {
	// The maximum number of results to return for the request in a single page. The
	// remaining results can be seen by sending another request with the next token
	// value.
	MaxResults *int32
	// The token to retrieve the next page of results.
	NextToken *string
	// The instance types. For more information, see Instance Types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	InstanceTypes []types.InstanceType
	// One or more filters. Filter names and values are case-sensitive.
	//
	//     *
	// auto-recovery-supported - Indicates whether auto recovery is supported. (true |
	// false)
	//
	//     * bare-metal - Indicates whether it is a bare metal instance type.
	// (true | false)
	//
	//     * burstable-performance-supported - Indicates whether it is
	// a burstable performance instance type. (true | false)
	//
	//     * current-generation
	// - Indicates whether this instance type is the latest generation instance type of
	// an instance family. (true | false)
	//
	//     *
	// ebs-info.ebs-optimized-info.baseline-bandwidth-in-mbps - The baseline bandwidth
	// performance for an EBS-optimized instance type, in Mbps.
	//
	//     *
	// ebs-info.ebs-optimized-info.baseline-throughput-in-mbps - The baseline
	// throughput performance for an EBS-optimized instance type, in MBps.
	//
	//     *
	// ebs-info.ebs-optimized-info.baseline-iops - The baseline input/output storage
	// operations per second for an EBS-optimized instance type.
	//
	//     *
	// ebs-info.ebs-optimized-info.maximum-bandwidth-in-mbps - The maximum bandwidth
	// performance for an EBS-optimized instance type, in Mbps.
	//
	//     *
	// ebs-info.ebs-optimized-info.maximum-throughput-in-mbps - The maximum throughput
	// performance for an EBS-optimized instance type, in MBps.
	//
	//     *
	// ebs-info.ebs-optimized-info.maximum-iops - The maximum input/output storage
	// operations per second for an EBS-optimized instance type.
	//
	//     *
	// ebs-info.ebs-optimized-support - Indicates whether the instance type is
	// EBS-optimized. (supported | unsupported | default)
	//
	//     *
	// ebs-info.encryption-support - Indicates whether EBS encryption is supported.
	// (supported | unsupported)
	//
	//     * ebs-info.nvme-support - Indicates whether
	// non-volatile memory express (NVMe) is supported or required. (required |
	// supported | unsupported)
	//
	//     * free-tier-eligible - Indicates whether the
	// instance type is eligible to use in the free tier. (true | false)
	//
	//     *
	// hibernation-supported - Indicates whether On-Demand hibernation is supported.
	// (true | false)
	//
	//     * hypervisor - The hypervisor used. (nitro | xen)
	//
	//     *
	// instance-storage-info.disk.count - The number of local disks.
	//
	//     *
	// instance-storage-info.disk.size-in-gb - The storage size of each instance
	// storage disk, in GB.
	//
	//     * instance-storage-info.disk.type - The storage
	// technology for the local instance storage disks. (hdd | ssd)
	//
	//     *
	// instance-storage-info.total-size-in-gb - The total amount of storage available
	// from all local instance storage, in GB.
	//
	//     * instance-storage-supported -
	// Indicates whether the instance type has local instance storage. (true | false)
	//
	//
	// * memory-info.size-in-mib - The memory size.
	//
	//     * network-info.ena-support -
	// Indicates whether Elastic Network Adapter (ENA) is supported or required.
	// (required | supported | unsupported)
	//
	//     * network-info.efa-supported -
	// Indicates whether the instance type supports Elastic Fabric Adapter (EFA). (true
	// | false)
	//
	//     * network-info.ipv4-addresses-per-interface - The maximum number
	// of private IPv4 addresses per network interface.
	//
	//     *
	// network-info.ipv6-addresses-per-interface - The maximum number of private IPv6
	// addresses per network interface.
	//
	//     * network-info.ipv6-supported - Indicates
	// whether the instance type supports IPv6. (true | false)
	//
	//     *
	// network-info.maximum-network-interfaces - The maximum number of network
	// interfaces per instance.
	//
	//     * network-info.network-performance - Describes the
	// network performance.
	//
	//     * processor-info.sustained-clock-speed-in-ghz - The
	// CPU clock speed, in GHz.
	//
	//     * vcpu-info.default-cores - The default number of
	// cores for the instance type.
	//
	//     * vcpu-info.default-threads-per-core - The
	// default number of threads per core for the instance type.
	//
	//     *
	// vcpu-info.default-vcpus - The default number of vCPUs for the instance type.
	Filters []*types.Filter
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DescribeInstanceTypesOutput struct {
	// The instance type. For more information, see Instance Types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	InstanceTypes []*types.InstanceTypeInfo
	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeInstanceTypesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeInstanceTypes{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeInstanceTypes{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeInstanceTypes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeInstanceTypes",
	}
}
