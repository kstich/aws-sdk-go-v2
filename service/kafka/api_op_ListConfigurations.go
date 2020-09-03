// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of all the MSK configurations in this Region.
func (c *Client) ListConfigurations(ctx context.Context, params *ListConfigurationsInput, optFns ...func(*Options)) (*ListConfigurationsOutput, error) {
	stack := middleware.NewStack("ListConfigurations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListConfigurationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListConfigurations(options.Region), middleware.Before)

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
			OperationName: "ListConfigurations",
			Err:           err,
		}
	}
	out := result.(*ListConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConfigurationsInput struct {
	// The paginated results marker. When the result of the operation is truncated, the
	// call returns NextToken in the response. To get the next batch, provide this
	// token in your next request.
	NextToken *string
	// The maximum number of results to return in the response. If there are more
	// results, the response includes a NextToken parameter.
	MaxResults *int32
}

type ListConfigurationsOutput struct {
	// The paginated results marker. When the result of a ListConfigurations operation
	// is truncated, the call returns NextToken in the response. To get another batch
	// of configurations, provide this token in your next request.
	NextToken *string
	// An array of MSK configurations.
	Configurations []*types.Configuration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListConfigurationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListConfigurations{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListConfigurations{}, middleware.After)
}

func newServiceMetadataMiddleware_opListConfigurations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kafka",
		OperationName: "ListConfigurations",
	}
}
