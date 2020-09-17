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

// Describes one or more of your network ACLs. For more information, see Network
// ACLs (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_ACLs.html) in the
// Amazon Virtual Private Cloud User Guide.
func (c *Client) DescribeNetworkAcls(ctx context.Context, params *DescribeNetworkAclsInput, optFns ...func(*Options)) (*DescribeNetworkAclsOutput, error) {
	stack := middleware.NewStack("DescribeNetworkAcls", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeNetworkAclsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNetworkAcls(options.Region), middleware.Before)
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
			OperationName: "DescribeNetworkAcls",
			Err:           err,
		}
	}
	out := result.(*DescribeNetworkAclsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeNetworkAclsInput struct {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// One or more network ACL IDs. Default: Describes all your network ACLs.
	NetworkAclIds []*string
	// One or more filters.
	//
	//     * association.association-id - The ID of an
	// association ID for the ACL.
	//
	//     * association.network-acl-id - The ID of the
	// network ACL involved in the association.
	//
	//     * association.subnet-id - The ID
	// of the subnet involved in the association.
	//
	//     * default - Indicates whether
	// the ACL is the default network ACL for the VPC.
	//
	//     * entry.cidr - The IPv4
	// CIDR range specified in the entry.
	//
	//     * entry.icmp.code - The ICMP code
	// specified in the entry, if any.
	//
	//     * entry.icmp.type - The ICMP type specified
	// in the entry, if any.
	//
	//     * entry.ipv6-cidr - The IPv6 CIDR range specified in
	// the entry.
	//
	//     * entry.port-range.from - The start of the port range specified
	// in the entry.
	//
	//     * entry.port-range.to - The end of the port range specified
	// in the entry.
	//
	//     * entry.protocol - The protocol specified in the entry (tcp |
	// udp | icmp or a protocol number).
	//
	//     * entry.rule-action - Allows or denies
	// the matching traffic (allow | deny).
	//
	//     * entry.rule-number - The number of an
	// entry (in other words, rule) in the set of ACL entries.
	//
	//     * network-acl-id -
	// The ID of the network ACL.
	//
	//     * owner-id - The ID of the AWS account that owns
	// the network ACL.
	//
	//     * tag: - The key/value combination of a tag assigned to
	// the resource. Use the tag key in the filter name and the tag value as the filter
	// value. For example, to find all resources that have a tag with the key Owner and
	// the value TeamA, specify tag:Owner for the filter name and TeamA for the filter
	// value.
	//
	//     * tag-key - The key of a tag assigned to the resource. Use this
	// filter to find all resources assigned a tag with a specific key, regardless of
	// the tag value.
	//
	//     * vpc-id - The ID of the VPC for the network ACL.
	Filters []*types.Filter
	// The token for the next page of results.
	NextToken *string
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32
}

type DescribeNetworkAclsOutput struct {
	// Information about one or more network ACLs.
	NetworkAcls []*types.NetworkAcl
	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeNetworkAclsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeNetworkAcls{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeNetworkAcls{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeNetworkAcls(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeNetworkAcls",
	}
}
