// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Searches an active index. Use this API to search your documents using query. The
// Query operation enables to do faceted search and to filter results based on
// document attributes. It also enables you to provide user context that Amazon
// Kendra uses to enforce document access control in the search results. Amazon
// Kendra searches your index for text content and question and answer (FAQ)
// content. By default the response contains three types of results.
//
//     *
// Relevant passages
//
//     * Matching FAQs
//
//     * Relevant documents
//
// You can
// specify that the query return only one type of result using the
// QueryResultTypeConfig parameter.
func (c *Client) Query(ctx context.Context, params *QueryInput, optFns ...func(*Options)) (*QueryOutput, error) {
	stack := middleware.NewStack("Query", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpQueryMiddlewares(stack)
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
	addOpQueryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opQuery(options.Region), middleware.Before)
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
			OperationName: "Query",
			Err:           err,
		}
	}
	out := result.(*QueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type QueryInput struct {
	// Provides information that determines how the results of the query are sorted.
	// You can set the field that Amazon Kendra should sort the results on, and specify
	// whether the results should be sorted in ascending or descending order. In the
	// case of ties in sorting the results, the results are sorted by relevance. If you
	// don't provide sorting configuration, the results are sorted by the relevance
	// that Amazon Kendra determines for the result.
	SortingConfiguration *types.SortingConfiguration
	// The unique identifier of the index to search. The identifier is returned in the
	// response from the operation.
	IndexId *string
	// The text to search for.
	QueryText *string
	// Enables filtered searches based on document attributes. You can only provide one
	// attribute filter; however, the AndAllFilters, NotFilter, and OrAllFilters
	// parameters contain a list of other filters. The AttributeFilter parameter
	// enables you to create a set of filtering rules that a document must satisfy to
	// be included in the query results.
	AttributeFilter *types.AttributeFilter
	// Query results are returned in pages the size of the PageSize parameter. By
	// default, Amazon Kendra returns the first page of results. Use this parameter to
	// get result pages after the first one.
	PageNumber *int32
	// Sets the number of results that are returned in each page of results. The
	// default page size is 10. The maximum number of results returned is 100. If you
	// ask for more than 100 results, only 100 are returned.
	PageSize *int32
	// An array of documents attributes. Amazon Kendra returns a count for each
	// attribute key specified. You can use this information to help narrow the search
	// for your user.
	Facets []*types.Facet
	// Sets the type of query. Only results for the specified query type are returned.
	QueryResultTypeFilter types.QueryResultType
	// An array of document attributes to include in the response. No other document
	// attributes are included in the response. By default all document attributes are
	// included in the response.
	RequestedDocumentAttributes []*string
}

type QueryOutput struct {
	// The results of the search.
	ResultItems []*types.QueryResultItem
	// Contains the facet results. A FacetResult contains the counts for each attribute
	// key that was specified in the Facets input parameter.
	FacetResults []*types.FacetResult
	// The number of items returned by the search. Use this to determine when you have
	// requested the last set of results.
	TotalNumberOfResults *int32
	// The unique identifier for the search. You use QueryId to identify the search
	// when using the feedback API.
	QueryId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpQueryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpQuery{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpQuery{}, middleware.After)
}

func newServiceMetadataMiddleware_opQuery(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "Query",
	}
}
