// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Update a field-level encryption configuration.
func (c *Client) UpdateFieldLevelEncryptionConfig(ctx context.Context, params *UpdateFieldLevelEncryptionConfigInput, optFns ...func(*Options)) (*UpdateFieldLevelEncryptionConfigOutput, error) {
	stack := middleware.NewStack("UpdateFieldLevelEncryptionConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpUpdateFieldLevelEncryptionConfigMiddlewares(stack)
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
	addOpUpdateFieldLevelEncryptionConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFieldLevelEncryptionConfig(options.Region), middleware.Before)
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
			OperationName: "UpdateFieldLevelEncryptionConfig",
			Err:           err,
		}
	}
	out := result.(*UpdateFieldLevelEncryptionConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFieldLevelEncryptionConfigInput struct {
	// Request to update a field-level encryption configuration.
	FieldLevelEncryptionConfig *types.FieldLevelEncryptionConfig
	// The ID of the configuration you want to update.
	Id *string
	// The value of the ETag header that you received when retrieving the configuration
	// identity to update. For example: E2QWRUHAPOMQZL.
	IfMatch *string
}

type UpdateFieldLevelEncryptionConfigOutput struct {
	// The value of the ETag header that you received when updating the configuration.
	// For example: E2QWRUHAPOMQZL.
	ETag *string
	// Return the results of updating the configuration.
	FieldLevelEncryption *types.FieldLevelEncryption

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpUpdateFieldLevelEncryptionConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpUpdateFieldLevelEncryptionConfig{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpUpdateFieldLevelEncryptionConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateFieldLevelEncryptionConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "UpdateFieldLevelEncryptionConfig",
	}
}
