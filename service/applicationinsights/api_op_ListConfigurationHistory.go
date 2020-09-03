// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationinsights

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationinsights/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Lists the INFO, WARN, and ERROR events for periodic configuration updates
// performed by Application Insights. Examples of events represented are:
//
//     *
// INFO: creating a new alarm or updating an alarm threshold.
//
//     * WARN: alarm
// not created due to insufficient data points used to predict thresholds.
//
//     *
// ERROR: alarm not created due to permission errors or exceeding quotas.
func (c *Client) ListConfigurationHistory(ctx context.Context, params *ListConfigurationHistoryInput, optFns ...func(*Options)) (*ListConfigurationHistoryOutput, error) {
	stack := middleware.NewStack("ListConfigurationHistory", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListConfigurationHistoryMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListConfigurationHistory(options.Region), middleware.Before)

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
			OperationName: "ListConfigurationHistory",
			Err:           err,
		}
	}
	out := result.(*ListConfigurationHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConfigurationHistoryInput struct {
	// The maximum number of results returned by ListConfigurationHistory in paginated
	// output. When this parameter is used, ListConfigurationHistory returns only
	// MaxResults in a single page along with a NextToken response element. The
	// remaining results of the initial request can be seen by sending another
	// ListConfigurationHistory request with the returned NextToken value. If this
	// parameter is not used, then ListConfigurationHistory returns all results.
	MaxResults *int32
	// Resource group to which the application belongs.
	ResourceGroupName *string
	// The NextToken value returned from a previous paginated ListConfigurationHistory
	// request where MaxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the NextToken value. This value is null when there are no more results
	// to return.
	NextToken *string
	// The status of the configuration update event. Possible values include INFO,
	// WARN, and ERROR.
	EventStatus types.ConfigurationEventStatus
	// The end time of the event.
	EndTime *time.Time
	// The start time of the event.
	StartTime *time.Time
}

type ListConfigurationHistoryOutput struct {
	// The list of configuration events and their corresponding details.
	EventList []*types.ConfigurationEvent
	// The NextToken value to include in a future ListConfigurationHistory request.
	// When the results of a ListConfigurationHistory request exceed MaxResults, this
	// value can be used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListConfigurationHistoryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListConfigurationHistory{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListConfigurationHistory{}, middleware.After)
}

func newServiceMetadataMiddleware_opListConfigurationHistory(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "applicationinsights",
		OperationName: "ListConfigurationHistory",
	}
}
