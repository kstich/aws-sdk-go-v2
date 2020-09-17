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

// Returns the number of compliant and noncompliant rules for one or more accounts
// and regions in an aggregator. The results can return an empty result page, but
// if you have a nextToken, the results are displayed on the next page.
func (c *Client) GetAggregateConfigRuleComplianceSummary(ctx context.Context, params *GetAggregateConfigRuleComplianceSummaryInput, optFns ...func(*Options)) (*GetAggregateConfigRuleComplianceSummaryOutput, error) {
	stack := middleware.NewStack("GetAggregateConfigRuleComplianceSummary", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetAggregateConfigRuleComplianceSummaryMiddlewares(stack)
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
	addOpGetAggregateConfigRuleComplianceSummaryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAggregateConfigRuleComplianceSummary(options.Region), middleware.Before)
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
			OperationName: "GetAggregateConfigRuleComplianceSummary",
			Err:           err,
		}
	}
	out := result.(*GetAggregateConfigRuleComplianceSummaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAggregateConfigRuleComplianceSummaryInput struct {
	// The name of the configuration aggregator.
	ConfigurationAggregatorName *string
	// Groups the result based on ACCOUNT_ID or AWS_REGION.
	GroupByKey types.ConfigRuleComplianceSummaryGroupKey
	// Filters the results based on the ConfigRuleComplianceSummaryFilters object.
	Filters *types.ConfigRuleComplianceSummaryFilters
	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
	// The maximum number of evaluation results returned on each page. The default is
	// 1000. You cannot specify a number greater than 1000. If you specify 0, AWS
	// Config uses the default.
	Limit *int32
}

type GetAggregateConfigRuleComplianceSummaryOutput struct {
	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
	// Returns a list of AggregateComplianceCounts object.
	AggregateComplianceCounts []*types.AggregateComplianceCount
	// Groups the result based on ACCOUNT_ID or AWS_REGION.
	GroupByKey *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetAggregateConfigRuleComplianceSummaryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetAggregateConfigRuleComplianceSummary{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAggregateConfigRuleComplianceSummary{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetAggregateConfigRuleComplianceSummary(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "GetAggregateConfigRuleComplianceSummary",
	}
}
