// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Lists endpoints.
func (c *Client) ListEndpoints(ctx context.Context, params *ListEndpointsInput, optFns ...func(*Options)) (*ListEndpointsOutput, error) {
	stack := middleware.NewStack("ListEndpoints", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListEndpointsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListEndpoints(options.Region), middleware.Before)
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
			OperationName: "ListEndpoints",
			Err:           err,
		}
	}
	out := result.(*ListEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEndpointsInput struct {
	// A filter that returns only endpoints with the specified status.
	StatusEquals types.EndpointStatus
	// A filter that returns only endpoints that were modified after the specified
	// timestamp.
	LastModifiedTimeAfter *time.Time
	// Sorts the list of results. The default is CreationTime.
	SortBy types.EndpointSortKey
	// A filter that returns only endpoints that were created before the specified time
	// (timestamp).
	CreationTimeBefore *time.Time
	// If the result of a ListEndpoints request was truncated, the response includes a
	// NextToken. To retrieve the next set of endpoints, use the token in the next
	// request.
	NextToken *string
	// The maximum number of endpoints to return in the response.
	MaxResults *int32
	// A string in endpoint names. This filter returns only endpoints whose name
	// contains the specified string.
	NameContains *string
	// The sort order for results. The default is Descending.
	SortOrder types.OrderKey
	// A filter that returns only endpoints with a creation time greater than or equal
	// to the specified time (timestamp).
	CreationTimeAfter *time.Time
	// A filter that returns only endpoints that were modified before the specified
	// timestamp.
	LastModifiedTimeBefore *time.Time
}

type ListEndpointsOutput struct {
	// An array or endpoint objects.
	Endpoints []*types.EndpointSummary
	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of training jobs, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListEndpointsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListEndpoints{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEndpoints{}, middleware.After)
}

func newServiceMetadataMiddleware_opListEndpoints(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListEndpoints",
	}
}
