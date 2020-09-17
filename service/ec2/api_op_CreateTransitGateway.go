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

// Creates a transit gateway. You can use a transit gateway to interconnect your
// virtual private clouds (VPC) and on-premises networks. After the transit gateway
// enters the available state, you can attach your VPCs and VPN connections to the
// transit gateway. To attach your VPCs, use CreateTransitGatewayVpcAttachment ().
// To attach a VPN connection, use CreateCustomerGateway () to create a customer
// gateway and specify the ID of the customer gateway and the ID of the transit
// gateway in a call to CreateVpnConnection (). When you create a transit gateway,
// we create a default transit gateway route table and use it as the default
// association route table and the default propagation route table. You can use
// CreateTransitGatewayRouteTable () to create additional transit gateway route
// tables. If you disable automatic route propagation, we do not create a default
// transit gateway route table. You can use
// EnableTransitGatewayRouteTablePropagation () to propagate routes from a resource
// attachment to a transit gateway route table. If you disable automatic
// associations, you can use AssociateTransitGatewayRouteTable () to associate a
// resource attachment with a transit gateway route table.
func (c *Client) CreateTransitGateway(ctx context.Context, params *CreateTransitGatewayInput, optFns ...func(*Options)) (*CreateTransitGatewayOutput, error) {
	stack := middleware.NewStack("CreateTransitGateway", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpCreateTransitGatewayMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTransitGateway(options.Region), middleware.Before)
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
			OperationName: "CreateTransitGateway",
			Err:           err,
		}
	}
	out := result.(*CreateTransitGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTransitGatewayInput struct {
	// The tags to apply to the transit gateway.
	TagSpecifications []*types.TagSpecification
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The transit gateway options.
	Options *types.TransitGatewayRequestOptions
	// A description of the transit gateway.
	Description *string
}

type CreateTransitGatewayOutput struct {
	// Information about the transit gateway.
	TransitGateway *types.TransitGateway

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpCreateTransitGatewayMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpCreateTransitGateway{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpCreateTransitGateway{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateTransitGateway(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateTransitGateway",
	}
}
