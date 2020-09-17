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

// Returns a paginated list of individual assessments based on filter settings.
// These filter settings can specify a combination of premigration assessment runs,
// migration tasks, and assessment status values.
func (c *Client) DescribeReplicationTaskIndividualAssessments(ctx context.Context, params *DescribeReplicationTaskIndividualAssessmentsInput, optFns ...func(*Options)) (*DescribeReplicationTaskIndividualAssessmentsOutput, error) {
	stack := middleware.NewStack("DescribeReplicationTaskIndividualAssessments", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeReplicationTaskIndividualAssessmentsMiddlewares(stack)
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
	addOpDescribeReplicationTaskIndividualAssessmentsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReplicationTaskIndividualAssessments(options.Region), middleware.Before)
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
			OperationName: "DescribeReplicationTaskIndividualAssessments",
			Err:           err,
		}
	}
	out := result.(*DescribeReplicationTaskIndividualAssessmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeReplicationTaskIndividualAssessmentsInput struct {
	// Filters applied to the individual assessments described in the form of key-value
	// pairs. Valid filter names: replication-task-assessment-run-arn,
	// replication-task-arn, status
	Filters []*types.Filter
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	MaxRecords *int32
	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string
}

//
type DescribeReplicationTaskIndividualAssessmentsOutput struct {
	// One or more individual assessments as specified by Filters.
	ReplicationTaskIndividualAssessments []*types.ReplicationTaskIndividualAssessment
	// A pagination token returned for you to pass to a subsequent request. If you pass
	// this token as the Marker value in a subsequent request, the response includes
	// only records beyond the marker, up to the value specified in the request by
	// MaxRecords.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeReplicationTaskIndividualAssessmentsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeReplicationTaskIndividualAssessments{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeReplicationTaskIndividualAssessments{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeReplicationTaskIndividualAssessments(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "DescribeReplicationTaskIndividualAssessments",
	}
}
