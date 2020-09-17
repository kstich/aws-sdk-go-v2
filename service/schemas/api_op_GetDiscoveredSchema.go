// Code generated by smithy-go-codegen DO NOT EDIT.

package schemas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/schemas/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Get the discovered schema that was generated based on sampled events.
func (c *Client) GetDiscoveredSchema(ctx context.Context, params *GetDiscoveredSchemaInput, optFns ...func(*Options)) (*GetDiscoveredSchemaOutput, error) {
	stack := middleware.NewStack("GetDiscoveredSchema", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetDiscoveredSchemaMiddlewares(stack)
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
	addOpGetDiscoveredSchemaValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDiscoveredSchema(options.Region), middleware.Before)
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
			OperationName: "GetDiscoveredSchema",
			Err:           err,
		}
	}
	out := result.(*GetDiscoveredSchemaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDiscoveredSchemaInput struct {
	// An array of strings where each string is a JSON event. These are the events that
	// were used to generate the schema. The array includes a single type of event and
	// has a maximum size of 10 events.
	Events []*string
	// The type of event.
	Type types.Type
}

type GetDiscoveredSchemaOutput struct {
	// The source of the schema definition.
	Content *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetDiscoveredSchemaMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetDiscoveredSchema{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDiscoveredSchema{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDiscoveredSchema(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "schemas",
		OperationName: "GetDiscoveredSchema",
	}
}
