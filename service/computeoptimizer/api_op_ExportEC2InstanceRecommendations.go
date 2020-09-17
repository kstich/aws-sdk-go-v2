// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Exports optimization recommendations for Amazon EC2 instances.
// <p>Recommendations are exported in a comma-separated values (.csv) file, and its
// metadata in a JavaScript Object Notation (.json) file, to an existing Amazon
// Simple Storage Service (Amazon S3) bucket that you specify. For more
// information, see <a
// href="https://docs.aws.amazon.com/compute-optimizer/latest/ug/exporting-recommendations.html">Exporting
// Recommendations</a> in the <i>Compute Optimizer User Guide</i>.</p> <p>You can
// have only one Amazon EC2 instance export job in progress per AWS Region.</p>
func (c *Client) ExportEC2InstanceRecommendations(ctx context.Context, params *ExportEC2InstanceRecommendationsInput, optFns ...func(*Options)) (*ExportEC2InstanceRecommendationsOutput, error) {
	stack := middleware.NewStack("ExportEC2InstanceRecommendations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpExportEC2InstanceRecommendationsMiddlewares(stack)
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
	addOpExportEC2InstanceRecommendationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opExportEC2InstanceRecommendations(options.Region), middleware.Before)
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
			OperationName: "ExportEC2InstanceRecommendations",
			Err:           err,
		}
	}
	out := result.(*ExportEC2InstanceRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportEC2InstanceRecommendationsInput struct {
	// The IDs of the AWS accounts for which to export instance recommendations. If
	// your account is the master account of an organization, use this parameter to
	// specify the member accounts for which you want to export recommendations. This
	// parameter cannot be specified together with the include member accounts
	// parameter. The parameters are mutually exclusive. Recommendations for member
	// accounts are not included in the export if this parameter, or the include member
	// accounts parameter, is omitted. You can specify multiple account IDs per
	// request.
	AccountIds []*string
	// The format of the export file. The only export file format currently supported
	// is Csv.
	FileFormat types.FileFormat
	// The recommendations data to include in the export file.
	FieldsToExport []types.ExportableInstanceField
	// Indicates whether to include recommendations for resources in all member
	// accounts of the organization if your account is the master account of an
	// organization. The member accounts must also be opted in to Compute Optimizer.
	// Recommendations for member accounts of the organization are not included in the
	// export file if this parameter is omitted. Recommendations for member accounts
	// are not included in the export if this parameter, or the account IDs parameter,
	// is omitted.
	IncludeMemberAccounts *bool
	// An array of objects that describe a filter to export a more specific set of
	// instance recommendations.
	Filters []*types.Filter
	// An object to specify the destination Amazon Simple Storage Service (Amazon S3)
	// bucket name and key prefix for the export job. You must create the destination
	// Amazon S3 bucket for your recommendations export before you create the export
	// job. Compute Optimizer does not create the S3 bucket for you. After you create
	// the S3 bucket, ensure that it has the required permission policy to allow
	// Compute Optimizer to write the export file to it. If you plan to specify an
	// object prefix when you create the export job, you must include the object prefix
	// in the policy that you add to the S3 bucket. For more information, see Amazon S3
	// Bucket Policy for Compute Optimizer
	// (https://docs.aws.amazon.com/compute-optimizer/latest/ug/create-s3-bucket-policy-for-compute-optimizer.html)
	// in the Compute Optimizer user guide.
	S3DestinationConfig *types.S3DestinationConfig
}

type ExportEC2InstanceRecommendationsOutput struct {
	// The identification number of the export job. Use the
	// DescribeRecommendationExportJobs action, and specify the job ID to view the
	// status of an export job.
	JobId *string
	// An object that describes the destination Amazon S3 bucket of a recommendations
	// export file.
	S3Destination *types.S3Destination

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpExportEC2InstanceRecommendationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpExportEC2InstanceRecommendations{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpExportEC2InstanceRecommendations{}, middleware.After)
}

func newServiceMetadataMiddleware_opExportEC2InstanceRecommendations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "compute-optimizer",
		OperationName: "ExportEC2InstanceRecommendations",
	}
}
