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

// Lists the trials in your account. Specify an experiment name to limit the list
// to the trials that are part of that experiment. Specify a trial component name
// to limit the list to the trials that associated with that trial component. The
// list can be filtered to show only trials that were created in a specific time
// range. The list can be sorted by trial name or creation time.
func (c *Client) ListTrials(ctx context.Context, params *ListTrialsInput, optFns ...func(*Options)) (*ListTrialsOutput, error) {
	stack := middleware.NewStack("ListTrials", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListTrialsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTrials(options.Region), middleware.Before)
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
			OperationName: "ListTrials",
			Err:           err,
		}
	}
	out := result.(*ListTrialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTrialsInput struct {
	// A filter that returns only trials that are part of the specified experiment.
	ExperimentName *string
	// If the previous call to ListTrials didn't return the full set of trials, the
	// call returns a token for getting the next set of trials.
	NextToken *string
	// The maximum number of trials to return in the response. The default value is 10.
	MaxResults *int32
	// The property used to sort results. The default value is CreationTime.
	SortBy types.SortTrialsBy
	// The sort order. The default value is Descending.
	SortOrder types.SortOrder
	// A filter that returns only trials created after the specified time.
	CreatedAfter *time.Time
	// A filter that returns only trials that are associated with the specified trial
	// component.
	TrialComponentName *string
	// A filter that returns only trials created before the specified time.
	CreatedBefore *time.Time
}

type ListTrialsOutput struct {
	// A token for getting the next set of trials, if there are any.
	NextToken *string
	// A list of the summaries of your trials.
	TrialSummaries []*types.TrialSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListTrialsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListTrials{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTrials{}, middleware.After)
}

func newServiceMetadataMiddleware_opListTrials(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListTrials",
	}
}
