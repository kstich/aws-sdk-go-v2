// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Aborts the upload of the specified document version that was previously
// initiated by InitiateDocumentVersionUpload (). The client should make this call
// only when it no longer intends to upload the document version, or fails to do
// so.
func (c *Client) AbortDocumentVersionUpload(ctx context.Context, params *AbortDocumentVersionUploadInput, optFns ...func(*Options)) (*AbortDocumentVersionUploadOutput, error) {
	stack := middleware.NewStack("AbortDocumentVersionUpload", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpAbortDocumentVersionUploadMiddlewares(stack)
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
	addOpAbortDocumentVersionUploadValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAbortDocumentVersionUpload(options.Region), middleware.Before)

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
			OperationName: "AbortDocumentVersionUpload",
			Err:           err,
		}
	}
	out := result.(*AbortDocumentVersionUploadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AbortDocumentVersionUploadInput struct {
	// The ID of the document.
	DocumentId *string
	// The ID of the version.
	VersionId *string
	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string
}

type AbortDocumentVersionUploadOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpAbortDocumentVersionUploadMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpAbortDocumentVersionUpload{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpAbortDocumentVersionUpload{}, middleware.After)
}

func newServiceMetadataMiddleware_opAbortDocumentVersionUpload(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "AbortDocumentVersionUpload",
	}
}
