// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns an array of resources successfully backed up by AWS Backup, including
// the time the resource was saved, an Amazon Resource Name (ARN) of the resource,
// and a resource type.
func (c *Client) ListProtectedResources(ctx context.Context, params *ListProtectedResourcesInput, optFns ...func(*Options)) (*ListProtectedResourcesOutput, error) {
	stack := middleware.NewStack("ListProtectedResources", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListProtectedResourcesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListProtectedResources(options.Region), middleware.Before)
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
			OperationName: "ListProtectedResources",
			Err:           err,
		}
	}
	out := result.(*ListProtectedResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProtectedResourcesInput struct {
	// The next item following a partial list of returned items. For example, if a
	// request is made to return maxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string
	// The maximum number of items to be returned.
	MaxResults *int32
}

type ListProtectedResourcesOutput struct {
	// The next item following a partial list of returned items. For example, if a
	// request is made to return maxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string
	// An array of resources successfully backed up by AWS Backup including the time
	// the resource was saved, an Amazon Resource Name (ARN) of the resource, and a
	// resource type.
	Results []*types.ProtectedResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListProtectedResourcesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListProtectedResources{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListProtectedResources{}, middleware.After)
}

func newServiceMetadataMiddleware_opListProtectedResources(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "ListProtectedResources",
	}
}
