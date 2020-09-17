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

// Deletes the association between the specified Direct Connect gateway and virtual
// private gateway. We recommend that you specify the associationID to delete the
// association. Alternatively, if you own virtual gateway and a Direct Connect
// gateway association, you can specify the virtualGatewayId and
// directConnectGatewayId to delete an association.
func (c *Client) DeleteDirectConnectGatewayAssociation(ctx context.Context, params *DeleteDirectConnectGatewayAssociationInput, optFns ...func(*Options)) (*DeleteDirectConnectGatewayAssociationOutput, error) {
	stack := middleware.NewStack("DeleteDirectConnectGatewayAssociation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteDirectConnectGatewayAssociationMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDirectConnectGatewayAssociation(options.Region), middleware.Before)
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
			OperationName: "DeleteDirectConnectGatewayAssociation",
			Err:           err,
		}
	}
	out := result.(*DeleteDirectConnectGatewayAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDirectConnectGatewayAssociationInput struct {
	// The ID of the virtual private gateway.
	VirtualGatewayId *string
	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string
	// The ID of the Direct Connect gateway association.
	AssociationId *string
}

type DeleteDirectConnectGatewayAssociationOutput struct {
	// Information about the deleted association.
	DirectConnectGatewayAssociation *types.DirectConnectGatewayAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteDirectConnectGatewayAssociationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteDirectConnectGatewayAssociation{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteDirectConnectGatewayAssociation{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteDirectConnectGatewayAssociation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "DeleteDirectConnectGatewayAssociation",
	}
}
