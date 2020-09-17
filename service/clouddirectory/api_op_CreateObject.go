// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an object in a Directory (). Additionally attaches the object to a
// parent, if a parent reference and LinkName is specified. An object is simply a
// collection of Facet () attributes. You can also use this API call to create a
// policy object, if the facet from which you create the object is a policy facet.
func (c *Client) CreateObject(ctx context.Context, params *CreateObjectInput, optFns ...func(*Options)) (*CreateObjectOutput, error) {
	stack := middleware.NewStack("CreateObject", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateObjectMiddlewares(stack)
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
	addOpCreateObjectValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateObject(options.Region), middleware.Before)
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
			OperationName: "CreateObject",
			Err:           err,
		}
	}
	out := result.(*CreateObjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateObjectInput struct {
	// The Amazon Resource Name (ARN) that is associated with the Directory () in which
	// the object will be created. For more information, see arns ().
	DirectoryArn *string
	// The name of link that is used to attach this object to a parent.
	LinkName *string
	// If specified, the parent reference to which this object will be attached.
	ParentReference *types.ObjectReference
	// A list of schema facets to be associated with the object. Do not provide minor
	// version components. See SchemaFacet () for details.
	SchemaFacets []*types.SchemaFacet
	// The attribute map whose attribute ARN contains the key and attribute value as
	// the map value.
	ObjectAttributeList []*types.AttributeKeyAndValue
}

type CreateObjectOutput struct {
	// The identifier that is associated with the object.
	ObjectIdentifier *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateObjectMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateObject{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateObject{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateObject(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "CreateObject",
	}
}
