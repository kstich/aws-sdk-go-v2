// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the specified attributes of the specified target group.
func (c *Client) ModifyTargetGroupAttributes(ctx context.Context, params *ModifyTargetGroupAttributesInput, optFns ...func(*Options)) (*ModifyTargetGroupAttributesOutput, error) {
	stack := middleware.NewStack("ModifyTargetGroupAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpModifyTargetGroupAttributesMiddlewares(stack)
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
	addOpModifyTargetGroupAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyTargetGroupAttributes(options.Region), middleware.Before)
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
			OperationName: "ModifyTargetGroupAttributes",
			Err:           err,
		}
	}
	out := result.(*ModifyTargetGroupAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyTargetGroupAttributesInput struct {
	// The Amazon Resource Name (ARN) of the target group.
	TargetGroupArn *string
	// The attributes.
	Attributes []*types.TargetGroupAttribute
}

type ModifyTargetGroupAttributesOutput struct {
	// Information about the attributes.
	Attributes []*types.TargetGroupAttribute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpModifyTargetGroupAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpModifyTargetGroupAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyTargetGroupAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyTargetGroupAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "ModifyTargetGroupAttributes",
	}
}
