// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about the provisioned products that meet the specified
// criteria.
func (c *Client) SearchProvisionedProducts(ctx context.Context, params *SearchProvisionedProductsInput, optFns ...func(*Options)) (*SearchProvisionedProductsOutput, error) {
	stack := middleware.NewStack("SearchProvisionedProducts", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpSearchProvisionedProductsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opSearchProvisionedProducts(options.Region), middleware.Before)
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
			OperationName: "SearchProvisionedProducts",
			Err:           err,
		}
	}
	out := result.(*SearchProvisionedProductsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchProvisionedProductsInput struct {
	// The maximum number of items to return with this call.
	PageSize *int32
	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string
	// The sort order. If no value is specified, the results are not sorted.
	SortOrder types.SortOrder
	// The language code.
	//
	//     * en - English (default)
	//
	//     * jp - Japanese
	//
	//     * zh
	// - Chinese
	AcceptLanguage *string
	// The sort field. If no value is specified, the results are not sorted. The valid
	// values are arn, id, name, and lastRecordId.
	SortBy *string
	// The search filters. When the key is SearchQuery, the searchable fields are arn,
	// createdTime, id, lastRecordId, idempotencyToken, name, physicalId, productId,
	// provisioningArtifact, type, status, tags, userArn, and userArnSession. Example:
	// "SearchQuery":["status:AVAILABLE"]
	Filters map[string][]*string
	// The access level to use to obtain results. The default is User.
	AccessLevelFilter *types.AccessLevelFilter
}

type SearchProvisionedProductsOutput struct {
	// The number of provisioned products found.
	TotalResultsCount *int32
	// Information about the provisioned products.
	ProvisionedProducts []*types.ProvisionedProductAttribute
	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpSearchProvisionedProductsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpSearchProvisionedProducts{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchProvisionedProducts{}, middleware.After)
}

func newServiceMetadataMiddleware_opSearchProvisionedProducts(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "SearchProvisionedProducts",
	}
}
