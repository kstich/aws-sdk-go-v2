// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates and defines the settings for a classification job.
func (c *Client) CreateClassificationJob(ctx context.Context, params *CreateClassificationJobInput, optFns ...func(*Options)) (*CreateClassificationJobOutput, error) {
	stack := middleware.NewStack("CreateClassificationJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateClassificationJobMiddlewares(stack)
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
	addIdempotencyToken_opCreateClassificationJobMiddleware(stack, options)
	addOpCreateClassificationJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateClassificationJob(options.Region), middleware.Before)
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
			OperationName: "CreateClassificationJob",
			Err:           err,
		}
	}
	out := result.(*CreateClassificationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateClassificationJobInput struct {
	// A map of key-value pairs that specifies the tags to associate with the job. A
	// job can have a maximum of 50 tags. Each tag consists of a required tag key and
	// an associated tag value. The maximum length of a tag key is 128 characters. The
	// maximum length of a tag value is 256 characters.
	Tags map[string]*string
	// The S3 buckets that contain the objects to analyze, and the scope of that
	// analysis.
	S3JobDefinition *types.S3JobDefinition
	// A custom description of the job. The description can contain as many as 200
	// characters.
	Description *string
	// A custom name for the job. The name can contain as many as 500 characters.
	Name *string
	// Specifies whether to run the job immediately, after it's created.
	InitialRun *bool
	// The schedule for running the job. Valid values are:
	//
	//     * ONE_TIME - Run the
	// job only once. If you specify this value, don't specify a value for the
	// scheduleFrequency property.
	//
	//     * SCHEDULED - Run the job on a daily, weekly,
	// or monthly basis. If you specify this value, use the scheduleFrequency property
	// to define the recurrence pattern for the job.
	JobType types.JobType
	// A unique, case-sensitive token that you provide to ensure the idempotency of the
	// request.
	ClientToken *string
	// The recurrence pattern for running the job. To run the job only once, don't
	// specify a value for this property and set the value of the jobType property to
	// ONE_TIME.
	ScheduleFrequency *types.JobScheduleFrequency
	// The custom data identifiers to use for data analysis and classification.
	CustomDataIdentifierIds []*string
	// The sampling depth, as a percentage, to apply when processing objects. This
	// value determines the percentage of eligible objects that the job analyzes. If
	// the value is less than 100, Amazon Macie randomly selects the objects to
	// analyze, up to the specified percentage.
	SamplingPercentage *int32
}

type CreateClassificationJobOutput struct {
	// The Amazon Resource Name (ARN) of the job.
	JobArn *string
	// The unique identifier for the job.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateClassificationJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateClassificationJob{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateClassificationJob{}, middleware.After)
}

type idempotencyToken_initializeOpCreateClassificationJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateClassificationJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateClassificationJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateClassificationJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateClassificationJobInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateClassificationJobMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateClassificationJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateClassificationJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "CreateClassificationJob",
	}
}
