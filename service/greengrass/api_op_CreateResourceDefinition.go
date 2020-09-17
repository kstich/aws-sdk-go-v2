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

// Creates a resource definition which contains a list of resources to be used in a
// group. You can create an initial version of the definition by providing a list
// of resources now, or use ''CreateResourceDefinitionVersion'' later.
func (c *Client) CreateResourceDefinition(ctx context.Context, params *CreateResourceDefinitionInput, optFns ...func(*Options)) (*CreateResourceDefinitionOutput, error) {
	stack := middleware.NewStack("CreateResourceDefinition", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateResourceDefinitionMiddlewares(stack)
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
	addOpCreateResourceDefinitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateResourceDefinition(options.Region), middleware.Before)
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
			OperationName: "CreateResourceDefinition",
			Err:           err,
		}
	}
	out := result.(*CreateResourceDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateResourceDefinitionInput struct {
	// The name of the resource definition.
	Name *string
	// Tag(s) to add to the new resource.
	Tags map[string]*string
	// Information about the initial version of the resource definition.
	InitialVersion *types.ResourceDefinitionVersion
	// A client token used to correlate requests and responses.
	AmznClientToken *string
}

type CreateResourceDefinitionOutput struct {
	// The time, in milliseconds since the epoch, when the definition was created.
	CreationTimestamp *string
	// The ID of the latest version associated with the definition.
	LatestVersion *string
	// The time, in milliseconds since the epoch, when the definition was last updated.
	LastUpdatedTimestamp *string
	// The ARN of the latest version associated with the definition.
	LatestVersionArn *string
	// The ID of the definition.
	Id *string
	// The name of the definition.
	Name *string
	// The ARN of the definition.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateResourceDefinitionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateResourceDefinition{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateResourceDefinition{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateResourceDefinition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "CreateResourceDefinition",
	}
}
