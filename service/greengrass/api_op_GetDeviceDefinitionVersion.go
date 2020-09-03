// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about a device definition version.
func (c *Client) GetDeviceDefinitionVersion(ctx context.Context, params *GetDeviceDefinitionVersionInput, optFns ...func(*Options)) (*GetDeviceDefinitionVersionOutput, error) {
	stack := middleware.NewStack("GetDeviceDefinitionVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetDeviceDefinitionVersionMiddlewares(stack)
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
	addOpGetDeviceDefinitionVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDeviceDefinitionVersion(options.Region), middleware.Before)

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
			OperationName: "GetDeviceDefinitionVersion",
			Err:           err,
		}
	}
	out := result.(*GetDeviceDefinitionVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDeviceDefinitionVersionInput struct {
	// The ID of the device definition version. This value maps to the ''Version''
	// property of the corresponding ''VersionInformation'' object, which is returned
	// by ''ListDeviceDefinitionVersions'' requests. If the version is the last one
	// that was associated with a device definition, the value also maps to the
	// ''LatestVersion'' property of the corresponding ''DefinitionInformation''
	// object.
	DeviceDefinitionVersionId *string
	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string
	// The ID of the device definition.
	DeviceDefinitionId *string
}

type GetDeviceDefinitionVersionOutput struct {
	// The version of the device definition version.
	Version *string
	// The ID of the device definition version.
	Id *string
	// The time, in milliseconds since the epoch, when the device definition version
	// was created.
	CreationTimestamp *string
	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string
	// Information about the device definition version.
	Definition *types.DeviceDefinitionVersion
	// The ARN of the device definition version.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetDeviceDefinitionVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetDeviceDefinitionVersion{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDeviceDefinitionVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDeviceDefinitionVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "GetDeviceDefinitionVersion",
	}
}
