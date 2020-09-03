// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the specified policies. If you specify a load balancer name, the
// action returns the descriptions of all policies created for the load balancer.
// If you specify a policy name associated with your load balancer, the action
// returns the description of that policy. If you don't specify a load balancer
// name, the action returns descriptions of the specified sample policies, or
// descriptions of all sample policies. The names of the sample policies have the
// ELBSample- prefix.
func (c *Client) DescribeLoadBalancerPolicies(ctx context.Context, params *DescribeLoadBalancerPoliciesInput, optFns ...func(*Options)) (*DescribeLoadBalancerPoliciesOutput, error) {
	stack := middleware.NewStack("DescribeLoadBalancerPolicies", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeLoadBalancerPoliciesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLoadBalancerPolicies(options.Region), middleware.Before)

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
			OperationName: "DescribeLoadBalancerPolicies",
			Err:           err,
		}
	}
	out := result.(*DescribeLoadBalancerPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeLoadBalancerPolicies.
type DescribeLoadBalancerPoliciesInput struct {
	// The name of the load balancer.
	LoadBalancerName *string
	// The names of the policies.
	PolicyNames []*string
}

// Contains the output of DescribeLoadBalancerPolicies.
type DescribeLoadBalancerPoliciesOutput struct {
	// Information about the policies.
	PolicyDescriptions []*types.PolicyDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeLoadBalancerPoliciesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeLoadBalancerPolicies{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeLoadBalancerPolicies{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeLoadBalancerPolicies(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "DescribeLoadBalancerPolicies",
	}
}
