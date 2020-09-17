// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes a statement from the permissions policy for a version of an AWS Lambda
// layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
// For more information, see AddLayerVersionPermission ().
func (c *Client) RemoveLayerVersionPermission(ctx context.Context, params *RemoveLayerVersionPermissionInput, optFns ...func(*Options)) (*RemoveLayerVersionPermissionOutput, error) {
	stack := middleware.NewStack("RemoveLayerVersionPermission", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpRemoveLayerVersionPermissionMiddlewares(stack)
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
	addOpRemoveLayerVersionPermissionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveLayerVersionPermission(options.Region), middleware.Before)
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
			OperationName: "RemoveLayerVersionPermission",
			Err:           err,
		}
	}
	out := result.(*RemoveLayerVersionPermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveLayerVersionPermissionInput struct {
	// The version number.
	VersionNumber *int64
	// The name or Amazon Resource Name (ARN) of the layer.
	LayerName *string
	// The identifier that was specified when the statement was added.
	StatementId *string
	// Only update the policy if the revision ID matches the ID specified. Use this
	// option to avoid modifying a policy that has changed since you last read it.
	RevisionId *string
}

type RemoveLayerVersionPermissionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpRemoveLayerVersionPermissionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpRemoveLayerVersionPermission{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpRemoveLayerVersionPermission{}, middleware.After)
}

func newServiceMetadataMiddleware_opRemoveLayerVersionPermission(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "RemoveLayerVersionPermission",
	}
}
