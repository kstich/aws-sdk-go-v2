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

// Lists all attributes that are associated with an object.
func (c *Client) ListObjectAttributes(ctx context.Context, params *ListObjectAttributesInput, optFns ...func(*Options)) (*ListObjectAttributesOutput, error) {
	stack := middleware.NewStack("ListObjectAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListObjectAttributesMiddlewares(stack)
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
	addOpListObjectAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListObjectAttributes(options.Region), middleware.Before)
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
			OperationName: "ListObjectAttributes",
			Err:           err,
		}
	}
	out := result.(*ListObjectAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListObjectAttributesInput struct {
	// The Amazon Resource Name (ARN) that is associated with the Directory () where
	// the object resides. For more information, see arns ().
	DirectoryArn *string
	// Represents the manner and timing in which the successful write or update of an
	// object is reflected in a subsequent read operation of that same object.
	ConsistencyLevel types.ConsistencyLevel
	// The reference that identifies the object whose attributes will be listed.
	ObjectReference *types.ObjectReference
	// The pagination token.
	NextToken *string
	// The maximum number of items to be retrieved in a single call. This is an
	// approximate number.
	MaxResults *int32
	// Used to filter the list of object attributes that are associated with a certain
	// facet.
	FacetFilter *types.SchemaFacet
}

type ListObjectAttributesOutput struct {
	// Attributes map that is associated with the object. AttributeArn is the key, and
	// attribute value is the value.
	Attributes []*types.AttributeKeyAndValue
	// The pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListObjectAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListObjectAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListObjectAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opListObjectAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListObjectAttributes",
	}
}
