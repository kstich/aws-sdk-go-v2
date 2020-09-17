// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new endpoint for an application or updates the settings and attributes
// of an existing endpoint for an application. You can also use this operation to
// define custom attributes for an endpoint. If an update includes one or more
// values for a custom attribute, Amazon Pinpoint replaces (overwrites) any
// existing values with the new values.
func (c *Client) UpdateEndpoint(ctx context.Context, params *UpdateEndpointInput, optFns ...func(*Options)) (*UpdateEndpointOutput, error) {
	stack := middleware.NewStack("UpdateEndpoint", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateEndpointMiddlewares(stack)
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
	addOpUpdateEndpointValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEndpoint(options.Region), middleware.Before)
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
			OperationName: "UpdateEndpoint",
			Err:           err,
		}
	}
	out := result.(*UpdateEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEndpointInput struct {
	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	ApplicationId *string
	// Specifies the channel type and other settings for an endpoint.
	EndpointRequest *types.EndpointRequest
	// The unique identifier for the endpoint.
	EndpointId *string
}

type UpdateEndpointOutput struct {
	// Provides information about an API request or response.
	MessageBody *types.MessageBody

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateEndpointMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateEndpoint{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateEndpoint{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateEndpoint(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "UpdateEndpoint",
	}
}
