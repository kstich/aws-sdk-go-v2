// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the network profile details by the network profile ARN.
func (c *Client) GetNetworkProfile(ctx context.Context, params *GetNetworkProfileInput, optFns ...func(*Options)) (*GetNetworkProfileOutput, error) {
	stack := middleware.NewStack("GetNetworkProfile", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetNetworkProfileMiddlewares(stack)
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
	addOpGetNetworkProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetNetworkProfile(options.Region), middleware.Before)
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
			OperationName: "GetNetworkProfile",
			Err:           err,
		}
	}
	out := result.(*GetNetworkProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetNetworkProfileInput struct {
	// The ARN of the network profile associated with a device.
	NetworkProfileArn *string
}

type GetNetworkProfileOutput struct {
	// The network profile associated with a device.
	NetworkProfile *types.NetworkProfile

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetNetworkProfileMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetNetworkProfile{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetNetworkProfile{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetNetworkProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "GetNetworkProfile",
	}
}
