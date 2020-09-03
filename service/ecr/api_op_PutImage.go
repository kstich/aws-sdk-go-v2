// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates or updates the image manifest and tags associated with an image. When an
// image is pushed and all new image layers have been uploaded, the PutImage API is
// called once to create or update the image manifest and the tags associated with
// the image.  <note> <p>This operation is used by the Amazon ECR proxy and is not
// generally used by customers for pulling and pushing images. In most cases, you
// should use the <code>docker</code> CLI to pull, tag, and push images.</p>
// </note>
func (c *Client) PutImage(ctx context.Context, params *PutImageInput, optFns ...func(*Options)) (*PutImageOutput, error) {
	stack := middleware.NewStack("PutImage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutImageMiddlewares(stack)
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
	addOpPutImageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutImage(options.Region), middleware.Before)

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
			OperationName: "PutImage",
			Err:           err,
		}
	}
	out := result.(*PutImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutImageInput struct {
	// The media type of the image manifest. If you push an image manifest that does
	// not contain the mediaType field, you must specify the imageManifestMediaType in
	// the request.
	ImageManifestMediaType *string
	// The AWS account ID associated with the registry that contains the repository in
	// which to put the image. If you do not specify a registry, the default registry
	// is assumed.
	RegistryId *string
	// The image manifest corresponding to the image to be uploaded.
	ImageManifest *string
	// The name of the repository in which to put the image.
	RepositoryName *string
	// The tag to associate with the image. This parameter is required for images that
	// use the Docker Image Manifest V2 Schema 2 or Open Container Initiative (OCI)
	// formats.
	ImageTag *string
	// The image digest of the image manifest corresponding to the image.
	ImageDigest *string
}

type PutImageOutput struct {
	// Details of the image uploaded.
	Image *types.Image

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutImageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutImage{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutImage{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutImage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "PutImage",
	}
}
