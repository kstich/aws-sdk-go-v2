// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes pending authorization requests for a specified aggregator account in a
// specified region.
func (c *Client) DeletePendingAggregationRequest(ctx context.Context, params *DeletePendingAggregationRequestInput, optFns ...func(*Options)) (*DeletePendingAggregationRequestOutput, error) {
	stack := middleware.NewStack("DeletePendingAggregationRequest", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeletePendingAggregationRequestMiddlewares(stack)
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
	addOpDeletePendingAggregationRequestValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePendingAggregationRequest(options.Region), middleware.Before)
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
			OperationName: "DeletePendingAggregationRequest",
			Err:           err,
		}
	}
	out := result.(*DeletePendingAggregationRequestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePendingAggregationRequestInput struct {
	// The 12-digit account ID of the account requesting to aggregate data.
	RequesterAccountId *string
	// The region requesting to aggregate data.
	RequesterAwsRegion *string
}

type DeletePendingAggregationRequestOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeletePendingAggregationRequestMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeletePendingAggregationRequest{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeletePendingAggregationRequest{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeletePendingAggregationRequest(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DeletePendingAggregationRequest",
	}
}
