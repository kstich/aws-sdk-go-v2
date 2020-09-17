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

// Describes the data feed for Spot Instances. For more information, see Spot
// Instance data feed
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-data-feeds.html) in
// the Amazon EC2 User Guide for Linux Instances.
func (c *Client) DescribeSpotDatafeedSubscription(ctx context.Context, params *DescribeSpotDatafeedSubscriptionInput, optFns ...func(*Options)) (*DescribeSpotDatafeedSubscriptionOutput, error) {
	stack := middleware.NewStack("DescribeSpotDatafeedSubscription", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeSpotDatafeedSubscriptionMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSpotDatafeedSubscription(options.Region), middleware.Before)
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
			OperationName: "DescribeSpotDatafeedSubscription",
			Err:           err,
		}
	}
	out := result.(*DescribeSpotDatafeedSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeSpotDatafeedSubscription.
type DescribeSpotDatafeedSubscriptionInput struct {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

// Contains the output of DescribeSpotDatafeedSubscription.
type DescribeSpotDatafeedSubscriptionOutput struct {
	// The Spot Instance data feed subscription.
	SpotDatafeedSubscription *types.SpotDatafeedSubscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeSpotDatafeedSubscriptionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeSpotDatafeedSubscription{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeSpotDatafeedSubscription{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeSpotDatafeedSubscription(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeSpotDatafeedSubscription",
	}
}
