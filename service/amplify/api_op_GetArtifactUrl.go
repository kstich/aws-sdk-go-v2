// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the artifact info that corresponds to an artifact id.
func (c *Client) GetArtifactUrl(ctx context.Context, params *GetArtifactUrlInput, optFns ...func(*Options)) (*GetArtifactUrlOutput, error) {
	stack := middleware.NewStack("GetArtifactUrl", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetArtifactUrlMiddlewares(stack)
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
	addOpGetArtifactUrlValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetArtifactUrl(options.Region), middleware.Before)

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
			OperationName: "GetArtifactUrl",
			Err:           err,
		}
	}
	out := result.(*GetArtifactUrlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Returns the request structure for the get artifact request.
type GetArtifactUrlInput struct {
	// The unique ID for an artifact.
	ArtifactId *string
}

// Returns the result structure for the get artifact request.
type GetArtifactUrlOutput struct {
	// The presigned URL for the artifact.
	ArtifactUrl *string
	// The unique ID for an artifact.
	ArtifactId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetArtifactUrlMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetArtifactUrl{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetArtifactUrl{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetArtifactUrl(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplify",
		OperationName: "GetArtifactUrl",
	}
}
