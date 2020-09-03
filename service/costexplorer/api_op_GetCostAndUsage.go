// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves cost and usage metrics for your account. You can specify which cost
// and usage-related metric, such as BlendedCosts or UsageQuantity, that you want
// the request to return. You can also filter and group your data by various
// dimensions, such as SERVICE or AZ, in a specific time range. For a complete list
// of valid dimensions, see the GetDimensionValues
// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_GetDimensionValues.html)
// operation. Master accounts in an organization in AWS Organizations have access
// to all member accounts.
func (c *Client) GetCostAndUsage(ctx context.Context, params *GetCostAndUsageInput, optFns ...func(*Options)) (*GetCostAndUsageOutput, error) {
	stack := middleware.NewStack("GetCostAndUsage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetCostAndUsageMiddlewares(stack)
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
	addOpGetCostAndUsageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCostAndUsage(options.Region), middleware.Before)

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
			OperationName: "GetCostAndUsage",
			Err:           err,
		}
	}
	out := result.(*GetCostAndUsageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCostAndUsageInput struct {
	// Sets the start and end dates for retrieving AWS costs. The start date is
	// inclusive, but the end date is exclusive. For example, if start is 2017-01-01
	// and end is 2017-05-01, then the cost and usage data is retrieved from 2017-01-01
	// up to and including 2017-04-30 but not including 2017-05-01.
	TimePeriod *types.DateInterval
	// The token to retrieve the next set of results. AWS provides the token when the
	// response from a previous call has more results than the maximum page size.
	NextPageToken *string
	// Filters AWS costs by different dimensions. For example, you can specify SERVICE
	// and LINKED_ACCOUNT and get the costs that are associated with that account's
	// usage of that service. You can nest Expression objects to define any combination
	// of dimension filters. For more information, see Expression
	// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html).
	Filter *types.Expression
	// Which metrics are returned in the query. For more information about blended and
	// unblended rates, see Why does the "blended" annotation appear on some line items
	// in my bill?
	// (http://aws.amazon.com/premiumsupport/knowledge-center/blended-rates-intro/).
	// Valid values are AmortizedCost, BlendedCost, NetAmortizedCost, NetUnblendedCost,
	// NormalizedUsageAmount, UnblendedCost, and UsageQuantity. If you return the
	// UsageQuantity metric, the service aggregates all usage numbers without taking
	// into account the units. For example, if you aggregate usageQuantity across all
	// of Amazon EC2, the results aren't meaningful because Amazon EC2 compute hours
	// and data transfer are measured in different units (for example, hours vs. GB).
	// To get more meaningful UsageQuantity metrics, filter by UsageType or
	// UsageTypeGroups. Metrics is required for GetCostAndUsage requests.
	Metrics []*string
	// Sets the AWS cost granularity to MONTHLY or DAILY, or HOURLY. If Granularity
	// isn't set, the response object doesn't include the Granularity, either MONTHLY
	// or DAILY, or HOURLY.
	Granularity types.Granularity
	// You can group AWS costs using up to two different groups, either dimensions, tag
	// keys, or both. When you group by tag key, you get all tag values, including
	// empty strings. Valid values are AZ, INSTANCE_TYPE, LEGAL_ENTITY_NAME,
	// LINKED_ACCOUNT, OPERATION, PLATFORM, PURCHASE_TYPE, SERVICE, TAGS, TENANCY,
	// RECORD_TYPE, and USAGE_TYPE.
	GroupBy []*types.GroupDefinition
}

type GetCostAndUsageOutput struct {
	// The groups that are specified by the Filter or GroupBy parameters in the
	// request.
	GroupDefinitions []*types.GroupDefinition
	// The token for the next set of retrievable results. AWS provides the token when
	// the response from a previous call has more results than the maximum page size.
	NextPageToken *string
	// The time period that is covered by the results in the response.
	ResultsByTime []*types.ResultByTime

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetCostAndUsageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetCostAndUsage{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCostAndUsage{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetCostAndUsage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "GetCostAndUsage",
	}
}
