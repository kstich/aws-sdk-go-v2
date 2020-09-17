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

// Disassociates a CIDR block from a VPC. To disassociate the CIDR block, you must
// specify its association ID. You can get the association ID by using DescribeVpcs
// (). You must detach or delete all gateways and resources that are associated
// with the CIDR block before you can disassociate it. You cannot disassociate the
// CIDR block with which you originally created the VPC (the primary CIDR block).
func (c *Client) DisassociateVpcCidrBlock(ctx context.Context, params *DisassociateVpcCidrBlockInput, optFns ...func(*Options)) (*DisassociateVpcCidrBlockOutput, error) {
	stack := middleware.NewStack("DisassociateVpcCidrBlock", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDisassociateVpcCidrBlockMiddlewares(stack)
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
	addOpDisassociateVpcCidrBlockValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateVpcCidrBlock(options.Region), middleware.Before)
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
			OperationName: "DisassociateVpcCidrBlock",
			Err:           err,
		}
	}
	out := result.(*DisassociateVpcCidrBlockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateVpcCidrBlockInput struct {
	// The association ID for the CIDR block.
	AssociationId *string
}

type DisassociateVpcCidrBlockOutput struct {
	// The ID of the VPC.
	VpcId *string
	// Information about the IPv4 CIDR block association.
	CidrBlockAssociation *types.VpcCidrBlockAssociation
	// Information about the IPv6 CIDR block association.
	Ipv6CidrBlockAssociation *types.VpcIpv6CidrBlockAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDisassociateVpcCidrBlockMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDisassociateVpcCidrBlock{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDisassociateVpcCidrBlock{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisassociateVpcCidrBlock(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DisassociateVpcCidrBlock",
	}
}
