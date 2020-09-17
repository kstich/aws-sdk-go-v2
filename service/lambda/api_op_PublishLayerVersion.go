// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an AWS Lambda layer
// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) from a
// ZIP archive. Each time you call PublishLayerVersion with the same layer name, a
// new version is created. Add layers to your function with CreateFunction () or
// UpdateFunctionConfiguration ().
func (c *Client) PublishLayerVersion(ctx context.Context, params *PublishLayerVersionInput, optFns ...func(*Options)) (*PublishLayerVersionOutput, error) {
	stack := middleware.NewStack("PublishLayerVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPublishLayerVersionMiddlewares(stack)
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
	addOpPublishLayerVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPublishLayerVersion(options.Region), middleware.Before)
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
			OperationName: "PublishLayerVersion",
			Err:           err,
		}
	}
	out := result.(*PublishLayerVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PublishLayerVersionInput struct {
	// The function layer archive.
	Content *types.LayerVersionContentInput
	// The description of the version.
	Description *string
	// The name or Amazon Resource Name (ARN) of the layer.
	LayerName *string
	// A list of compatible function runtimes
	// (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html). Used for
	// filtering with ListLayers () and ListLayerVersions ().
	CompatibleRuntimes []types.Runtime
	// The layer's software license. It can be any of the following:
	//
	//     * An SPDX
	// license identifier (https://spdx.org/licenses/). For example, MIT.
	//
	//     * The
	// URL of a license hosted on the internet. For example,
	// https://opensource.org/licenses/MIT.
	//
	//     * The full text of the license.
	LicenseInfo *string
}

type PublishLayerVersionOutput struct {
	// The description of the version.
	Description *string
	// The ARN of the layer version.
	LayerVersionArn *string
	// The version number.
	Version *int64
	// The date that the layer version was created, in ISO-8601 format
	// (https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).
	CreatedDate *string
	// The ARN of the layer.
	LayerArn *string
	// Details about the layer version.
	Content *types.LayerVersionContentOutput
	// The layer's compatible runtimes.
	CompatibleRuntimes []types.Runtime
	// The layer's software license.
	LicenseInfo *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPublishLayerVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPublishLayerVersion{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPublishLayerVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opPublishLayerVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "PublishLayerVersion",
	}
}
