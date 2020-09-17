// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Purchase a reservation with configurations that match those of your Dedicated
// Host. You must have active Dedicated Hosts in your account before you purchase a
// reservation. This action results in the specified reservation being purchased
// and charged to your account.
func (c *Client) PurchaseHostReservation(ctx context.Context, params *PurchaseHostReservationInput, optFns ...func(*Options)) (*PurchaseHostReservationOutput, error) {
	stack := middleware.NewStack("PurchaseHostReservation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpPurchaseHostReservationMiddlewares(stack)
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
	addOpPurchaseHostReservationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPurchaseHostReservation(options.Region), middleware.Before)
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
			OperationName: "PurchaseHostReservation",
			Err:           err,
		}
	}
	out := result.(*PurchaseHostReservationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PurchaseHostReservationInput struct {
	// The tags to apply to the Dedicated Host Reservation during purchase.
	TagSpecifications []*types.TagSpecification
	// The specified limit is checked against the total upfront cost of the reservation
	// (calculated as the offering's upfront cost multiplied by the host count). If the
	// total upfront cost is greater than the specified price limit, the request fails.
	// This is used to ensure that the purchase does not exceed the expected upfront
	// cost of the purchase. At this time, the only supported currency is USD. For
	// example, to indicate a limit price of USD 100, specify 100.00.
	LimitPrice *string
	// The currency in which the totalUpfrontPrice, LimitPrice, and totalHourlyPrice
	// amounts are specified. At this time, the only supported currency is USD.
	CurrencyCode types.CurrencyCodeValues
	// The IDs of the Dedicated Hosts with which the reservation will be associated.
	HostIdSet []*string
	// The ID of the offering.
	OfferingId *string
	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see How to Ensure Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string
}

type PurchaseHostReservationOutput struct {
	// The currency in which the totalUpfrontPrice and totalHourlyPrice amounts are
	// specified. At this time, the only supported currency is USD.
	CurrencyCode types.CurrencyCodeValues
	// Describes the details of the purchase.
	Purchase []*types.Purchase
	// The total amount charged to your account when you purchase the reservation.
	TotalUpfrontPrice *string
	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see How to Ensure Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string
	// The total hourly price of the reservation calculated per hour.
	TotalHourlyPrice *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpPurchaseHostReservationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpPurchaseHostReservation{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpPurchaseHostReservation{}, middleware.After)
}

func newServiceMetadataMiddleware_opPurchaseHostReservation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "PurchaseHostReservation",
	}
}
