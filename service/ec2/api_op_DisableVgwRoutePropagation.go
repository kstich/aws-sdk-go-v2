// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables a virtual private gateway (VGW) from propagating routes to a specified
// route table of a VPC.
func (c *Client) DisableVgwRoutePropagation(ctx context.Context, params *DisableVgwRoutePropagationInput, optFns ...func(*Options)) (*DisableVgwRoutePropagationOutput, error) {
	stack := middleware.NewStack("DisableVgwRoutePropagation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDisableVgwRoutePropagationMiddlewares(stack)
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
	addOpDisableVgwRoutePropagationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisableVgwRoutePropagation(options.Region), middleware.Before)
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
			OperationName: "DisableVgwRoutePropagation",
			Err:           err,
		}
	}
	out := result.(*DisableVgwRoutePropagationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DisableVgwRoutePropagation.
type DisableVgwRoutePropagationInput struct {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The ID of the virtual private gateway.
	GatewayId *string
	// The ID of the route table.
	RouteTableId *string
}

type DisableVgwRoutePropagationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDisableVgwRoutePropagationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDisableVgwRoutePropagation{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDisableVgwRoutePropagation{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisableVgwRoutePropagation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DisableVgwRoutePropagation",
	}
}
