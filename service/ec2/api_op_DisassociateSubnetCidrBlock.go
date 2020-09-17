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

// Disassociates a CIDR block from a subnet. Currently, you can disassociate an
// IPv6 CIDR block only. You must detach or delete all gateways and resources that
// are associated with the CIDR block before you can disassociate it.
func (c *Client) DisassociateSubnetCidrBlock(ctx context.Context, params *DisassociateSubnetCidrBlockInput, optFns ...func(*Options)) (*DisassociateSubnetCidrBlockOutput, error) {
	stack := middleware.NewStack("DisassociateSubnetCidrBlock", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDisassociateSubnetCidrBlockMiddlewares(stack)
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
	addOpDisassociateSubnetCidrBlockValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateSubnetCidrBlock(options.Region), middleware.Before)
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
			OperationName: "DisassociateSubnetCidrBlock",
			Err:           err,
		}
	}
	out := result.(*DisassociateSubnetCidrBlockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateSubnetCidrBlockInput struct {
	// The association ID for the CIDR block.
	AssociationId *string
}

type DisassociateSubnetCidrBlockOutput struct {
	// Information about the IPv6 CIDR block association.
	Ipv6CidrBlockAssociation *types.SubnetIpv6CidrBlockAssociation
	// The ID of the subnet.
	SubnetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDisassociateSubnetCidrBlockMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDisassociateSubnetCidrBlock{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDisassociateSubnetCidrBlock{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisassociateSubnetCidrBlock(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DisassociateSubnetCidrBlock",
	}
}
