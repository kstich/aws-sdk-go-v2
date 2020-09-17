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

// Cancels the specified Reserved Instance listing in the Reserved Instance
// Marketplace. For more information, see Reserved Instance Marketplace
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ri-market-general.html) in
// the Amazon Elastic Compute Cloud User Guide.
func (c *Client) CancelReservedInstancesListing(ctx context.Context, params *CancelReservedInstancesListingInput, optFns ...func(*Options)) (*CancelReservedInstancesListingOutput, error) {
	stack := middleware.NewStack("CancelReservedInstancesListing", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpCancelReservedInstancesListingMiddlewares(stack)
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
	addOpCancelReservedInstancesListingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCancelReservedInstancesListing(options.Region), middleware.Before)
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
			OperationName: "CancelReservedInstancesListing",
			Err:           err,
		}
	}
	out := result.(*CancelReservedInstancesListingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CancelReservedInstancesListing.
type CancelReservedInstancesListingInput struct {
	// The ID of the Reserved Instance listing.
	ReservedInstancesListingId *string
}

// Contains the output of CancelReservedInstancesListing.
type CancelReservedInstancesListingOutput struct {
	// The Reserved Instance listing.
	ReservedInstancesListings []*types.ReservedInstancesListing

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpCancelReservedInstancesListingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpCancelReservedInstancesListing{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpCancelReservedInstancesListing{}, middleware.After)
}

func newServiceMetadataMiddleware_opCancelReservedInstancesListing(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CancelReservedInstancesListing",
	}
}
