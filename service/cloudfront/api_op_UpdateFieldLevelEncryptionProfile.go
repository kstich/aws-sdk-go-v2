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

// Update a field-level encryption profile.
func (c *Client) UpdateFieldLevelEncryptionProfile(ctx context.Context, params *UpdateFieldLevelEncryptionProfileInput, optFns ...func(*Options)) (*UpdateFieldLevelEncryptionProfileOutput, error) {
	stack := middleware.NewStack("UpdateFieldLevelEncryptionProfile", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpUpdateFieldLevelEncryptionProfileMiddlewares(stack)
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
	addOpUpdateFieldLevelEncryptionProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFieldLevelEncryptionProfile(options.Region), middleware.Before)
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
			OperationName: "UpdateFieldLevelEncryptionProfile",
			Err:           err,
		}
	}
	out := result.(*UpdateFieldLevelEncryptionProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFieldLevelEncryptionProfileInput struct {
	// The value of the ETag header that you received when retrieving the profile
	// identity to update. For example: E2QWRUHAPOMQZL.
	IfMatch *string
	// Request to update a field-level encryption profile.
	FieldLevelEncryptionProfileConfig *types.FieldLevelEncryptionProfileConfig
	// The ID of the field-level encryption profile request.
	Id *string
}

type UpdateFieldLevelEncryptionProfileOutput struct {
	// The result of the field-level encryption profile request.
	ETag *string
	// Return the results of updating the profile.
	FieldLevelEncryptionProfile *types.FieldLevelEncryptionProfile

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpUpdateFieldLevelEncryptionProfileMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpUpdateFieldLevelEncryptionProfile{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpUpdateFieldLevelEncryptionProfile{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateFieldLevelEncryptionProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "UpdateFieldLevelEncryptionProfile",
	}
}
