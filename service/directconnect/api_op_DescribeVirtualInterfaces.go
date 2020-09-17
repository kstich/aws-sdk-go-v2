// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Displays all virtual interfaces for an AWS account. Virtual interfaces deleted
// fewer than 15 minutes before you make the request are also returned. If you
// specify a connection ID, only the virtual interfaces associated with the
// connection are returned. If you specify a virtual interface ID, then only a
// single virtual interface is returned. A virtual interface (VLAN) transmits the
// traffic between the AWS Direct Connect location and the customer network.
func (c *Client) DescribeVirtualInterfaces(ctx context.Context, params *DescribeVirtualInterfacesInput, optFns ...func(*Options)) (*DescribeVirtualInterfacesOutput, error) {
	stack := middleware.NewStack("DescribeVirtualInterfaces", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeVirtualInterfacesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVirtualInterfaces(options.Region), middleware.Before)
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
			OperationName: "DescribeVirtualInterfaces",
			Err:           err,
		}
	}
	out := result.(*DescribeVirtualInterfacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVirtualInterfacesInput struct {
	// The ID of the virtual interface.
	VirtualInterfaceId *string
	// The ID of the connection.
	ConnectionId *string
}

type DescribeVirtualInterfacesOutput struct {
	// The virtual interfaces
	VirtualInterfaces []*types.VirtualInterface

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeVirtualInterfacesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeVirtualInterfaces{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeVirtualInterfaces{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeVirtualInterfaces(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "DescribeVirtualInterfaces",
	}
}
