// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an existing virtual router in a specified service mesh.
func (c *Client) UpdateVirtualRouter(ctx context.Context, params *UpdateVirtualRouterInput, optFns ...func(*Options)) (*UpdateVirtualRouterOutput, error) {
	stack := middleware.NewStack("UpdateVirtualRouter", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateVirtualRouterMiddlewares(stack)
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
	addOpUpdateVirtualRouterValidationMiddleware(stack)

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
			OperationName: "UpdateVirtualRouter",
			Err:           err,
		}
	}
	out := result.(*UpdateVirtualRouterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type UpdateVirtualRouterInput struct {
	// The name of the virtual router to update.
	VirtualRouterName *string
	// The name of the service mesh that the virtual router resides in.
	MeshName *string
	// The new virtual router specification to apply. This overwrites the existing
	// data.
	Spec *types.VirtualRouterSpec
	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. Up to 36 letters, numbers, hyphens, and underscores are allowed.
	ClientToken *string
	// The AWS IAM account ID of the service mesh owner. If the account ID is not your
	// own, then it's the ID of the account that shared the mesh with your account. For
	// more information about mesh sharing, see Working with shared meshes
	// (https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).
	MeshOwner *string
}

//
type UpdateVirtualRouterOutput struct {
	// A full description of the virtual router that was updated.
	VirtualRouter *types.VirtualRouterData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateVirtualRouterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateVirtualRouter{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateVirtualRouter{}, middleware.After)
}
