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

// Accepts a resource type and returns a list of resource identifiers for the
// resources of that type. A resource identifier includes the resource type, ID,
// and (if available) the custom resource name. The results consist of resources
// that AWS Config has discovered, including those that AWS Config is not currently
// recording. You can narrow the results to include only resources that have
// specific resource IDs or a resource name. You can specify either resource IDs or
// a resource name, but not both, in the same request. The response is paginated.
// By default, AWS Config lists 100 resource identifiers on each page. You can
// customize this number with the limit parameter. The response includes a
// nextToken string. To get the next page of results, run the request again and
// specify the string for the nextToken parameter.
func (c *Client) ListDiscoveredResources(ctx context.Context, params *ListDiscoveredResourcesInput, optFns ...func(*Options)) (*ListDiscoveredResourcesOutput, error) {
	stack := middleware.NewStack("ListDiscoveredResources", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListDiscoveredResourcesMiddlewares(stack)
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
	addOpListDiscoveredResourcesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDiscoveredResources(options.Region), middleware.Before)

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
			OperationName: "ListDiscoveredResources",
			Err:           err,
		}
	}
	out := result.(*ListDiscoveredResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ListDiscoveredResourcesInput struct {
	// The maximum number of resource identifiers returned on each page. The default is
	// 100. You cannot specify a number greater than 100. If you specify 0, AWS Config
	// uses the default.
	Limit *int32
	// Specifies whether AWS Config includes deleted resources in the results. By
	// default, deleted resources are not included.
	IncludeDeletedResources *bool
	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
	// The type of resources that you want AWS Config to list in the response.
	ResourceType types.ResourceType
	// The IDs of only those resources that you want AWS Config to list in the
	// response. If you do not specify this parameter, AWS Config lists all resources
	// of the specified type that it has discovered.
	ResourceIds []*string
	// The custom name of only those resources that you want AWS Config to list in the
	// response. If you do not specify this parameter, AWS Config lists all resources
	// of the specified type that it has discovered.
	ResourceName *string
}

//
type ListDiscoveredResourcesOutput struct {
	// The string that you use in a subsequent request to get the next page of results
	// in a paginated response.
	NextToken *string
	// The details that identify a resource that is discovered by AWS Config, including
	// the resource type, ID, and (if available) the custom resource name.
	ResourceIdentifiers []*types.ResourceIdentifier

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListDiscoveredResourcesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListDiscoveredResources{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDiscoveredResources{}, middleware.After)
}

func newServiceMetadataMiddleware_opListDiscoveredResources(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "ListDiscoveredResources",
	}
}
