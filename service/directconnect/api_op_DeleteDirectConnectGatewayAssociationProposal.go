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

// Deletes the association proposal request between the specified Direct Connect
// gateway and virtual private gateway or transit gateway.
func (c *Client) DeleteDirectConnectGatewayAssociationProposal(ctx context.Context, params *DeleteDirectConnectGatewayAssociationProposalInput, optFns ...func(*Options)) (*DeleteDirectConnectGatewayAssociationProposalOutput, error) {
	stack := middleware.NewStack("DeleteDirectConnectGatewayAssociationProposal", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteDirectConnectGatewayAssociationProposalMiddlewares(stack)
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
	addOpDeleteDirectConnectGatewayAssociationProposalValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDirectConnectGatewayAssociationProposal(options.Region), middleware.Before)

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
			OperationName: "DeleteDirectConnectGatewayAssociationProposal",
			Err:           err,
		}
	}
	out := result.(*DeleteDirectConnectGatewayAssociationProposalOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDirectConnectGatewayAssociationProposalInput struct {
	// The ID of the proposal.
	ProposalId *string
}

type DeleteDirectConnectGatewayAssociationProposalOutput struct {
	// The ID of the associated gateway.
	DirectConnectGatewayAssociationProposal *types.DirectConnectGatewayAssociationProposal

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteDirectConnectGatewayAssociationProposalMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteDirectConnectGatewayAssociationProposal{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteDirectConnectGatewayAssociationProposal{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteDirectConnectGatewayAssociationProposal(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "DeleteDirectConnectGatewayAssociationProposal",
	}
}
