// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts a new premigration assessment run for one or more individual assessments
// of a migration task. The assessments that you can specify depend on the source
// and target database engine and the migration type defined for the given task. To
// run this operation, your migration task must already be created. After you run
// this operation, you can review the status of each individual assessment. You can
// also run the migration task manually after the assessment run and its individual
// assessments complete.
func (c *Client) StartReplicationTaskAssessmentRun(ctx context.Context, params *StartReplicationTaskAssessmentRunInput, optFns ...func(*Options)) (*StartReplicationTaskAssessmentRunOutput, error) {
	stack := middleware.NewStack("StartReplicationTaskAssessmentRun", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartReplicationTaskAssessmentRunMiddlewares(stack)
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
	addOpStartReplicationTaskAssessmentRunValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartReplicationTaskAssessmentRun(options.Region), middleware.Before)
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
			OperationName: "StartReplicationTaskAssessmentRun",
			Err:           err,
		}
	}
	out := result.(*StartReplicationTaskAssessmentRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type StartReplicationTaskAssessmentRunInput struct {
	// ARN of a service role needed to start the assessment run.
	ServiceAccessRoleArn *string
	// Unique name to identify the assessment run.
	AssessmentRunName *string
	// Folder within an Amazon S3 bucket where you want AWS DMS to store the results of
	// this assessment run.
	ResultLocationFolder *string
	// Encryption mode that you can specify to encrypt the results of this assessment
	// run. If you don't specify this request parameter, AWS DMS stores the assessment
	// run results without encryption. You can specify one of the options following:
	//
	//
	// * "SSE_S3" – The server-side encryption provided as a default by Amazon S3.
	//
	//
	// * "SSE_KMS" – AWS Key Management Service (AWS KMS) encryption. This encryption
	// can use either a custom KMS encryption key that you specify or the default KMS
	// encryption key that DMS provides.
	ResultEncryptionMode *string
	// Space-separated list of names for specific individual assessments that you want
	// to include. These names come from the default list of individual assessments
	// that AWS DMS supports for the associated migration task. This task is specified
	// by ReplicationTaskArn. You can't set a value for IncludeOnly if you also set a
	// value for Exclude in the API operation. To identify the names of the default
	// individual assessments that AWS DMS supports for the associated migration task,
	// run the DescribeApplicableIndividualAssessments operation using its own
	// ReplicationTaskArn request parameter.
	IncludeOnly []*string
	// Space-separated list of names for specific individual assessments that you want
	// to exclude. These names come from the default list of individual assessments
	// that AWS DMS supports for the associated migration task. This task is specified
	// by ReplicationTaskArn. You can't set a value for Exclude if you also set a value
	// for IncludeOnly in the API operation. To identify the names of the default
	// individual assessments that AWS DMS supports for the associated migration task,
	// run the DescribeApplicableIndividualAssessments operation using its own
	// ReplicationTaskArn request parameter.
	Exclude []*string
	// Amazon S3 bucket where you want AWS DMS to store the results of this assessment
	// run.
	ResultLocationBucket *string
	// Amazon Resource Name (ARN) of the migration task associated with the
	// premigration assessment run that you want to start.
	ReplicationTaskArn *string
	// ARN of a custom KMS encryption key that you specify when you set
	// ResultEncryptionMode to "SSE_KMS".
	ResultKmsKeyArn *string
}

//
type StartReplicationTaskAssessmentRunOutput struct {
	// The premigration assessment run that was started.
	ReplicationTaskAssessmentRun *types.ReplicationTaskAssessmentRun

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartReplicationTaskAssessmentRunMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartReplicationTaskAssessmentRun{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartReplicationTaskAssessmentRun{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartReplicationTaskAssessmentRun(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "StartReplicationTaskAssessmentRun",
	}
}
