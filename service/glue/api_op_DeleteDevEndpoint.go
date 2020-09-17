// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a specified development endpoint.
func (c *Client) DeleteDevEndpoint(ctx context.Context, params *DeleteDevEndpointInput, optFns ...func(*Options)) (*DeleteDevEndpointOutput, error) {
	stack := middleware.NewStack("DeleteDevEndpoint", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteDevEndpointMiddlewares(stack)
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
	addOpDeleteDevEndpointValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDevEndpoint(options.Region), middleware.Before)
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
			OperationName: "DeleteDevEndpoint",
			Err:           err,
		}
	}
	out := result.(*DeleteDevEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDevEndpointInput struct {
	// The name of the DevEndpoint.
	EndpointName *string
}

type DeleteDevEndpointOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteDevEndpointMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteDevEndpoint{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteDevEndpoint{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteDevEndpoint(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "DeleteDevEndpoint",
	}
}
