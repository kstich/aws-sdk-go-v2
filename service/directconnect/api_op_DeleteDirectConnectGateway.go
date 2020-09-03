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

// Deletes the specified Direct Connect gateway. You must first delete all virtual
// interfaces that are attached to the Direct Connect gateway and disassociate all
// virtual private gateways associated with the Direct Connect gateway.
func (c *Client) DeleteDirectConnectGateway(ctx context.Context, params *DeleteDirectConnectGatewayInput, optFns ...func(*Options)) (*DeleteDirectConnectGatewayOutput, error) {
	stack := middleware.NewStack("DeleteDirectConnectGateway", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteDirectConnectGatewayMiddlewares(stack)
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
	addOpDeleteDirectConnectGatewayValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDirectConnectGateway(options.Region), middleware.Before)

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
			OperationName: "DeleteDirectConnectGateway",
			Err:           err,
		}
	}
	out := result.(*DeleteDirectConnectGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDirectConnectGatewayInput struct {
	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string
}

type DeleteDirectConnectGatewayOutput struct {
	// The Direct Connect gateway.
	DirectConnectGateway *types.DirectConnectGateway

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteDirectConnectGatewayMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteDirectConnectGateway{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteDirectConnectGateway{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteDirectConnectGateway(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "DeleteDirectConnectGateway",
	}
}
