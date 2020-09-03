// Code generated by smithy-go-codegen DO NOT EDIT.

package iotdataplane

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the shadows for the specified thing.
func (c *Client) ListNamedShadowsForThing(ctx context.Context, params *ListNamedShadowsForThingInput, optFns ...func(*Options)) (*ListNamedShadowsForThingOutput, error) {
	stack := middleware.NewStack("ListNamedShadowsForThing", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListNamedShadowsForThingMiddlewares(stack)
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
	addOpListNamedShadowsForThingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListNamedShadowsForThing(options.Region), middleware.Before)

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
			OperationName: "ListNamedShadowsForThing",
			Err:           err,
		}
	}
	out := result.(*ListNamedShadowsForThingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNamedShadowsForThingInput struct {
	// The token to retrieve the next set of results.
	NextToken *string
	// The name of the thing.
	ThingName *string
	// The result page size.
	PageSize *int32
}

type ListNamedShadowsForThingOutput struct {
	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string
	// The list of shadows for the specified thing.
	Results []*string
	// The Epoch date and time the response was generated by AWS IoT.
	Timestamp *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListNamedShadowsForThingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListNamedShadowsForThing{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListNamedShadowsForThing{}, middleware.After)
}

func newServiceMetadataMiddleware_opListNamedShadowsForThing(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotdata",
		OperationName: "ListNamedShadowsForThing",
	}
}