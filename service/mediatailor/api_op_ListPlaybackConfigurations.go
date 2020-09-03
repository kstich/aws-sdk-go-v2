// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of the playback configurations defined in AWS Elemental
// MediaTailor. You can specify a maximum number of configurations to return at a
// time. The default maximum is 50. Results are returned in pagefuls. If
// MediaTailor has more configurations than the specified maximum, it provides
// parameters in the response that you can use to retrieve the next pageful.
func (c *Client) ListPlaybackConfigurations(ctx context.Context, params *ListPlaybackConfigurationsInput, optFns ...func(*Options)) (*ListPlaybackConfigurationsOutput, error) {
	stack := middleware.NewStack("ListPlaybackConfigurations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListPlaybackConfigurationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPlaybackConfigurations(options.Region), middleware.Before)

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
			OperationName: "ListPlaybackConfigurations",
			Err:           err,
		}
	}
	out := result.(*ListPlaybackConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPlaybackConfigurationsInput struct {
	// Maximum number of records to return.
	MaxResults *int32
	// Pagination token returned by the GET list request when results exceed the
	// maximum allowed. Use the token to fetch the next page of results.
	NextToken *string
}

type ListPlaybackConfigurationsOutput struct {
	// Pagination token returned by the GET list request when results exceed the
	// maximum allowed. Use the token to fetch the next page of results.
	NextToken *string
	// Array of playback configurations. This might be all the available configurations
	// or a subset, depending on the settings that you provide and the total number of
	// configurations stored.
	Items []*types.PlaybackConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListPlaybackConfigurationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListPlaybackConfigurations{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListPlaybackConfigurations{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPlaybackConfigurations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediatailor",
		OperationName: "ListPlaybackConfigurations",
	}
}
