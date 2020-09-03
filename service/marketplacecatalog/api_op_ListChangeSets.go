// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplacecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/marketplacecatalog/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the list of change sets owned by the account being used to make the
// call. You can filter this list by providing any combination of entityId,
// ChangeSetName, and status. If you provide more than one filter, the API
// operation applies a logical AND between the filters.  <p>You can describe a
// change during the 60-day request history retention period for API calls.</p>
func (c *Client) ListChangeSets(ctx context.Context, params *ListChangeSetsInput, optFns ...func(*Options)) (*ListChangeSetsOutput, error) {
	stack := middleware.NewStack("ListChangeSets", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListChangeSetsMiddlewares(stack)
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
	addOpListChangeSetsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListChangeSets(options.Region), middleware.Before)

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
			OperationName: "ListChangeSets",
			Err:           err,
		}
	}
	out := result.(*ListChangeSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListChangeSetsInput struct {
	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string
	// An object that contains two attributes, SortBy and SortOrder.
	Sort *types.Sort
	// An array of filter objects.
	FilterList []*types.Filter
	// The catalog related to the request. Fixed value: AWSMarketplace
	Catalog *string
	// The maximum number of results returned by a single call. This value must be
	// provided in the next call to retrieve the next set of results. By default, this
	// value is 20.
	MaxResults *int32
}

type ListChangeSetsOutput struct {
	// The value of the next token, if it exists. Null if there are no more results.
	NextToken *string
	// Array of ChangeSetSummaryListItem objects.
	ChangeSetSummaryList []*types.ChangeSetSummaryListItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListChangeSetsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListChangeSets{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListChangeSets{}, middleware.After)
}

func newServiceMetadataMiddleware_opListChangeSets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aws-marketplace",
		OperationName: "ListChangeSets",
	}
}
