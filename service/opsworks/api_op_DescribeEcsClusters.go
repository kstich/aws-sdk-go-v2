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

// Describes Amazon ECS clusters that are registered with a stack. If you specify
// only a stack ID, you can use the MaxResults and NextToken parameters to paginate
// the response. However, AWS OpsWorks Stacks currently supports only one cluster
// per layer, so the result set has a maximum of one element. Required Permissions:
// To use this action, an IAM user must have a Show, Deploy, or Manage permissions
// level for the stack or an attached policy that explicitly grants permission. For
// more information about user permissions, see Managing User Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
// This call accepts only one resource-identifying parameter.
func (c *Client) DescribeEcsClusters(ctx context.Context, params *DescribeEcsClustersInput, optFns ...func(*Options)) (*DescribeEcsClustersOutput, error) {
	stack := middleware.NewStack("DescribeEcsClusters", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeEcsClustersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEcsClusters(options.Region), middleware.Before)

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
			OperationName: "DescribeEcsClusters",
			Err:           err,
		}
	}
	out := result.(*DescribeEcsClustersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEcsClustersInput struct {
	// A list of ARNs, one for each cluster to be described.
	EcsClusterArns []*string
	// If the previous paginated request did not return all of the remaining results,
	// the response object'sNextToken parameter value is set to a token. To retrieve
	// the next set of results, call DescribeEcsClusters again and assign that token to
	// the request object's NextToken parameter. If there are no remaining results, the
	// previous response object's NextToken parameter is set to null.
	NextToken *string
	// A stack ID. DescribeEcsClusters returns a description of the cluster that is
	// registered with the stack.
	StackId *string
	// To receive a paginated response, use this parameter to specify the maximum
	// number of results to be returned with a single call. If the number of available
	// results exceeds this maximum, the response includes a NextToken value that you
	// can assign to the NextToken request parameter to get the next set of results.
	MaxResults *int32
}

// Contains the response to a DescribeEcsClusters request.
type DescribeEcsClustersOutput struct {
	// A list of EcsCluster objects containing the cluster descriptions.
	EcsClusters []*types.EcsCluster
	// If a paginated request does not return all of the remaining results, this
	// parameter is set to a token that you can assign to the request object's
	// NextToken parameter to retrieve the next set of results. If the previous
	// paginated request returned all of the remaining results, this parameter is set
	// to null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeEcsClustersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEcsClusters{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEcsClusters{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeEcsClusters(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "DescribeEcsClusters",
	}
}
