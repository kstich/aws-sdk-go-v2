// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Accepts a resource type and returns a list of resource identifiers that are
// aggregated for a specific resource type across accounts and regions. A resource
// identifier includes the resource type, ID, (if available) the custom resource
// name, source account, and source region. You can narrow the results to include
// only resources that have specific resource IDs, or a resource name, or source
// account ID, or source region. For example, if the input consists of accountID
// 12345678910 and the region is us-east-1 for resource type AWS::EC2::Instance
// then the API returns all the EC2 instance identifiers of accountID 12345678910
// and region us-east-1.
func (c *Client) ListAggregateDiscoveredResources(ctx context.Context, params *ListAggregateDiscoveredResourcesInput, optFns ...func(*Options)) (*ListAggregateDiscoveredResourcesOutput, error) {
	stack := middleware.NewStack("ListAggregateDiscoveredResources", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListAggregateDiscoveredResourcesMiddlewares(stack)
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
	addOpListAggregateDiscoveredResourcesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAggregateDiscoveredResources(options.Region), middleware.Before)

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
			OperationName: "ListAggregateDiscoveredResources",
			Err:           err,
		}
	}
	out := result.(*ListAggregateDiscoveredResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAggregateDiscoveredResourcesInput struct {
	// Filters the results based on the ResourceFilters object.
	Filters *types.ResourceFilters
	// The type of resources that you want AWS Config to list in the response.
	ResourceType types.ResourceType
	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
	// The maximum number of resource identifiers returned on each page. The default is
	// 100. You cannot specify a number greater than 100. If you specify 0, AWS Config
	// uses the default.
	Limit *int32
	// The name of the configuration aggregator.
	ConfigurationAggregatorName *string
}

type ListAggregateDiscoveredResourcesOutput struct {
	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
	// Returns a list of ResourceIdentifiers objects.
	ResourceIdentifiers []*types.AggregateResourceIdentifier

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListAggregateDiscoveredResourcesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListAggregateDiscoveredResources{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAggregateDiscoveredResources{}, middleware.After)
}

func newServiceMetadataMiddleware_opListAggregateDiscoveredResources(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "ListAggregateDiscoveredResources",
	}
}
