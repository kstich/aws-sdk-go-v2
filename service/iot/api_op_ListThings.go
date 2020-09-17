// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists your things. Use the attributeName and attributeValue parameters to filter
// your things. For example, calling ListThings with attributeName=Color and
// attributeValue=Red retrieves all things in the registry that contain an
// attribute Color with the value Red.
func (c *Client) ListThings(ctx context.Context, params *ListThingsInput, optFns ...func(*Options)) (*ListThingsOutput, error) {
	stack := middleware.NewStack("ListThings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListThingsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListThings(options.Region), middleware.Before)
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
			OperationName: "ListThings",
			Err:           err,
		}
	}
	out := result.(*ListThingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the ListThings operation.
type ListThingsInput struct {
	// The attribute value used to search for things.
	AttributeValue *string
	// The name of the thing type used to search for things.
	ThingTypeName *string
	// The maximum number of results to return in this operation.
	MaxResults *int32
	// The token to retrieve the next set of results.
	NextToken *string
	// The attribute name used to search for things.
	AttributeName *string
}

// The output from the ListThings operation.
type ListThingsOutput struct {
	// The token used to get the next set of results, or null if there are no
	// additional results.
	NextToken *string
	// The things.
	Things []*types.ThingAttribute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListThingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListThings{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListThings{}, middleware.After)
}

func newServiceMetadataMiddleware_opListThings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListThings",
	}
}
