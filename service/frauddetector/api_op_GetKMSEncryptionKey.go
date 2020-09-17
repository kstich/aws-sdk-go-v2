// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the encryption key if a Key Management Service (KMS) customer master key
// (CMK) has been specified to be used to encrypt content in Amazon Fraud Detector.
func (c *Client) GetKMSEncryptionKey(ctx context.Context, params *GetKMSEncryptionKeyInput, optFns ...func(*Options)) (*GetKMSEncryptionKeyOutput, error) {
	stack := middleware.NewStack("GetKMSEncryptionKey", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetKMSEncryptionKeyMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetKMSEncryptionKey(options.Region), middleware.Before)
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
			OperationName: "GetKMSEncryptionKey",
			Err:           err,
		}
	}
	out := result.(*GetKMSEncryptionKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetKMSEncryptionKeyInput struct {
}

type GetKMSEncryptionKeyOutput struct {
	// The KMS encryption key.
	KmsKey *types.KMSKey

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetKMSEncryptionKeyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetKMSEncryptionKey{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetKMSEncryptionKey{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetKMSEncryptionKey(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "GetKMSEncryptionKey",
	}
}
