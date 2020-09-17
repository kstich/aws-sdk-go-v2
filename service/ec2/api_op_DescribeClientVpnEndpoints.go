// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes one or more Client VPN endpoints in the account.
func (c *Client) DescribeClientVpnEndpoints(ctx context.Context, params *DescribeClientVpnEndpointsInput, optFns ...func(*Options)) (*DescribeClientVpnEndpointsOutput, error) {
	stack := middleware.NewStack("DescribeClientVpnEndpoints", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeClientVpnEndpointsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClientVpnEndpoints(options.Region), middleware.Before)
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
			OperationName: "DescribeClientVpnEndpoints",
			Err:           err,
		}
	}
	out := result.(*DescribeClientVpnEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeClientVpnEndpointsInput struct {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The token to retrieve the next page of results.
	NextToken *string
	// One or more filters. Filter names and values are case-sensitive.
	//
	//     *
	// endpoint-id - The ID of the Client VPN endpoint.
	//
	//     * transport-protocol - The
	// transport protocol (tcp | udp).
	Filters []*types.Filter
	// The maximum number of results to return for the request in a single page. The
	// remaining results can be seen by sending another request with the nextToken
	// value.
	MaxResults *int32
	// The ID of the Client VPN endpoint.
	ClientVpnEndpointIds []*string
}

type DescribeClientVpnEndpointsOutput struct {
	// Information about the Client VPN endpoints.
	ClientVpnEndpoints []*types.ClientVpnEndpoint
	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeClientVpnEndpointsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeClientVpnEndpoints{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeClientVpnEndpoints{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeClientVpnEndpoints(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeClientVpnEndpoints",
	}
}
