// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the identity attribute order for a specific TypedLinkFacet (). For more
// information, see Typed Links
// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
func (c *Client) GetTypedLinkFacetInformation(ctx context.Context, params *GetTypedLinkFacetInformationInput, optFns ...func(*Options)) (*GetTypedLinkFacetInformationOutput, error) {
	stack := middleware.NewStack("GetTypedLinkFacetInformation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetTypedLinkFacetInformationMiddlewares(stack)
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
	addOpGetTypedLinkFacetInformationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTypedLinkFacetInformation(options.Region), middleware.Before)
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
			OperationName: "GetTypedLinkFacetInformation",
			Err:           err,
		}
	}
	out := result.(*GetTypedLinkFacetInformationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTypedLinkFacetInformationInput struct {
	// The Amazon Resource Name (ARN) that is associated with the schema. For more
	// information, see arns ().
	SchemaArn *string
	// The unique name of the typed link facet.
	Name *string
}

type GetTypedLinkFacetInformationOutput struct {
	// The order of identity attributes for the facet, from most significant to least
	// significant. The ability to filter typed links considers the order that the
	// attributes are defined on the typed link facet. When providing ranges to typed
	// link selection, any inexact ranges must be specified at the end. Any attributes
	// that do not have a range specified are presumed to match the entire range.
	// Filters are interpreted in the order of the attributes on the typed link facet,
	// not the order in which they are supplied to any API calls. For more information
	// about identity attributes, see Typed Links
	// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
	IdentityAttributeOrder []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetTypedLinkFacetInformationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetTypedLinkFacetInformation{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTypedLinkFacetInformation{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetTypedLinkFacetInformation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "GetTypedLinkFacetInformation",
	}
}
