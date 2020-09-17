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
)

// Gets a list of TrainingJobSummary () objects that describe the training jobs
// that a hyperparameter tuning job launched.
func (c *Client) ListTrainingJobsForHyperParameterTuningJob(ctx context.Context, params *ListTrainingJobsForHyperParameterTuningJobInput, optFns ...func(*Options)) (*ListTrainingJobsForHyperParameterTuningJobOutput, error) {
	stack := middleware.NewStack("ListTrainingJobsForHyperParameterTuningJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListTrainingJobsForHyperParameterTuningJobMiddlewares(stack)
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
	addOpListTrainingJobsForHyperParameterTuningJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTrainingJobsForHyperParameterTuningJob(options.Region), middleware.Before)
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
			OperationName: "ListTrainingJobsForHyperParameterTuningJob",
			Err:           err,
		}
	}
	out := result.(*ListTrainingJobsForHyperParameterTuningJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTrainingJobsForHyperParameterTuningJobInput struct {
	// The field to sort results by. The default is Name. If the value of this field is
	// FinalObjectiveMetricValue, any training jobs that did not return an objective
	// metric are not listed.
	SortBy types.TrainingJobSortByOptions
	// The maximum number of training jobs to return. The default value is 10.
	MaxResults *int32
	// If the result of the previous ListTrainingJobsForHyperParameterTuningJob request
	// was truncated, the response includes a NextToken. To retrieve the next set of
	// training jobs, use the token in the next request.
	NextToken *string
	// The sort order for results. The default is Ascending.
	SortOrder types.SortOrder
	// The name of the tuning job whose training jobs you want to list.
	HyperParameterTuningJobName *string
	// A filter that returns only training jobs with the specified status.
	StatusEquals types.TrainingJobStatus
}

type ListTrainingJobsForHyperParameterTuningJobOutput struct {
	// A list of TrainingJobSummary () objects that describe the training jobs that the
	// ListTrainingJobsForHyperParameterTuningJob request returned.
	TrainingJobSummaries []*types.HyperParameterTrainingJobSummary
	// If the result of this ListTrainingJobsForHyperParameterTuningJob request was
	// truncated, the response includes a NextToken. To retrieve the next set of
	// training jobs, use the token in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListTrainingJobsForHyperParameterTuningJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListTrainingJobsForHyperParameterTuningJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTrainingJobsForHyperParameterTuningJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opListTrainingJobsForHyperParameterTuningJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListTrainingJobsForHyperParameterTuningJob",
	}
}
