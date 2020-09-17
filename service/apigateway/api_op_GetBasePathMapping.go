// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describe a BasePathMapping () resource.
func (c *Client) GetBasePathMapping(ctx context.Context, params *GetBasePathMappingInput, optFns ...func(*Options)) (*GetBasePathMappingOutput, error) {
	stack := middleware.NewStack("GetBasePathMapping", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetBasePathMappingMiddlewares(stack)
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
	addOpGetBasePathMappingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBasePathMapping(options.Region), middleware.Before)
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
			OperationName: "GetBasePathMapping",
			Err:           err,
		}
	}
	out := result.(*GetBasePathMappingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to describe a BasePathMapping () resource.
type GetBasePathMappingInput struct {
	// [Required] The domain name of the BasePathMapping () resource to be described.
	DomainName *string
	// [Required] The base path name that callers of the API must provide as part of
	// the URL after the domain name. This value must be unique for all of the mappings
	// across a single API. Specify '(none)' if you do not want callers to specify any
	// base path name after the domain name.
	BasePath         *string
	Title            *string
	Name             *string
	TemplateSkipList []*string
	Template         *bool
}

// Represents the base path that callers of the API must provide as part of the URL
// after the domain name. A custom domain name plus a BasePathMapping specification
// identifies a deployed RestApi () in a given stage of the owner Account (). Use
// Custom Domain Names
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html)
type GetBasePathMappingOutput struct {
	// The name of the associated stage.
	Stage *string
	// The base path name that callers of the API must provide as part of the URL after
	// the domain name.
	BasePath *string
	// The string identifier of the associated RestApi ().
	RestApiId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetBasePathMappingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetBasePathMapping{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBasePathMapping{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetBasePathMapping(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetBasePathMapping",
	}
}
