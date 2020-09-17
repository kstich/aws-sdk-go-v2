// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Establishes a VPC peering connection between a virtual private cloud (VPC) in an
// AWS account with the VPC for your Amazon GameLift fleet. VPC peering enables the
// game servers on your fleet to communicate directly with other AWS resources. You
// can peer with VPCs in any AWS account that you have access to, including the
// account that you use to manage your Amazon GameLift fleets. You cannot peer with
// VPCs that are in different Regions. For more information, see VPC Peering with
// Amazon GameLift Fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/vpc-peering.html).
// Before calling this operation to establish the peering connection, you first
// need to call CreateVpcPeeringAuthorization () and identify the VPC you want to
// peer with. Once the authorization for the specified VPC is issued, you have 24
// hours to establish the connection. These two operations handle all tasks
// necessary to peer the two VPCs, including acceptance, updating routing tables,
// etc. To establish the connection, call this operation from the AWS account that
// is used to manage the Amazon GameLift fleets. Identify the following values: (1)
// The ID of the fleet you want to be enable a VPC peering connection for; (2) The
// AWS account with the VPC that you want to peer with; and (3) The ID of the VPC
// you want to peer with. This operation is asynchronous. If successful, a
// VpcPeeringConnection () request is created. You can use continuous polling to
// track the request's status using DescribeVpcPeeringConnections (), or by
// monitoring fleet events for success or failure using DescribeFleetEvents ().
//
//
// * CreateVpcPeeringAuthorization ()
//
//     * DescribeVpcPeeringAuthorizations ()
//
//
// * DeleteVpcPeeringAuthorization ()
//
//     * CreateVpcPeeringConnection ()
//
//     *
// DescribeVpcPeeringConnections ()
//
//     * DeleteVpcPeeringConnection ()
func (c *Client) CreateVpcPeeringConnection(ctx context.Context, params *CreateVpcPeeringConnectionInput, optFns ...func(*Options)) (*CreateVpcPeeringConnectionOutput, error) {
	stack := middleware.NewStack("CreateVpcPeeringConnection", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateVpcPeeringConnectionMiddlewares(stack)
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
	addOpCreateVpcPeeringConnectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVpcPeeringConnection(options.Region), middleware.Before)
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
			OperationName: "CreateVpcPeeringConnection",
			Err:           err,
		}
	}
	out := result.(*CreateVpcPeeringConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type CreateVpcPeeringConnectionInput struct {
	// A unique identifier for a fleet. You can use either the fleet ID or ARN value.
	// This tells Amazon GameLift which GameLift VPC to peer with.
	FleetId *string
	// A unique identifier for a VPC with resources to be accessed by your Amazon
	// GameLift fleet. The VPC must be in the same Region where your fleet is deployed.
	// Look up a VPC ID using the VPC Dashboard (https://console.aws.amazon.com/vpc/)
	// in the AWS Management Console. Learn more about VPC peering in VPC Peering with
	// Amazon GameLift Fleets
	// (https://docs.aws.amazon.com/gamelift/latest/developerguide/vpc-peering.html).
	PeerVpcId *string
	// A unique identifier for the AWS account with the VPC that you want to peer your
	// Amazon GameLift fleet with. You can find your Account ID in the AWS Management
	// Console under account settings.
	PeerVpcAwsAccountId *string
}

type CreateVpcPeeringConnectionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateVpcPeeringConnectionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateVpcPeeringConnection{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateVpcPeeringConnection{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateVpcPeeringConnection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "CreateVpcPeeringConnection",
	}
}
