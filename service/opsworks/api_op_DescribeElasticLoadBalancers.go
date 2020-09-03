// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes a stack's Elastic Load Balancing instances. This call accepts only one
// resource-identifying parameter. Required Permissions: To use this action, an IAM
// user must have a Show, Deploy, or Manage permissions level for the stack, or an
// attached policy that explicitly grants permissions. For more information about
// user permissions, see Managing User Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) DescribeElasticLoadBalancers(ctx context.Context, params *DescribeElasticLoadBalancersInput, optFns ...func(*Options)) (*DescribeElasticLoadBalancersOutput, error) {
	stack := middleware.NewStack("DescribeElasticLoadBalancers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeElasticLoadBalancersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeElasticLoadBalancers(options.Region), middleware.Before)

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
			OperationName: "DescribeElasticLoadBalancers",
			Err:           err,
		}
	}
	out := result.(*DescribeElasticLoadBalancersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeElasticLoadBalancersInput struct {
	// A stack ID. The action describes the stack's Elastic Load Balancing instances.
	StackId *string
	// A list of layer IDs. The action describes the Elastic Load Balancing instances
	// for the specified layers.
	LayerIds []*string
}

// Contains the response to a DescribeElasticLoadBalancers request.
type DescribeElasticLoadBalancersOutput struct {
	// A list of ElasticLoadBalancer objects that describe the specified Elastic Load
	// Balancing instances.
	ElasticLoadBalancers []*types.ElasticLoadBalancer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeElasticLoadBalancersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeElasticLoadBalancers{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeElasticLoadBalancers{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeElasticLoadBalancers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "DescribeElasticLoadBalancers",
	}
}
