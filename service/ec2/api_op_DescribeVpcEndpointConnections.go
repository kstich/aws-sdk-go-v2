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

// Describes the VPC endpoint connections to your VPC endpoint services, including
// any endpoints that are pending your acceptance.
func (c *Client) DescribeVpcEndpointConnections(ctx context.Context, params *DescribeVpcEndpointConnectionsInput, optFns ...func(*Options)) (*DescribeVpcEndpointConnectionsOutput, error) {
	stack := middleware.NewStack("DescribeVpcEndpointConnections", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeVpcEndpointConnectionsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpcEndpointConnections(options.Region), middleware.Before)

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
			OperationName: "DescribeVpcEndpointConnections",
			Err:           err,
		}
	}
	out := result.(*DescribeVpcEndpointConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVpcEndpointConnectionsInput struct {
	// One or more filters.
	//
	//     * service-id - The ID of the service.
	//
	//     *
	// vpc-endpoint-owner - The AWS account number of the owner of the endpoint.
	//
	//     *
	// vpc-endpoint-state - The state of the endpoint (pendingAcceptance | pending |
	// available | deleting | deleted | rejected | failed).
	//
	//     * vpc-endpoint-id -
	// The ID of the endpoint.
	Filters []*types.Filter
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The maximum number of results to return for the request in a single page. The
	// remaining results of the initial request can be seen by sending another request
	// with the returned NextToken value. This value can be between 5 and 1,000; if
	// MaxResults is given a value larger than 1,000, only 1,000 results are returned.
	MaxResults *int32
	// The token to retrieve the next page of results.
	NextToken *string
}

type DescribeVpcEndpointConnectionsOutput struct {
	// Information about one or more VPC endpoint connections.
	VpcEndpointConnections []*types.VpcEndpointConnection
	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeVpcEndpointConnectionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeVpcEndpointConnections{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVpcEndpointConnections{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeVpcEndpointConnections(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVpcEndpointConnections",
	}
}