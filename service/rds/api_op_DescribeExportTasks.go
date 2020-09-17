// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about a snapshot export to Amazon S3. This API operation
// supports pagination.
func (c *Client) DescribeExportTasks(ctx context.Context, params *DescribeExportTasksInput, optFns ...func(*Options)) (*DescribeExportTasksOutput, error) {
	stack := middleware.NewStack("DescribeExportTasks", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeExportTasksMiddlewares(stack)
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
	addOpDescribeExportTasksValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeExportTasks(options.Region), middleware.Before)
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
			OperationName: "DescribeExportTasks",
			Err:           err,
		}
	}
	out := result.(*DescribeExportTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeExportTasksInput struct {
	// An optional pagination token provided by a previous DescribeExportTasks request.
	// If you specify this parameter, the response includes only records beyond the
	// marker, up to the value specified by the MaxRecords parameter.
	Marker *string
	// The maximum number of records to include in the response. If more records exist
	// than the specified value, a pagination token called a marker is included in the
	// response. You can use the marker in a later DescribeExportTasks request to
	// retrieve the remaining results. Default: 100 Constraints: Minimum 20, maximum
	// 100.
	MaxRecords *int32
	// The Amazon Resource Name (ARN) of the snapshot exported to Amazon S3.
	SourceArn *string
	// Filters specify one or more snapshot exports to describe. The filters are
	// specified as name-value pairs that define what to include in the output.
	// Supported filters include the following:
	//
	//     * export-task-identifier - An
	// identifier for the snapshot export task.
	//
	//     * s3-bucket - The Amazon S3 bucket
	// the snapshot is exported to.
	//
	//     * source-arn - The Amazon Resource Name (ARN)
	// of the snapshot exported to Amazon S3
	//
	//     * status - The status of the export
	// task.
	Filters []*types.Filter
	// The identifier of the snapshot export task to be described.
	ExportTaskIdentifier *string
}

type DescribeExportTasksOutput struct {
	// A pagination token that can be used in a later DescribeExportTasks request. A
	// marker is used for pagination to identify the location to begin output for the
	// next response of DescribeExportTasks.
	Marker *string
	// Information about an export of a snapshot to Amazon S3.
	ExportTasks []*types.ExportTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeExportTasksMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeExportTasks{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeExportTasks{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeExportTasks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeExportTasks",
	}
}
