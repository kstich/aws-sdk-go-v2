// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Allows the destination domain owner to reject an inbound cross-cluster search
// connection request.
func (c *Client) RejectInboundCrossClusterSearchConnection(ctx context.Context, params *RejectInboundCrossClusterSearchConnectionInput, optFns ...func(*Options)) (*RejectInboundCrossClusterSearchConnectionOutput, error) {
	stack := middleware.NewStack("RejectInboundCrossClusterSearchConnection", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpRejectInboundCrossClusterSearchConnectionMiddlewares(stack)
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
	addOpRejectInboundCrossClusterSearchConnectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRejectInboundCrossClusterSearchConnection(options.Region), middleware.Before)
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
			OperationName: "RejectInboundCrossClusterSearchConnection",
			Err:           err,
		}
	}
	out := result.(*RejectInboundCrossClusterSearchConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the RejectInboundCrossClusterSearchConnection ()
// operation.
type RejectInboundCrossClusterSearchConnectionInput struct {
	// The id of the inbound connection that you want to reject.
	CrossClusterSearchConnectionId *string
}

// The result of a RejectInboundCrossClusterSearchConnection () operation. Contains
// details of rejected inbound connection.
type RejectInboundCrossClusterSearchConnectionOutput struct {
	// Specifies the InboundCrossClusterSearchConnection () of rejected inbound
	// connection.
	CrossClusterSearchConnection *types.InboundCrossClusterSearchConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpRejectInboundCrossClusterSearchConnectionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpRejectInboundCrossClusterSearchConnection{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpRejectInboundCrossClusterSearchConnection{}, middleware.After)
}

func newServiceMetadataMiddleware_opRejectInboundCrossClusterSearchConnection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "RejectInboundCrossClusterSearchConnection",
	}
}
