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

// Unassigns one or more secondary private IP addresses from a network interface.
func (c *Client) UnassignPrivateIpAddresses(ctx context.Context, params *UnassignPrivateIpAddressesInput, optFns ...func(*Options)) (*UnassignPrivateIpAddressesOutput, error) {
	stack := middleware.NewStack("UnassignPrivateIpAddresses", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpUnassignPrivateIpAddressesMiddlewares(stack)
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
	addOpUnassignPrivateIpAddressesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUnassignPrivateIpAddresses(options.Region), middleware.Before)
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
			OperationName: "UnassignPrivateIpAddresses",
			Err:           err,
		}
	}
	out := result.(*UnassignPrivateIpAddressesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for UnassignPrivateIpAddresses.
type UnassignPrivateIpAddressesInput struct {
	// The ID of the network interface.
	NetworkInterfaceId *string
	// The secondary private IP addresses to unassign from the network interface. You
	// can specify this option multiple times to unassign more than one IP address.
	PrivateIpAddresses []*string
}

type UnassignPrivateIpAddressesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpUnassignPrivateIpAddressesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpUnassignPrivateIpAddresses{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpUnassignPrivateIpAddresses{}, middleware.After)
}

func newServiceMetadataMiddleware_opUnassignPrivateIpAddresses(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "UnassignPrivateIpAddresses",
	}
}
