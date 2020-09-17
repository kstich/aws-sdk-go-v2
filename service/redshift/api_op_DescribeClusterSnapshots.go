// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns one or more snapshot objects, which contain metadata about your cluster
// snapshots. By default, this operation returns information about all snapshots of
// all clusters that are owned by you AWS customer account. No information is
// returned for snapshots owned by inactive AWS customer accounts. If you specify
// both tag keys and tag values in the same request, Amazon Redshift returns all
// snapshots that match any combination of the specified keys and values. For
// example, if you have owner and environment for tag keys, and admin and test for
// tag values, all snapshots that have any combination of those values are
// returned. Only snapshots that you own are returned in the response; shared
// snapshots are not returned with the tag key and tag value request parameters. If
// both tag keys and values are omitted from the request, snapshots are returned
// regardless of whether they have tag keys or values associated with them.
func (c *Client) DescribeClusterSnapshots(ctx context.Context, params *DescribeClusterSnapshotsInput, optFns ...func(*Options)) (*DescribeClusterSnapshotsOutput, error) {
	stack := middleware.NewStack("DescribeClusterSnapshots", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeClusterSnapshotsMiddlewares(stack)
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
	addOpDescribeClusterSnapshotsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClusterSnapshots(options.Region), middleware.Before)
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
			OperationName: "DescribeClusterSnapshots",
			Err:           err,
		}
	}
	out := result.(*DescribeClusterSnapshotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeClusterSnapshotsInput struct {
	// The identifier of the cluster which generated the requested snapshots.
	ClusterIdentifier *string
	// The snapshot identifier of the snapshot about which to return information.
	SnapshotIdentifier *string
	// The type of snapshots for which you are requesting information. By default,
	// snapshots of all types are returned. Valid Values: automated | manual
	SnapshotType *string
	// A value that requests only snapshots created at or after the specified time. The
	// time value is specified in ISO 8601 format. For more information about ISO 8601,
	// go to the ISO8601 Wikipedia page. (http://en.wikipedia.org/wiki/ISO_8601)
	// Example: 2012-07-16T18:00:00Z
	StartTime *time.Time
	// A time value that requests only snapshots created at or before the specified
	// time. The time value is specified in ISO 8601 format. For more information about
	// ISO 8601, go to the ISO8601 Wikipedia page.
	// (http://en.wikipedia.org/wiki/ISO_8601) Example: 2012-07-16T18:00:00Z
	EndTime *time.Time
	// A tag key or keys for which you want to return all matching cluster snapshots
	// that are associated with the specified key or keys. For example, suppose that
	// you have snapshots that are tagged with keys called owner and environment. If
	// you specify both of these tag keys in the request, Amazon Redshift returns a
	// response with the snapshots that have either or both of these tag keys
	// associated with them.
	TagKeys []*string
	//
	SortingEntities []*types.SnapshotSortingEntity
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int32
	// A value that indicates whether to return snapshots only for an existing cluster.
	// You can perform table-level restore only by using a snapshot of an existing
	// cluster, that is, a cluster that has not been deleted. Values for this parameter
	// work as follows:
	//
	//     * If ClusterExists is set to true, ClusterIdentifier is
	// required.
	//
	//     * If ClusterExists is set to false and ClusterIdentifier isn't
	// specified, all snapshots associated with deleted clusters (orphaned snapshots)
	// are returned.
	//
	//     * If ClusterExists is set to false and ClusterIdentifier is
	// specified for a deleted cluster, snapshots associated with that cluster are
	// returned.
	//
	//     * If ClusterExists is set to false and ClusterIdentifier is
	// specified for an existing cluster, no snapshots are returned.
	ClusterExists *bool
	// The AWS customer account used to create or copy the snapshot. Use this field to
	// filter the results to snapshots owned by a particular account. To describe
	// snapshots you own, either specify your AWS customer account, or do not specify
	// the parameter.
	OwnerAccount *string
	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeClusterSnapshots () request
	// exceed the value specified in MaxRecords, AWS returns a value in the Marker
	// field of the response. You can retrieve the next set of response records by
	// providing the returned marker value in the Marker parameter and retrying the
	// request.
	Marker *string
	// A tag value or values for which you want to return all matching cluster
	// snapshots that are associated with the specified tag value or values. For
	// example, suppose that you have snapshots that are tagged with values called
	// admin and test. If you specify both of these tag values in the request, Amazon
	// Redshift returns a response with the snapshots that have either or both of these
	// tag values associated with them.
	TagValues []*string
}

// Contains the output from the DescribeClusterSnapshots () action.
type DescribeClusterSnapshotsOutput struct {
	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string
	// A list of Snapshot () instances.
	Snapshots []*types.Snapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeClusterSnapshotsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeClusterSnapshots{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeClusterSnapshots{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeClusterSnapshots(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeClusterSnapshots",
	}
}
