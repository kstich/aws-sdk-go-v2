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

// Represents a collection of BasePathMapping () resources.
func (c *Client) GetBasePathMappings(ctx context.Context, params *GetBasePathMappingsInput, optFns ...func(*Options)) (*GetBasePathMappingsOutput, error) {
	stack := middleware.NewStack("GetBasePathMappings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetBasePathMappingsMiddlewares(stack)
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
	addOpGetBasePathMappingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBasePathMappings(options.Region), middleware.Before)
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
			OperationName: "GetBasePathMappings",
			Err:           err,
		}
	}
	out := result.(*GetBasePathMappingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to get information about a collection of BasePathMapping () resources.
type GetBasePathMappingsInput struct {
	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit *int32
	Name  *string
	// [Required] The domain name of a BasePathMapping () resource.
	DomainName       *string
	Template         *bool
	Title            *string
	TemplateSkipList []*string
	// The current pagination position in the paged result set.
	Position *string
}

// Represents a collection of BasePathMapping () resources. Use Custom Domain Names
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html)
type GetBasePathMappingsOutput struct {
	// The current page of elements from this collection.
	Items []*types.BasePathMapping
	// The current pagination position in the paged result set.
	Position *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetBasePathMappingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetBasePathMappings{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBasePathMappings{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetBasePathMappings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetBasePathMappings",
	}
}
