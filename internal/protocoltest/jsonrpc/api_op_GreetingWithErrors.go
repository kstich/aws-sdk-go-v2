// Code generated by smithy-go-codegen DO NOT EDIT.

package jsonrpc

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation has three possible return values:
//
//     * A successful response in
// the form of GreetingWithErrorsOutput
//
//     * An InvalidGreeting error.
//
//     * A
// ComplexError error.
//
// Implementations must be able to successfully take a
// response and properly deserialize successful and error responses.
func (c *Client) GreetingWithErrors(ctx context.Context, params *GreetingWithErrorsInput, optFns ...func(*Options)) (*GreetingWithErrorsOutput, error) {
	stack := middleware.NewStack("GreetingWithErrors", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	v4.AddHTTPSignerMiddleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGreetingWithErrors(options.Region), middleware.Before)
	addawsAwsjson11_serdeOpGreetingWithErrorsMiddlewares(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "GreetingWithErrors",
			Err:           err,
		}
	}
	out := result.(*GreetingWithErrorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GreetingWithErrorsInput struct {
}

type GreetingWithErrorsOutput struct {
	Greeting *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGreetingWithErrorsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGreetingWithErrors{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGreetingWithErrors{}, middleware.After)
}

func newServiceMetadataMiddleware_opGreetingWithErrors(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Json Protocol",
		ServiceID:      "jsonprotocol",
		EndpointPrefix: "jsonprotocol",
		SigningName:    "foo",
		OperationName:  "GreetingWithErrors",
	}
}