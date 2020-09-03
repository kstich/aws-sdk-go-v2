// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Searches a set of tables based on properties in the table metadata as well as on
// the parent database. You can search against text or filter conditions. You can
// only get tables that you have access to based on the security policies defined
// in Lake Formation. You need at least a read-only access to the table for it to
// be returned. If you do not have access to all the columns in the table, these
// columns will not be searched against when returning the list of tables back to
// you. If you have access to the columns but not the data in the columns, those
// columns and the associated metadata for those columns will be included in the
// search.
func (c *Client) SearchTables(ctx context.Context, params *SearchTablesInput, optFns ...func(*Options)) (*SearchTablesOutput, error) {
	stack := middleware.NewStack("SearchTables", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpSearchTablesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opSearchTables(options.Region), middleware.Before)

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
			OperationName: "SearchTables",
			Err:           err,
		}
	}
	out := result.(*SearchTablesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchTablesInput struct {
	// A list of criteria for sorting the results by a field name, in an ascending or
	// descending order.
	SortCriteria []*types.SortCriterion
	// A list of key-value pairs, and a comparator used to filter the search results.
	// Returns all entities matching the predicate.  <p>The <code>Comparator</code>
	// member of the <code>PropertyPredicate</code> struct is used only for time
	// fields, and can be omitted for other field types. Also, when comparing string
	// values, such as when <code>Key=Name</code>, a fuzzy match algorithm is used. The
	// <code>Key</code> field (for example, the value of the <code>Name</code> field)
	// is split on certain punctuation characters, for example, -, :, #, etc. into
	// tokens. Then each token is exact-match compared with the <code>Value</code>
	// member of <code>PropertyPredicate</code>. For example, if <code>Key=Name</code>
	// and <code>Value=link</code>, tables named <code>customer-link</code> and
	// <code>xx-link-yy</code> are returned, but <code>xxlinkyy</code> is not
	// returned.</p>
	Filters []*types.PropertyPredicate
	// A unique identifier, consisting of  account_id .
	CatalogId *string
	// Allows you to specify that you want to search the tables shared with your
	// account. The allowable values are FOREIGN or ALL.  <ul> <li> <p>If set to
	// <code>FOREIGN</code>, will search the tables shared with your account. </p>
	// </li> <li> <p>If set to <code>ALL</code>, will search the tables shared with
	// your account, as well as the tables in yor local account. </p> </li> </ul>
	ResourceShareType types.ResourceShareType
	// The maximum number of tables to return in a single response.
	MaxResults *int32
	// A continuation token, included if this is a continuation call.
	NextToken *string
	// A string used for a text search. Specifying a value in quotes filters based on
	// an exact match to the value.
	SearchText *string
}

type SearchTablesOutput struct {
	// A continuation token, present if the current list segment is not the last.
	NextToken *string
	// A list of the requested Table objects. The SearchTables response returns only
	// the tables that you have access to.
	TableList []*types.Table

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpSearchTablesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpSearchTables{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchTables{}, middleware.After)
}

func newServiceMetadataMiddleware_opSearchTables(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "SearchTables",
	}
}
