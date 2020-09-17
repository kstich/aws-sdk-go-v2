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

// Retrieves the reservation coverage for your account. This enables you to see how
// much of your Amazon Elastic Compute Cloud, Amazon ElastiCache, Amazon Relational
// Database Service, or Amazon Redshift usage is covered by a reservation. An
// organization's master account can see the coverage of the associated member
// accounts. This supports dimensions, Cost Categories, and nested expressions. For
// any time period, you can filter data about reservation usage by the following
// dimensions:
//
//     * AZ
//
//     * CACHE_ENGINE
//
//     * DATABASE_ENGINE
//
//     *
// DEPLOYMENT_OPTION
//
//     * INSTANCE_TYPE
//
//     * LINKED_ACCOUNT
//
//     *
// OPERATING_SYSTEM
//
//     * PLATFORM
//
//     * REGION
//
//     * SERVICE
//
//     * TAG
//
//     *
// TENANCY
//
// To determine valid values for a dimension, use the GetDimensionValues
// operation.
func (c *Client) GetReservationCoverage(ctx context.Context, params *GetReservationCoverageInput, optFns ...func(*Options)) (*GetReservationCoverageOutput, error) {
	stack := middleware.NewStack("GetReservationCoverage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetReservationCoverageMiddlewares(stack)
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
	addOpGetReservationCoverageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetReservationCoverage(options.Region), middleware.Before)
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
			OperationName: "GetReservationCoverage",
			Err:           err,
		}
	}
	out := result.(*GetReservationCoverageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// You can use the following request parameters to query for how much of your
// instance usage a reservation covered.
type GetReservationCoverageInput struct {
	// Filters utilization data by dimensions. You can filter by the following
	// dimensions:
	//
	//     * AZ
	//
	//     * CACHE_ENGINE
	//
	//     * DATABASE_ENGINE
	//
	//     *
	// DEPLOYMENT_OPTION
	//
	//     * INSTANCE_TYPE
	//
	//     * LINKED_ACCOUNT
	//
	//     *
	// OPERATING_SYSTEM
	//
	//     * PLATFORM
	//
	//     * REGION
	//
	//     * SERVICE
	//
	//     * TAG
	//
	//     *
	// TENANCY
	//
	// GetReservationCoverage uses the same Expression
	// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html)
	// object as the other operations, but only AND is supported among each dimension.
	// You can nest only one level deep. If there are multiple values for a dimension,
	// they are OR'd together. If you don't provide a SERVICE filter, Cost Explorer
	// defaults to EC2. Cost category is also supported.
	Filter *types.Expression
	// The granularity of the AWS cost data for the reservation. Valid values are
	// MONTHLY and DAILY. If GroupBy is set, Granularity can't be set. If Granularity
	// isn't set, the response object doesn't include Granularity, either MONTHLY or
	// DAILY. The GetReservationCoverage operation supports only DAILY and MONTHLY
	// granularities.
	Granularity types.Granularity
	// The start and end dates of the period that you want to retrieve data about
	// reservation coverage for. You can retrieve data for a maximum of 13 months: the
	// last 12 months and the current month. The start date is inclusive, but the end
	// date is exclusive. For example, if start is 2017-01-01 and end is 2017-05-01,
	// then the cost and usage data is retrieved from 2017-01-01 up to and including
	// 2017-04-30 but not including 2017-05-01.
	TimePeriod *types.DateInterval
	// The measurement that you want your reservation coverage reported in. Valid
	// values are Hour, Unit, and Cost. You can use multiple values in a request.
	Metrics []*string
	// The token to retrieve the next set of results. AWS provides the token when the
	// response from a previous call has more results than the maximum page size.
	NextPageToken *string
	// You can group the data by the following attributes:
	//
	//     * AZ
	//
	//     *
	// CACHE_ENGINE
	//
	//     * DATABASE_ENGINE
	//
	//     * DEPLOYMENT_OPTION
	//
	//     *
	// INSTANCE_TYPE
	//
	//     * LINKED_ACCOUNT
	//
	//     * OPERATING_SYSTEM
	//
	//     * PLATFORM
	//
	//
	// * REGION
	//
	//     * TENANCY
	GroupBy []*types.GroupDefinition
}

type GetReservationCoverageOutput struct {
	// The total amount of instance usage that a reservation covered.
	Total *types.Coverage
	// The token for the next set of retrievable results. AWS provides the token when
	// the response from a previous call has more results than the maximum page size.
	NextPageToken *string
	// The amount of time that your reservations covered.
	CoveragesByTime []*types.CoverageByTime

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetReservationCoverageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetReservationCoverage{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetReservationCoverage{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetReservationCoverage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "GetReservationCoverage",
	}
}
