// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists Amazon GuardDuty usage statistics over the last 30 days for the specified
// detector ID. For newly enabled detectors or data sources the cost returned will
// include only the usage so far under 30 days, this may differ from the cost
// metrics in the console, which projects usage over 30 days to provide a monthly
// cost estimate. For more information see Understanding How Usage Costs are
// Calculated
// (https://docs.aws.amazon.com/guardduty/latest/ug/monitoring_costs.html#usage-calculations).
func (c *Client) GetUsageStatistics(ctx context.Context, params *GetUsageStatisticsInput, optFns ...func(*Options)) (*GetUsageStatisticsOutput, error) {
	stack := middleware.NewStack("GetUsageStatistics", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetUsageStatisticsMiddlewares(stack)
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
	addOpGetUsageStatisticsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetUsageStatistics(options.Region), middleware.Before)

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
			OperationName: "GetUsageStatistics",
			Err:           err,
		}
	}
	out := result.(*GetUsageStatisticsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUsageStatisticsInput struct {
	// Represents the criteria used for querying usage.
	UsageCriteria *types.UsageCriteria
	// The type of usage statistics to retrieve.
	UsageStatisticType types.UsageStatisticType
	// The ID of the detector that specifies the GuardDuty service whose usage
	// statistics you want to retrieve.
	DetectorId *string
	// The currency unit you would like to view your usage statistics in. Current valid
	// values are USD.
	Unit *string
	// The maximum number of results to return in the response.
	MaxResults *int32
	// A token to use for paginating results that are returned in the response. Set the
	// value of this parameter to null for the first request to a list action. For
	// subsequent calls, use the NextToken value returned from the previous request to
	// continue listing results after the first page.
	NextToken *string
}

type GetUsageStatisticsOutput struct {
	// The pagination parameter to be used on the next list operation to retrieve more
	// items.
	NextToken *string
	// The usage statistics object. If a UsageStatisticType was provided, the objects
	// representing other types will be null.
	UsageStatistics *types.UsageStatistics

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetUsageStatisticsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetUsageStatistics{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUsageStatistics{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetUsageStatistics(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "GetUsageStatistics",
	}
}
