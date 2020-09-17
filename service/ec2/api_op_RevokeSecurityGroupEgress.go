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

// [VPC only] Removes the specified egress rules from a security group for EC2-VPC.
// This action doesn't apply to security groups for use in EC2-Classic. To remove a
// rule, the values that you specify (for example, ports) must match the existing
// rule's values exactly. Each rule consists of the protocol and the IPv4 or IPv6
// CIDR range or source security group. For the TCP and UDP protocols, you must
// also specify the destination port or range of ports. For the ICMP protocol, you
// must also specify the ICMP type and code. If the security group rule has a
// description, you do not have to specify the description to revoke the rule. Rule
// changes are propagated to instances within the security group as quickly as
// possible. However, a small delay might occur.
func (c *Client) RevokeSecurityGroupEgress(ctx context.Context, params *RevokeSecurityGroupEgressInput, optFns ...func(*Options)) (*RevokeSecurityGroupEgressOutput, error) {
	stack := middleware.NewStack("RevokeSecurityGroupEgress", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpRevokeSecurityGroupEgressMiddlewares(stack)
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
	addOpRevokeSecurityGroupEgressValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRevokeSecurityGroupEgress(options.Region), middleware.Before)
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
			OperationName: "RevokeSecurityGroupEgress",
			Err:           err,
		}
	}
	out := result.(*RevokeSecurityGroupEgressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RevokeSecurityGroupEgressInput struct {
	// Not supported. Use a set of IP permissions to specify the port.
	FromPort *int32
	// The ID of the security group.
	GroupId *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The sets of IP permissions. You can't specify a destination security group and a
	// CIDR IP address range in the same set of permissions.
	IpPermissions []*types.IpPermission
	// Not supported. Use a set of IP permissions to specify a destination security
	// group.
	SourceSecurityGroupName *string
	// Not supported. Use a set of IP permissions to specify the CIDR.
	CidrIp *string
	// Not supported. Use a set of IP permissions to specify a destination security
	// group.
	SourceSecurityGroupOwnerId *string
	// Not supported. Use a set of IP permissions to specify the port.
	ToPort *int32
	// Not supported. Use a set of IP permissions to specify the protocol name or
	// number.
	IpProtocol *string
}

type RevokeSecurityGroupEgressOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpRevokeSecurityGroupEgressMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpRevokeSecurityGroupEgress{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpRevokeSecurityGroupEgress{}, middleware.After)
}

func newServiceMetadataMiddleware_opRevokeSecurityGroupEgress(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "RevokeSecurityGroupEgress",
	}
}
