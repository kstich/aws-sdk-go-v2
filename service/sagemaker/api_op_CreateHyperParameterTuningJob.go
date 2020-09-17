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

// Starts a hyperparameter tuning job. A hyperparameter tuning job finds the best
// version of a model by running many training jobs on your dataset using the
// algorithm you choose and values for hyperparameters within ranges that you
// specify. It then chooses the hyperparameter values that result in a model that
// performs the best, as measured by an objective metric that you choose.
func (c *Client) CreateHyperParameterTuningJob(ctx context.Context, params *CreateHyperParameterTuningJobInput, optFns ...func(*Options)) (*CreateHyperParameterTuningJobOutput, error) {
	stack := middleware.NewStack("CreateHyperParameterTuningJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateHyperParameterTuningJobMiddlewares(stack)
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
	addOpCreateHyperParameterTuningJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHyperParameterTuningJob(options.Region), middleware.Before)
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
			OperationName: "CreateHyperParameterTuningJob",
			Err:           err,
		}
	}
	out := result.(*CreateHyperParameterTuningJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateHyperParameterTuningJobInput struct {
	// An array of key-value pairs. You can use tags to categorize your AWS resources
	// in different ways, for example, by purpose, owner, or environment. For more
	// information, see AWS Tagging Strategies
	// (https://aws.amazon.com/answers/account-management/aws-tagging-strategies/).
	// Tags that you specify for the tuning job are also added to all training jobs
	// that the tuning job launches.
	Tags []*types.Tag
	// The name of the tuning job. This name is the prefix for the names of all
	// training jobs that this tuning job launches. The name must be unique within the
	// same AWS account and AWS Region. The name must have { } to { } characters. Valid
	// characters are a-z, A-Z, 0-9, and : + = @ _ % - (hyphen). The name is not case
	// sensitive.
	HyperParameterTuningJobName *string
	// A list of the HyperParameterTrainingJobDefinition () objects launched for this
	// tuning job.
	TrainingJobDefinitions []*types.HyperParameterTrainingJobDefinition
	// The HyperParameterTrainingJobDefinition () object that describes the training
	// jobs that this tuning job launches, including static hyperparameters, input data
	// configuration, output data configuration, resource configuration, and stopping
	// condition.
	TrainingJobDefinition *types.HyperParameterTrainingJobDefinition
	// The HyperParameterTuningJobConfig () object that describes the tuning job,
	// including the search strategy, the objective metric used to evaluate training
	// jobs, ranges of parameters to search, and resource limits for the tuning job.
	// For more information, see How Hyperparameter Tuning Works
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-how-it-works.html).
	HyperParameterTuningJobConfig *types.HyperParameterTuningJobConfig
	// Specifies the configuration for starting the hyperparameter tuning job using one
	// or more previous tuning jobs as a starting point. The results of previous tuning
	// jobs are used to inform which combinations of hyperparameters to search over in
	// the new tuning job. All training jobs launched by the new hyperparameter tuning
	// job are evaluated by using the objective metric. If you specify
	// IDENTICAL_DATA_AND_ALGORITHM as the WarmStartType value for the warm start
	// configuration, the training job that performs the best in the new tuning job is
	// compared to the best training jobs from the parent tuning jobs. From these, the
	// training job that performs the best as measured by the objective metric is
	// returned as the overall best training job. All training jobs launched by parent
	// hyperparameter tuning jobs and the new hyperparameter tuning jobs count against
	// the limit of training jobs for the tuning job.
	WarmStartConfig *types.HyperParameterTuningJobWarmStartConfig
}

type CreateHyperParameterTuningJobOutput struct {
	// The Amazon Resource Name (ARN) of the tuning job. Amazon SageMaker assigns an
	// ARN to a hyperparameter tuning job when you create it.
	HyperParameterTuningJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateHyperParameterTuningJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateHyperParameterTuningJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateHyperParameterTuningJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateHyperParameterTuningJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateHyperParameterTuningJob",
	}
}
