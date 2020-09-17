// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables a tape gateway when the gateway is no longer functioning. For example,
// if your gateway VM is damaged, you can disable the gateway so you can recover
// virtual tapes.  <p>Use this operation for a tape gateway that is not reachable
// or not functioning. This operation is only supported in the tape gateway
// type.</p> <important> <p>After a gateway is disabled, it cannot be enabled.</p>
// </important>
func (c *Client) DisableGateway(ctx context.Context, params *DisableGatewayInput, optFns ...func(*Options)) (*DisableGatewayOutput, error) {
	stack := middleware.NewStack("DisableGateway", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDisableGatewayMiddlewares(stack)
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
	addOpDisableGatewayValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisableGateway(options.Region), middleware.Before)
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
			OperationName: "DisableGateway",
			Err:           err,
		}
	}
	out := result.(*DisableGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DisableGatewayInput
type DisableGatewayInput struct {
	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string
}

// DisableGatewayOutput
type DisableGatewayOutput struct {
	// The unique Amazon Resource Name (ARN) of the disabled gateway.
	GatewayARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDisableGatewayMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDisableGateway{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisableGateway{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisableGateway(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DisableGateway",
	}
}
