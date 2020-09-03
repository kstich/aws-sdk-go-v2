// Code generated by smithy-go-codegen DO NOT EDIT.

package elastictranscoder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The DeletePreset operation removes a preset that you've added in an AWS region.
// You can't delete the default presets that are included with Elastic Transcoder.
func (c *Client) DeletePreset(ctx context.Context, params *DeletePresetInput, optFns ...func(*Options)) (*DeletePresetOutput, error) {
	stack := middleware.NewStack("DeletePreset", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeletePresetMiddlewares(stack)
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
	addOpDeletePresetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePreset(options.Region), middleware.Before)

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
			OperationName: "DeletePreset",
			Err:           err,
		}
	}
	out := result.(*DeletePresetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The DeletePresetRequest structure.
type DeletePresetInput struct {
	// The identifier of the preset for which you want to get detailed information.
	Id *string
}

// The DeletePresetResponse structure.
type DeletePresetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeletePresetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeletePreset{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeletePreset{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeletePreset(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elastictranscoder",
		OperationName: "DeletePreset",
	}
}
