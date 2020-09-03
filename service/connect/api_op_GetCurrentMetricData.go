// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Gets the real-time metric data from the specified Amazon Connect instance. For
// more information, see Real-time Metrics Reports
// (https://docs.aws.amazon.com/connect/latest/adminguide/real-time-metrics-reports.html)
// in the Amazon Connect Administrator Guide.
func (c *Client) GetCurrentMetricData(ctx context.Context, params *GetCurrentMetricDataInput, optFns ...func(*Options)) (*GetCurrentMetricDataOutput, error) {
	stack := middleware.NewStack("GetCurrentMetricData", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetCurrentMetricDataMiddlewares(stack)
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
	addOpGetCurrentMetricDataValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCurrentMetricData(options.Region), middleware.Before)

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
			OperationName: "GetCurrentMetricData",
			Err:           err,
		}
	}
	out := result.(*GetCurrentMetricDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCurrentMetricDataInput struct {
	// The queues, up to 100, or channels, to use to filter the metrics returned.
	// Metric data is retrieved only for the resources associated with the queues or
	// channels included in the filter. You can include both queue IDs and queue ARNs
	// in the same request. The only supported channel is VOICE.
	Filters *types.Filters
	// The grouping applied to the metrics returned. For example, when grouped by
	// QUEUE, the metrics returned apply to each queue rather than aggregated for all
	// queues. If you group by CHANNEL, you should include a Channels filter. The only
	// supported channel is VOICE. If no Grouping is included in the request, a summary
	// of metrics is returned.
	Groupings []types.Grouping
	// The maximimum number of results to return per page.
	MaxResults *int32
	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results. The token
	// expires after 5 minutes from the time it is created. Subsequent requests that
	// use the token must use the same request parameters as the request that generated
	// the token.
	NextToken *string
	// The metrics to retrieve. Specify the name and unit for each metric. The
	// following metrics are available. For a description of each metric, see Real-time
	// Metrics Definitions
	// (https://docs.aws.amazon.com/connect/latest/adminguide/real-time-metrics-definitions.html)
	// in the Amazon Connect Administrator Guide. AGENTS_AFTER_CONTACT_WORK Unit: COUNT
	// AGENTS_AVAILABLE Unit: COUNT AGENTS_ERROR Unit: COUNT AGENTS_NON_PRODUCTIVE
	// Unit: COUNT AGENTS_ON_CALL Unit: COUNT AGENTS_ON_CONTACT Unit: COUNT
	// AGENTS_ONLINE Unit: COUNT AGENTS_STAFFED Unit: COUNT CONTACTS_IN_QUEUE Unit:
	// COUNT CONTACTS_SCHEDULED Unit: COUNT OLDEST_CONTACT_AGE Unit: SECONDS
	// SLOTS_ACTIVE Unit: COUNT SLOTS_AVAILABLE Unit: COUNT
	CurrentMetrics []*types.CurrentMetric
	// The identifier of the Amazon Connect instance.
	InstanceId *string
}

type GetCurrentMetricDataOutput struct {
	// Information about the real-time metrics.
	MetricResults []*types.CurrentMetricResult
	// If there are additional results, this is the token for the next set of results.
	// The token expires after 5 minutes from the time it is created. Subsequent
	// requests that use the token must use the same request parameters as the request
	// that generated the token.
	NextToken *string
	// The time at which the metrics were retrieved and cached for pagination.
	DataSnapshotTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetCurrentMetricDataMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetCurrentMetricData{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCurrentMetricData{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetCurrentMetricData(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "GetCurrentMetricData",
	}
}
