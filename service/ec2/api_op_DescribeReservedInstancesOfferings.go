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

// Describes Reserved Instance offerings that are available for purchase. With
// Reserved Instances, you purchase the right to launch instances for a period of
// time. During that time period, you do not receive insufficient capacity errors,
// and you pay a lower usage rate than the rate charged for On-Demand instances for
// the actual time used. If you have listed your own Reserved Instances for sale in
// the Reserved Instance Marketplace, they will be excluded from these results.
// This is to ensure that you do not purchase your own Reserved Instances. For more
// information, see Reserved Instance Marketplace
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ri-market-general.html) in
// the Amazon Elastic Compute Cloud User Guide.
func (c *Client) DescribeReservedInstancesOfferings(ctx context.Context, params *DescribeReservedInstancesOfferingsInput, optFns ...func(*Options)) (*DescribeReservedInstancesOfferingsOutput, error) {
	stack := middleware.NewStack("DescribeReservedInstancesOfferings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeReservedInstancesOfferingsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReservedInstancesOfferings(options.Region), middleware.Before)
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
			OperationName: "DescribeReservedInstancesOfferings",
			Err:           err,
		}
	}
	out := result.(*DescribeReservedInstancesOfferingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeReservedInstancesOfferings.
type DescribeReservedInstancesOfferingsInput struct {
	// The offering class of the Reserved Instance. Can be standard or convertible.
	OfferingClass types.OfferingClassType
	// The instance type that the reservation will cover (for example, m1.small). For
	// more information, see Instance Types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	InstanceType types.InstanceType
	// The token to retrieve the next page of results.
	NextToken *string
	// The Reserved Instance product platform description. Instances that include
	// (Amazon VPC) in the description are for use with Amazon VPC.
	ProductDescription types.RIProductDescription
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The Availability Zone in which the Reserved Instance can be used.
	AvailabilityZone *string
	// The tenancy of the instances covered by the reservation. A Reserved Instance
	// with a tenancy of dedicated is applied to instances that run in a VPC on
	// single-tenant hardware (i.e., Dedicated Instances). Important: The host value
	// cannot be used with this parameter. Use the default or dedicated values only.
	// Default: default
	InstanceTenancy types.Tenancy
	// The Reserved Instance offering type. If you are using tools that predate the
	// 2011-11-01 API version, you only have access to the Medium Utilization Reserved
	// Instance offering type.
	OfferingType types.OfferingTypeValues
	// The minimum duration (in seconds) to filter when searching for offerings.
	// Default: 2592000 (1 month)
	MinDuration *int64
	// Include Reserved Instance Marketplace offerings in the response.
	IncludeMarketplace *bool
	// One or more filters.
	//
	//     * availability-zone - The Availability Zone where the
	// Reserved Instance can be used.
	//
	//     * duration - The duration of the Reserved
	// Instance (for example, one year or three years), in seconds (31536000 |
	// 94608000).
	//
	//     * fixed-price - The purchase price of the Reserved Instance (for
	// example, 9800.0).
	//
	//     * instance-type - The instance type that is covered by
	// the reservation.
	//
	//     * marketplace - Set to true to show only Reserved Instance
	// Marketplace offerings. When this filter is not used, which is the default
	// behavior, all offerings from both AWS and the Reserved Instance Marketplace are
	// listed.
	//
	//     * product-description - The Reserved Instance product platform
	// description. Instances that include (Amazon VPC) in the product platform
	// description will only be displayed to EC2-Classic account holders and are for
	// use with Amazon VPC. (Linux/UNIX | Linux/UNIX (Amazon VPC) | SUSE Linux | SUSE
	// Linux (Amazon VPC) | Red Hat Enterprise Linux | Red Hat Enterprise Linux (Amazon
	// VPC) | Windows | Windows (Amazon VPC) | Windows with SQL Server Standard |
	// Windows with SQL Server Standard (Amazon VPC) | Windows with SQL Server Web |
	// Windows with SQL Server Web (Amazon VPC) | Windows with SQL Server Enterprise |
	// Windows with SQL Server Enterprise (Amazon VPC))
	//
	//     *
	// reserved-instances-offering-id - The Reserved Instances offering ID.
	//
	//     *
	// scope - The scope of the Reserved Instance (Availability Zone or Region).
	//
	//     *
	// usage-price - The usage price of the Reserved Instance, per hour (for example,
	// 0.84).
	Filters []*types.Filter
	// The maximum duration (in seconds) to filter when searching for offerings.
	// Default: 94608000 (3 years)
	MaxDuration *int64
	// The maximum number of results to return for the request in a single page. The
	// remaining results of the initial request can be seen by sending another request
	// with the returned NextToken value. The maximum is 100. Default: 100
	MaxResults *int32
	// The maximum number of instances to filter when searching for offerings. Default:
	// 20
	MaxInstanceCount *int32
	// One or more Reserved Instances offering IDs.
	ReservedInstancesOfferingIds []*string
}

// Contains the output of DescribeReservedInstancesOfferings.
type DescribeReservedInstancesOfferingsOutput struct {
	// A list of Reserved Instances offerings.
	ReservedInstancesOfferings []*types.ReservedInstancesOffering
	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeReservedInstancesOfferingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeReservedInstancesOfferings{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeReservedInstancesOfferings{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeReservedInstancesOfferings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeReservedInstancesOfferings",
	}
}
