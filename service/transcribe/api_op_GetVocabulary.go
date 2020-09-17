// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Gets information about a vocabulary.
func (c *Client) GetVocabulary(ctx context.Context, params *GetVocabularyInput, optFns ...func(*Options)) (*GetVocabularyOutput, error) {
	stack := middleware.NewStack("GetVocabulary", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetVocabularyMiddlewares(stack)
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
	addOpGetVocabularyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetVocabulary(options.Region), middleware.Before)
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
			OperationName: "GetVocabulary",
			Err:           err,
		}
	}
	out := result.(*GetVocabularyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetVocabularyInput struct {
	// The name of the vocabulary to return information about. The name is
	// case-sensitive.
	VocabularyName *string
}

type GetVocabularyOutput struct {
	// The date and time that the vocabulary was last modified.
	LastModifiedTime *time.Time
	// The S3 location where the vocabulary is stored. Use this URI to get the contents
	// of the vocabulary. The URI is available for a limited time.
	DownloadUri *string
	// The processing state of the vocabulary.
	VocabularyState types.VocabularyState
	// The name of the vocabulary to return.
	VocabularyName *string
	// The language code of the vocabulary entries.
	LanguageCode types.LanguageCode
	// If the VocabularyState field is FAILED, this field contains information about
	// why the job failed.
	FailureReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetVocabularyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetVocabulary{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetVocabulary{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetVocabulary(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "GetVocabulary",
	}
}
