// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts an asynchronous sentiment detection job for a collection of documents.
// use the operation to track the status of a job.
func (c *Client) StartSentimentDetectionJob(ctx context.Context, params *StartSentimentDetectionJobInput, optFns ...func(*Options)) (*StartSentimentDetectionJobOutput, error) {
	stack := middleware.NewStack("StartSentimentDetectionJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartSentimentDetectionJobMiddlewares(stack)
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
	addIdempotencyToken_opStartSentimentDetectionJobMiddleware(stack, options)
	addOpStartSentimentDetectionJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartSentimentDetectionJob(options.Region), middleware.Before)
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
			OperationName: "StartSentimentDetectionJob",
			Err:           err,
		}
	}
	out := result.(*StartSentimentDetectionJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartSentimentDetectionJobInput struct {
	// Specifies the format and location of the input data for the job.
	InputDataConfig *types.InputDataConfig
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM)
	// role that grants Amazon Comprehend read access to your input data. For more
	// information, see
	// https://docs.aws.amazon.com/comprehend/latest/dg/access-control-managing-permissions.html#auth-role-permissions
	// (https://docs.aws.amazon.com/comprehend/latest/dg/access-control-managing-permissions.html#auth-role-permissions).
	DataAccessRoleArn *string
	// ID for the AWS Key Management Service (KMS) key that Amazon Comprehend uses to
	// encrypt data on the storage volume attached to the ML compute instance(s) that
	// process the analysis job. The VolumeKmsKeyId can be either of the following
	// formats:
	//
	//     * KMS Key ID: "1234abcd-12ab-34cd-56ef-1234567890ab"
	//
	//     * Amazon
	// Resource Name (ARN) of a KMS Key:
	// "arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"
	VolumeKmsKeyId *string
	// Specifies where to send the output files.
	OutputDataConfig *types.OutputDataConfig
	// Configuration parameters for an optional private Virtual Private Cloud (VPC)
	// containing the resources you are using for your sentiment detection job. For
	// more information, see Amazon VPC
	// (https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html).
	VpcConfig *types.VpcConfig
	// A unique identifier for the request. If you don't set the client request token,
	// Amazon Comprehend generates one.
	ClientRequestToken *string
	// The language of the input documents. You can specify any of the primary
	// languages supported by Amazon Comprehend. All documents must be in the same
	// language.
	LanguageCode types.LanguageCode
	// The identifier of the job.
	JobName *string
}

type StartSentimentDetectionJobOutput struct {
	// The identifier generated for the job. To get the status of a job, use this
	// identifier with the operation.
	JobId *string
	// The status of the job.
	//
	//     * SUBMITTED - The job has been received and is
	// queued for processing.
	//
	//     * IN_PROGRESS - Amazon Comprehend is processing the
	// job.
	//
	//     * COMPLETED - The job was successfully completed and the output is
	// available.
	//
	//     * FAILED - The job did not complete. To get details, use the
	// operation.
	JobStatus types.JobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartSentimentDetectionJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartSentimentDetectionJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartSentimentDetectionJob{}, middleware.After)
}

type idempotencyToken_initializeOpStartSentimentDetectionJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartSentimentDetectionJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartSentimentDetectionJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartSentimentDetectionJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartSentimentDetectionJobInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartSentimentDetectionJobMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpStartSentimentDetectionJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartSentimentDetectionJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "StartSentimentDetectionJob",
	}
}
