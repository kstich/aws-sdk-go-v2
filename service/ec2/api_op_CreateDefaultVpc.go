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

// Creates a default VPC with a size /16 IPv4 CIDR block and a default subnet in
// each Availability Zone. For more information about the components of a default
// VPC, see Default VPC and Default Subnets
// (https://docs.aws.amazon.com/vpc/latest/userguide/default-vpc.html) in the
// Amazon Virtual Private Cloud User Guide. You cannot specify the components of
// the default VPC yourself. If you deleted your previous default VPC, you can
// create a default VPC. You cannot have more than one default VPC per Region. If
// your account supports EC2-Classic, you cannot use this action to create a
// default VPC in a Region that supports EC2-Classic. If you want a default VPC in
// a Region that supports EC2-Classic, see "I really want a default VPC for my
// existing EC2 account. Is that possible?" in the Default VPCs FAQ
// (http://aws.amazon.com/vpc/faqs/#Default_VPCs).
func (c *Client) CreateDefaultVpc(ctx context.Context, params *CreateDefaultVpcInput, optFns ...func(*Options)) (*CreateDefaultVpcOutput, error) {
	stack := middleware.NewStack("CreateDefaultVpc", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpCreateDefaultVpcMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDefaultVpc(options.Region), middleware.Before)
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
			OperationName: "CreateDefaultVpc",
			Err:           err,
		}
	}
	out := result.(*CreateDefaultVpcOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDefaultVpcInput struct {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type CreateDefaultVpcOutput struct {
	// Information about the VPC.
	Vpc *types.Vpc

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpCreateDefaultVpcMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpCreateDefaultVpc{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpCreateDefaultVpc{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDefaultVpc(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateDefaultVpc",
	}
}
