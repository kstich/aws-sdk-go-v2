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

// Describes the connection notifications for VPC endpoints and VPC endpoint
// services.
func (c *Client) DescribeVpcEndpointConnectionNotifications(ctx context.Context, params *DescribeVpcEndpointConnectionNotificationsInput, optFns ...func(*Options)) (*DescribeVpcEndpointConnectionNotificationsOutput, error) {
	stack := middleware.NewStack("DescribeVpcEndpointConnectionNotifications", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeVpcEndpointConnectionNotificationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpcEndpointConnectionNotifications(options.Region), middleware.Before)
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
			OperationName: "DescribeVpcEndpointConnectionNotifications",
			Err:           err,
		}
	}
	out := result.(*DescribeVpcEndpointConnectionNotificationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVpcEndpointConnectionNotificationsInput struct {
	// One or more filters.
	//
	//     * connection-notification-arn - The ARN of the SNS
	// topic for the notification.
	//
	//     * connection-notification-id - The ID of the
	// notification.
	//
	//     * connection-notification-state - The state of the
	// notification (Enabled | Disabled).
	//
	//     * connection-notification-type - The
	// type of notification (Topic).
	//
	//     * service-id - The ID of the endpoint
	// service.
	//
	//     * vpc-endpoint-id - The ID of the VPC endpoint.
	Filters []*types.Filter
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another request with the returned NextToken value.
	MaxResults *int32
	// The token to request the next page of results.
	NextToken *string
	// The ID of the notification.
	ConnectionNotificationId *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DescribeVpcEndpointConnectionNotificationsOutput struct {
	// One or more notifications.
	ConnectionNotificationSet []*types.ConnectionNotification
	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeVpcEndpointConnectionNotificationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeVpcEndpointConnectionNotifications{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVpcEndpointConnectionNotifications{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeVpcEndpointConnectionNotifications(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVpcEndpointConnectionNotifications",
	}
}
