// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about a work team provided by a vendor. It returns details
// about the subscription with a vendor in the AWS Marketplace.
func (c *Client) DescribeSubscribedWorkteam(ctx context.Context, params *DescribeSubscribedWorkteamInput, optFns ...func(*Options)) (*DescribeSubscribedWorkteamOutput, error) {
	stack := middleware.NewStack("DescribeSubscribedWorkteam", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeSubscribedWorkteamMiddlewares(stack)
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
	addOpDescribeSubscribedWorkteamValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSubscribedWorkteam(options.Region), middleware.Before)

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
			OperationName: "DescribeSubscribedWorkteam",
			Err:           err,
		}
	}
	out := result.(*DescribeSubscribedWorkteamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSubscribedWorkteamInput struct {
	// The Amazon Resource Name (ARN) of the subscribed work team to describe.
	WorkteamArn *string
}

type DescribeSubscribedWorkteamOutput struct {
	// A Workteam instance that contains information about the work team.
	SubscribedWorkteam *types.SubscribedWorkteam

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeSubscribedWorkteamMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeSubscribedWorkteam{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeSubscribedWorkteam{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeSubscribedWorkteam(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeSubscribedWorkteam",
	}
}