// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Attaches an existing object to another object. An object can be accessed in two
// ways:
//
//     * Using the path
//
//     * Using ObjectIdentifier
func (c *Client) AttachObject(ctx context.Context, params *AttachObjectInput, optFns ...func(*Options)) (*AttachObjectOutput, error) {
	stack := middleware.NewStack("AttachObject", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpAttachObjectMiddlewares(stack)
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
	addOpAttachObjectValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAttachObject(options.Region), middleware.Before)
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
			OperationName: "AttachObject",
			Err:           err,
		}
	}
	out := result.(*AttachObjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachObjectInput struct {
	// Amazon Resource Name (ARN) that is associated with the Directory () where both
	// objects reside. For more information, see arns ().
	DirectoryArn *string
	// The link name with which the child object is attached to the parent.
	LinkName *string
	// The child object reference to be attached to the object.
	ChildReference *types.ObjectReference
	// The parent object reference.
	ParentReference *types.ObjectReference
}

type AttachObjectOutput struct {
	// The attached ObjectIdentifier, which is the child ObjectIdentifier.
	AttachedObjectIdentifier *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpAttachObjectMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpAttachObject{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpAttachObject{}, middleware.After)
}

func newServiceMetadataMiddleware_opAttachObject(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "AttachObject",
	}
}
