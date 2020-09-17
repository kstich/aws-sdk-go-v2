// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Revokes an ingress rule in an Amazon Redshift security group for a previously
// authorized IP range or Amazon EC2 security group. To add an ingress rule, see
// AuthorizeClusterSecurityGroupIngress (). For information about managing security
// groups, go to Amazon Redshift Cluster Security Groups
// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-security-groups.html)
// in the Amazon Redshift Cluster Management Guide.
func (c *Client) RevokeClusterSecurityGroupIngress(ctx context.Context, params *RevokeClusterSecurityGroupIngressInput, optFns ...func(*Options)) (*RevokeClusterSecurityGroupIngressOutput, error) {
	stack := middleware.NewStack("RevokeClusterSecurityGroupIngress", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpRevokeClusterSecurityGroupIngressMiddlewares(stack)
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
	addOpRevokeClusterSecurityGroupIngressValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRevokeClusterSecurityGroupIngress(options.Region), middleware.Before)
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
			OperationName: "RevokeClusterSecurityGroupIngress",
			Err:           err,
		}
	}
	out := result.(*RevokeClusterSecurityGroupIngressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type RevokeClusterSecurityGroupIngressInput struct {
	// The AWS account number of the owner of the security group specified in the
	// EC2SecurityGroupName parameter. The AWS access key ID is not an acceptable
	// value. If EC2SecurityGroupOwnerId is specified, EC2SecurityGroupName must also
	// be provided. and CIDRIP cannot be provided. Example: 111122223333
	EC2SecurityGroupOwnerId *string
	// The name of the security Group from which to revoke the ingress rule.
	ClusterSecurityGroupName *string
	// The IP range for which to revoke access. This range must be a valid Classless
	// Inter-Domain Routing (CIDR) block of IP addresses. If CIDRIP is specified,
	// EC2SecurityGroupName and EC2SecurityGroupOwnerId cannot be provided.
	CIDRIP *string
	// The name of the EC2 Security Group whose access is to be revoked. If
	// EC2SecurityGroupName is specified, EC2SecurityGroupOwnerId must also be provided
	// and CIDRIP cannot be provided.
	EC2SecurityGroupName *string
}

type RevokeClusterSecurityGroupIngressOutput struct {
	// Describes a security group.
	ClusterSecurityGroup *types.ClusterSecurityGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpRevokeClusterSecurityGroupIngressMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpRevokeClusterSecurityGroupIngress{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpRevokeClusterSecurityGroupIngress{}, middleware.After)
}

func newServiceMetadataMiddleware_opRevokeClusterSecurityGroupIngress(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "RevokeClusterSecurityGroupIngress",
	}
}
