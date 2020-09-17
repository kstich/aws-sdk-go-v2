// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is used by SaaS partners to write events to a customer's partner event bus.
// AWS customers do not use this operation.
func (c *Client) PutPartnerEvents(ctx context.Context, params *PutPartnerEventsInput, optFns ...func(*Options)) (*PutPartnerEventsOutput, error) {
	stack := middleware.NewStack("PutPartnerEvents", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutPartnerEventsMiddlewares(stack)
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
	addOpPutPartnerEventsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutPartnerEvents(options.Region), middleware.Before)
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
			OperationName: "PutPartnerEvents",
			Err:           err,
		}
	}
	out := result.(*PutPartnerEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutPartnerEventsInput struct {
	// The list of events to write to the event bus.
	Entries []*types.PutPartnerEventsRequestEntry
}

type PutPartnerEventsOutput struct {
	// The number of events from this operation that could not be written to the
	// partner event bus.
	FailedEntryCount *int32
	// The list of events from this operation that were successfully written to the
	// partner event bus.
	Entries []*types.PutPartnerEventsResultEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutPartnerEventsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutPartnerEvents{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutPartnerEvents{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutPartnerEvents(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "PutPartnerEvents",
	}
}
