// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a job that imports training data from your data source (an Amazon S3
// bucket) to an Amazon Personalize dataset. To allow Amazon Personalize to import
// the training data, you must specify an AWS Identity and Access Management (IAM)
// role that has permission to read from the data source. The dataset import job
// replaces any previous data in the dataset. Status A dataset import job can be in
// one of the following states:
//
//     * CREATE PENDING > CREATE IN_PROGRESS > ACTIVE
// -or- CREATE FAILED
//
// To get the status of the import job, call
// DescribeDatasetImportJob (), providing the Amazon Resource Name (ARN) of the
// dataset import job. The dataset import is complete when the status shows as
// ACTIVE. If the status shows as CREATE FAILED, the response includes a
// failureReason key, which describes why the job failed. Importing takes time. You
// must wait until the status shows as ACTIVE before training a model using the
// dataset.  <p class="title"> <b>Related APIs</b> </p> <ul> <li> <p>
// <a>ListDatasetImportJobs</a> </p> </li> <li> <p> <a>DescribeDatasetImportJob</a>
// </p> </li> </ul>
func (c *Client) CreateDatasetImportJob(ctx context.Context, params *CreateDatasetImportJobInput, optFns ...func(*Options)) (*CreateDatasetImportJobOutput, error) {
	stack := middleware.NewStack("CreateDatasetImportJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateDatasetImportJobMiddlewares(stack)
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
	addOpCreateDatasetImportJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDatasetImportJob(options.Region), middleware.Before)

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
			OperationName: "CreateDatasetImportJob",
			Err:           err,
		}
	}
	out := result.(*CreateDatasetImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDatasetImportJobInput struct {
	// The Amazon S3 bucket that contains the training data to import.
	DataSource *types.DataSource
	// The name for the dataset import job.
	JobName *string
	// The ARN of the IAM role that has permissions to read from the Amazon S3 data
	// source.
	RoleArn *string
	// The ARN of the dataset that receives the imported data.
	DatasetArn *string
}

type CreateDatasetImportJobOutput struct {
	// The ARN of the dataset import job.
	DatasetImportJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateDatasetImportJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDatasetImportJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDatasetImportJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDatasetImportJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "CreateDatasetImportJob",
	}
}
