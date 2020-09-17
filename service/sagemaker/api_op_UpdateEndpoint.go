// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deploys the new EndpointConfig specified in the request, switches to using newly
// created endpoint, and then deletes resources provisioned for the endpoint using
// the previous EndpointConfig (there is no availability loss). When Amazon
// SageMaker receives the request, it sets the endpoint status to Updating. After
// updating the endpoint, it sets the status to InService. To check the status of
// an endpoint, use the DescribeEndpoint () API.  </p> <note> <p>You must not
// delete an <code>EndpointConfig</code> in use by an endpoint that is live or
// while the <code>UpdateEndpoint</code> or <code>CreateEndpoint</code> operations
// are being performed on the endpoint. To update an endpoint, you must create a
// new <code>EndpointConfig</code>.</p> <p>If you delete the
// <code>EndpointConfig</code> of an endpoint that is active or being created or
// updated you may lose visibility into the instance type the endpoint is using.
// The endpoint must be deleted in order to stop incurring charges.</p> </note>
func (c *Client) UpdateEndpoint(ctx context.Context, params *UpdateEndpointInput, optFns ...func(*Options)) (*UpdateEndpointOutput, error) {
	stack := middleware.NewStack("UpdateEndpoint", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateEndpointMiddlewares(stack)
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
	// The name of the new endpoint configuration.
	EndpointConfigName *string
	// When updating endpoint resources, enables or disables the retention of variant
	// properties, such as the instance count or the variant weight. To retain the
	// variant properties of an endpoint when updating it, set
	// RetainAllVariantProperties to true. To use the variant properties specified in a
	// new EndpointConfig call when updating an endpoint, set
	// RetainAllVariantProperties to false.
	RetainAllVariantProperties *bool
	// When you are updating endpoint resources with
	// UpdateEndpointInput$RetainAllVariantProperties (), whose value is set to true,
	// ExcludeRetainedVariantProperties specifies the list of type VariantProperty ()
	// to override with the values provided by EndpointConfig. If you don't specify a
	// value for ExcludeAllVariantProperties, no variant properties are overridden.
	ExcludeRetainedVariantProperties []*types.VariantProperty
	// The name of the endpoint whose configuration you want to update.
	EndpointName *string
}

type UpdateEndpointOutput struct {
	// The Amazon Resource Name (ARN) of the endpoint.
	EndpointArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateEndpointMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateEndpoint{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateEndpoint{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateEndpoint(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "UpdateEndpoint",
	}
}
