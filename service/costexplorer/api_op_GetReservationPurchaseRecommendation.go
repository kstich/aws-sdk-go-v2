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

// Gets recommendations for which reservations to purchase. These recommendations
// could help you reduce your costs. Reservations provide a discounted hourly rate
// (up to 75%) compared to On-Demand pricing. AWS generates your recommendations by
// identifying your On-Demand usage during a specific time period and collecting
// your usage into categories that are eligible for a reservation. After AWS has
// these categories, it simulates every combination of reservations in each
// category of usage to identify the best number of each type of RI to purchase to
// maximize your estimated savings. For example, AWS automatically aggregates your
// Amazon EC2 Linux, shared tenancy, and c4 family usage in the US West (Oregon)
// Region and recommends that you buy size-flexible regional reservations to apply
// to the c4 family usage. AWS recommends the smallest size instance in an instance
// family. This makes it easier to purchase a size-flexible RI. AWS also shows the
// equal number of normalized units so that you can purchase any instance size that
// you want. For this example, your RI recommendation would be for c4.large because
// that is the smallest size instance in the c4 instance family.
func (c *Client) GetReservationPurchaseRecommendation(ctx context.Context, params *GetReservationPurchaseRecommendationInput, optFns ...func(*Options)) (*GetReservationPurchaseRecommendationOutput, error) {
	stack := middleware.NewStack("GetReservationPurchaseRecommendation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetReservationPurchaseRecommendationMiddlewares(stack)
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
	addOpGetReservationPurchaseRecommendationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetReservationPurchaseRecommendation(options.Region), middleware.Before)
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
			OperationName: "GetReservationPurchaseRecommendation",
			Err:           err,
		}
	}
	out := result.(*GetReservationPurchaseRecommendationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetReservationPurchaseRecommendationInput struct {
	// The number of recommendations that you want returned in a single response
	// object.
	PageSize *int32
	// The account ID that is associated with the recommendation.
	AccountId *string
	// The specific service that you want recommendations for.
	Service *string
	// The number of previous days that you want AWS to consider when it calculates
	// your recommendations.
	LookbackPeriodInDays types.LookbackPeriodInDays
	// The reservation term that you want recommendations for.
	TermInYears types.TermInYears
	// The pagination token that indicates the next set of results that you want to
	// retrieve.
	NextPageToken *string
	// The hardware specifications for the service instances that you want
	// recommendations for, such as standard or convertible Amazon EC2 instances.
	ServiceSpecification *types.ServiceSpecification
	// The reservation purchase option that you want recommendations for.
	PaymentOption types.PaymentOption
	// The account scope that you want your recommendations for. Amazon Web Services
	// calculates recommendations including the payer account and linked accounts if
	// the value is set to PAYER. If the value is LINKED, recommendations are
	// calculated for individual linked accounts only.
	AccountScope types.AccountScope
}

type GetReservationPurchaseRecommendationOutput struct {
	// Recommendations for reservations to purchase.
	Recommendations []*types.ReservationPurchaseRecommendation
	// Information about this specific recommendation call, such as the time stamp for
	// when Cost Explorer generated this recommendation.
	Metadata *types.ReservationPurchaseRecommendationMetadata
	// The pagination token for the next set of retrievable results.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetReservationPurchaseRecommendationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetReservationPurchaseRecommendation{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetReservationPurchaseRecommendation{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetReservationPurchaseRecommendation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "GetReservationPurchaseRecommendation",
	}
}
