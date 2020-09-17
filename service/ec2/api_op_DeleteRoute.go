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

// Deletes the specified route from the specified route table.
func (c *Client) DeleteRoute(ctx context.Context, params *DeleteRouteInput, optFns ...func(*Options)) (*DeleteRouteOutput, error) {
	stack := middleware.NewStack("DeleteRoute", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDeleteRouteMiddlewares(stack)
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
	addOpDeleteRouteValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRoute(options.Region), middleware.Before)
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
			OperationName: "DeleteRoute",
			Err:           err,
		}
	}
	out := result.(*DeleteRouteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRouteInput struct {
	// The IPv6 CIDR range for the route. The value you specify must match the CIDR for
	// the route exactly.
	DestinationIpv6CidrBlock *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The IPv4 CIDR range for the route. The value you specify must match the CIDR for
	// the route exactly.
	DestinationCidrBlock *string
	// The ID of the route table.
	RouteTableId *string
	// The ID of the prefix list for the route.
	DestinationPrefixListId *string
}

type DeleteRouteOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDeleteRouteMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDeleteRoute{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDeleteRoute{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteRoute(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeleteRoute",
	}
}
