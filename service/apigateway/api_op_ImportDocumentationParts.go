// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

func (c *Client) ImportDocumentationParts(ctx context.Context, params *ImportDocumentationPartsInput, optFns ...func(*Options)) (*ImportDocumentationPartsOutput, error) {
	stack := middleware.NewStack("ImportDocumentationParts", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpImportDocumentationPartsMiddlewares(stack)
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
	addOpImportDocumentationPartsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opImportDocumentationParts(options.Region), middleware.Before)
	addResponseErrorWrapper(stack)
	addAcceptHeader(stack)

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
			OperationName: "ImportDocumentationParts",
			Err:           err,
		}
	}
	out := result.(*ImportDocumentationPartsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Import documentation parts from an external (e.g., OpenAPI) definition file.
type ImportDocumentationPartsInput struct {
	Name *string
	// [Required] The string identifier of the associated RestApi ().
	RestApiId *string
	Title     *string
	// A query parameter to specify whether to rollback the documentation importation
	// (true) or not (false) when a warning is encountered. The default value is false.
	FailOnWarnings   *bool
	TemplateSkipList []*string
	Template         *bool
	// A query parameter to indicate whether to overwrite (OVERWRITE) any existing
	// DocumentationParts () definition or to merge (MERGE) the new definition into the
	// existing one. The default value is MERGE.
	Mode types.PutMode
}

// A collection of the imported DocumentationPart () identifiers. This is used to
// return the result when documentation parts in an external (e.g., OpenAPI) file
// are imported into API Gateway Documenting an API
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api.html),
// documentationpart:import
// (https://docs.aws.amazon.com/apigateway/api-reference/link-relation/documentationpart-import/),
// DocumentationPart ()
type ImportDocumentationPartsOutput struct {
	// A list of the returned documentation part identifiers.
	Ids []*string
	// A list of warning messages reported during import of documentation parts.
	Warnings []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpImportDocumentationPartsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpImportDocumentationParts{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpImportDocumentationParts{}, middleware.After)
}

func newServiceMetadataMiddleware_opImportDocumentationParts(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "ImportDocumentationParts",
	}
}
