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

// Lists all the experiments in your account. The list can be filtered to show only
// experiments that were created in a specific time range. The list can be sorted
// by experiment name or creation time.
func (c *Client) ListExperiments(ctx context.Context, params *ListExperimentsInput, optFns ...func(*Options)) (*ListExperimentsOutput, error) {
	stack := middleware.NewStack("ListExperiments", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListExperimentsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListExperiments(options.Region), middleware.Before)
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
			OperationName: "ListExperiments",
			Err:           err,
		}
	}
	out := result.(*ListExperimentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListExperimentsInput struct {
	// The sort order. The default value is Descending.
	SortOrder types.SortOrder
	// A filter that returns only experiments created before the specified time.
	CreatedBefore *time.Time
	// A filter that returns only experiments created after the specified time.
	CreatedAfter *time.Time
	// The property used to sort results. The default value is CreationTime.
	SortBy types.SortExperimentsBy
	// If the previous call to ListExperiments didn't return the full set of
	// experiments, the call returns a token for getting the next set of experiments.
	NextToken *string
	// The maximum number of experiments to return in the response. The default value
	// is 10.
	MaxResults *int32
}

type ListExperimentsOutput struct {
	// A token for getting the next set of experiments, if there are any.
	NextToken *string
	// A list of the summaries of your experiments.
	ExperimentSummaries []*types.ExperimentSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListExperimentsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListExperiments{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListExperiments{}, middleware.After)
}

func newServiceMetadataMiddleware_opListExperiments(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListExperiments",
	}
}
