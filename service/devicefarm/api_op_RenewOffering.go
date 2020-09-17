// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Explicitly sets the quantity of devices to renew for an offering, starting from
// the effectiveDate of the next period. The API returns a NotEligible error if the
// user is not permitted to invoke the operation. If you must be able to invoke
// this operation, contact aws-devicefarm-support@amazon.com
// (mailto:aws-devicefarm-support@amazon.com).
func (c *Client) RenewOffering(ctx context.Context, params *RenewOfferingInput, optFns ...func(*Options)) (*RenewOfferingOutput, error) {
	stack := middleware.NewStack("RenewOffering", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRenewOfferingMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opRenewOffering(options.Region), middleware.Before)
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
			OperationName: "RenewOffering",
			Err:           err,
		}
	}
	out := result.(*RenewOfferingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request that represents an offering renewal.
type RenewOfferingInput struct {
	// The ID of a request to renew an offering.
	OfferingId *string
	// The quantity requested in an offering renewal.
	Quantity *int32
}

// The result of a renewal offering.
type RenewOfferingOutput struct {
	// Represents the status of the offering transaction for the renewal.
	OfferingTransaction *types.OfferingTransaction

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRenewOfferingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRenewOffering{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRenewOffering{}, middleware.After)
}

func newServiceMetadataMiddleware_opRenewOffering(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "RenewOffering",
	}
}
