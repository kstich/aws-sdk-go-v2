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

// Describes one or more attachments between resources and transit gateways. By
// default, all attachments are described. Alternatively, you can filter the
// results by attachment ID, attachment state, resource ID, or resource owner.
func (c *Client) DescribeTransitGatewayAttachments(ctx context.Context, params *DescribeTransitGatewayAttachmentsInput, optFns ...func(*Options)) (*DescribeTransitGatewayAttachmentsOutput, error) {
	stack := middleware.NewStack("DescribeTransitGatewayAttachments", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeTransitGatewayAttachmentsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTransitGatewayAttachments(options.Region), middleware.Before)
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
			OperationName: "DescribeTransitGatewayAttachments",
			Err:           err,
		}
	}
	out := result.(*DescribeTransitGatewayAttachmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTransitGatewayAttachmentsInput struct {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32
	// The token for the next page of results.
	NextToken *string
	// One or more filters. The possible values are:
	//
	//     * association.state - The
	// state of the association (associating | associated | disassociating).
	//
	//     *
	// association.transit-gateway-route-table-id - The ID of the route table for the
	// transit gateway.
	//
	//     * resource-id - The ID of the resource.
	//
	//     *
	// resource-owner-id - The ID of the AWS account that owns the resource.
	//
	//     *
	// resource-type - The resource type (vpc | vpn).
	//
	//     * state - The state of the
	// attachment (available | deleted | deleting | failed | modifying |
	// pendingAcceptance | pending | rollingBack | rejected | rejecting).
	//
	//     *
	// transit-gateway-attachment-id - The ID of the attachment.
	//
	//     *
	// transit-gateway-id - The ID of the transit gateway.
	//
	//     *
	// transit-gateway-owner-id - The ID of the AWS account that owns the transit
	// gateway.
	Filters []*types.Filter
	// The IDs of the attachments.
	TransitGatewayAttachmentIds []*string
}

type DescribeTransitGatewayAttachmentsOutput struct {
	// Information about the attachments.
	TransitGatewayAttachments []*types.TransitGatewayAttachment
	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeTransitGatewayAttachmentsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeTransitGatewayAttachments{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeTransitGatewayAttachments{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeTransitGatewayAttachments(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeTransitGatewayAttachments",
	}
}
