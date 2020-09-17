// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes one or more custom attributes from an Amazon ECS resource.
func (c *Client) DeleteAttributes(ctx context.Context, params *DeleteAttributesInput, optFns ...func(*Options)) (*DeleteAttributesOutput, error) {
	stack := middleware.NewStack("DeleteAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteAttributesMiddlewares(stack)
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
	addOpDeleteAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAttributes(options.Region), middleware.Before)
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
			OperationName: "DeleteAttributes",
			Err:           err,
		}
	}
	out := result.(*DeleteAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAttributesInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that contains
	// the resource to delete attributes. If you do not specify a cluster, the default
	// cluster is assumed.
	Cluster *string
	// The attributes to delete from your resource. You can specify up to 10 attributes
	// per request. For custom attributes, specify the attribute name and target ID,
	// but do not specify the value. If you specify the target ID using the short form,
	// you must also specify the target type.
	Attributes []*types.Attribute
}

type DeleteAttributesOutput struct {
	// A list of attribute objects that were successfully deleted from your resource.
	Attributes []*types.Attribute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "DeleteAttributes",
	}
}
