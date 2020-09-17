// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsmv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about AWS CloudHSM clusters. This is a paginated operation,
// which means that each response might contain only a subset of all the clusters.
// When the response contains only a subset of clusters, it includes a NextToken
// value. Use this value in a subsequent DescribeClusters request to get more
// clusters. When you receive a response with no NextToken (or an empty or null
// value), that means there are no more clusters to get.
func (c *Client) DescribeClusters(ctx context.Context, params *DescribeClustersInput, optFns ...func(*Options)) (*DescribeClustersOutput, error) {
	stack := middleware.NewStack("DescribeClusters", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeClustersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClusters(options.Region), middleware.Before)
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
			OperationName: "DescribeClusters",
			Err:           err,
		}
	}
	out := result.(*DescribeClustersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeClustersInput struct {
	// One or more filters to limit the items returned in the response. Use the
	// clusterIds filter to return only the specified clusters. Specify clusters by
	// their cluster identifier (ID). Use the vpcIds filter to return only the clusters
	// in the specified virtual private clouds (VPCs). Specify VPCs by their VPC
	// identifier (ID). Use the states filter to return only clusters that match the
	// specified state.
	Filters map[string][]*string
	// The NextToken value that you received in the previous response. Use this value
	// to get more clusters.
	NextToken *string
	// The maximum number of clusters to return in the response. When there are more
	// clusters than the number you specify, the response contains a NextToken value.
	MaxResults *int32
}

type DescribeClustersOutput struct {
	// An opaque string that indicates that the response contains only a subset of
	// clusters. Use this value in a subsequent DescribeClusters request to get more
	// clusters.
	NextToken *string
	// A list of clusters.
	Clusters []*types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeClustersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeClusters{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeClusters{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeClusters(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudhsm",
		OperationName: "DescribeClusters",
	}
}
