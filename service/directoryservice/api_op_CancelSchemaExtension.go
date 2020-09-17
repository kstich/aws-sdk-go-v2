// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Cancels an in-progress schema extension to a Microsoft AD directory. Once a
// schema extension has started replicating to all domain controllers, the task can
// no longer be canceled. A schema extension can be canceled during any of the
// following states; Initializing, CreatingSnapshot, and UpdatingSchema.
func (c *Client) CancelSchemaExtension(ctx context.Context, params *CancelSchemaExtensionInput, optFns ...func(*Options)) (*CancelSchemaExtensionOutput, error) {
	stack := middleware.NewStack("CancelSchemaExtension", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCancelSchemaExtensionMiddlewares(stack)
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
	addOpCancelSchemaExtensionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCancelSchemaExtension(options.Region), middleware.Before)
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
			OperationName: "CancelSchemaExtension",
			Err:           err,
		}
	}
	out := result.(*CancelSchemaExtensionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelSchemaExtensionInput struct {
	// The identifier of the directory whose schema extension will be canceled.
	DirectoryId *string
	// The identifier of the schema extension that will be canceled.
	SchemaExtensionId *string
}

type CancelSchemaExtensionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCancelSchemaExtensionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCancelSchemaExtension{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCancelSchemaExtension{}, middleware.After)
}

func newServiceMetadataMiddleware_opCancelSchemaExtension(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "CancelSchemaExtension",
	}
}
