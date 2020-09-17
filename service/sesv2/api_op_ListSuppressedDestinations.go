// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieves a list of email addresses that are on the suppression list for your
// account.
func (c *Client) ListSuppressedDestinations(ctx context.Context, params *ListSuppressedDestinationsInput, optFns ...func(*Options)) (*ListSuppressedDestinationsOutput, error) {
	stack := middleware.NewStack("ListSuppressedDestinations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListSuppressedDestinationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListSuppressedDestinations(options.Region), middleware.Before)
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
			OperationName: "ListSuppressedDestinations",
			Err:           err,
		}
	}
	out := result.(*ListSuppressedDestinationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to obtain a list of email destinations that are on the suppression
// list for your account.
type ListSuppressedDestinationsInput struct {
	// Used to filter the list of suppressed email destinations so that it only
	// includes addresses that were added to the list before a specific date. The date
	// that you specify should be in Unix time format.
	EndDate *time.Time
	// A token returned from a previous call to ListSuppressedDestinations to indicate
	// the position in the list of suppressed email addresses.
	NextToken *string
	// The number of results to show in a single call to ListSuppressedDestinations. If
	// the number of results is larger than the number you specified in this parameter,
	// then the response includes a NextToken element, which you can use to obtain
	// additional results.
	PageSize *int32
	// The factors that caused the email address to be added to .
	Reasons []types.SuppressionListReason
	// Used to filter the list of suppressed email destinations so that it only
	// includes addresses that were added to the list after a specific date. The date
	// that you specify should be in Unix time format.
	StartDate *time.Time
}

// A list of suppressed email addresses.
type ListSuppressedDestinationsOutput struct {
	// A list of summaries, each containing a summary for a suppressed email
	// destination.
	SuppressedDestinationSummaries []*types.SuppressedDestinationSummary
	// A token that indicates that there are additional email addresses on the
	// suppression list for your account. To view additional suppressed addresses,
	// issue another request to ListSuppressedDestinations, and pass this token in the
	// NextToken parameter.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListSuppressedDestinationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListSuppressedDestinations{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListSuppressedDestinations{}, middleware.After)
}

func newServiceMetadataMiddleware_opListSuppressedDestinations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListSuppressedDestinations",
	}
}
