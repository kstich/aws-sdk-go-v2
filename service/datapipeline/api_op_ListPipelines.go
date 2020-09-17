// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the pipeline identifiers for all active pipelines that you have permission
// to access.
func (c *Client) ListPipelines(ctx context.Context, params *ListPipelinesInput, optFns ...func(*Options)) (*ListPipelinesOutput, error) {
	stack := middleware.NewStack("ListPipelines", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListPipelinesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPipelines(options.Region), middleware.Before)
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
			OperationName: "ListPipelines",
			Err:           err,
		}
	}
	out := result.(*ListPipelinesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for ListPipelines.
type ListPipelinesInput struct {
	// The starting point for the results to be returned. For the first call, this
	// value should be empty. As long as there are more results, continue to call
	// ListPipelines with the marker value from the previous call to retrieve the next
	// set of results.
	Marker *string
}

// Contains the output of ListPipelines.
type ListPipelinesOutput struct {
	// Indicates whether there are more results that can be obtained by a subsequent
	// call.
	HasMoreResults *bool
	// The pipeline identifiers. If you require additional information about the
	// pipelines, you can use these identifiers to call DescribePipelines () and
	// GetPipelineDefinition ().
	PipelineIdList []*types.PipelineIdName
	// The starting point for the next page of results. To view the next page of
	// results, call ListPipelinesOutput again with this marker value. If the value is
	// null, there are no more results.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListPipelinesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListPipelines{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPipelines{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPipelines(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datapipeline",
		OperationName: "ListPipelines",
	}
}
