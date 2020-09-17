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

// Returns information about provisioned Aurora DB clusters. This API supports
// pagination. For more information on Amazon Aurora, see  What Is Amazon Aurora?
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. This operation can also return information for
// Amazon Neptune DB instances and Amazon DocumentDB instances.
func (c *Client) DescribeDBClusters(ctx context.Context, params *DescribeDBClustersInput, optFns ...func(*Options)) (*DescribeDBClustersOutput, error) {
	stack := middleware.NewStack("DescribeDBClusters", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeDBClustersMiddlewares(stack)
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
	addOpDescribeDBClustersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBClusters(options.Region), middleware.Before)
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
			OperationName: "DescribeDBClusters",
			Err:           err,
		}
	}
	out := result.(*DescribeDBClustersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeDBClustersInput struct {
	// A filter that specifies one or more DB clusters to describe. Supported
	// filters:
	//
	//     * db-cluster-id - Accepts DB cluster identifiers and DB cluster
	// Amazon Resource Names (ARNs). The results list will only include information
	// about the DB clusters identified by these ARNs.
	Filters []*types.Filter
	// Optional Boolean parameter that specifies whether the output includes
	// information about clusters shared from other AWS accounts.
	IncludeShared *bool
	// The user-supplied DB cluster identifier. If this parameter is specified,
	// information from only the specific DB cluster is returned. This parameter isn't
	// case-sensitive. Constraints:
	//
	//     * If supplied, must match an existing
	// DBClusterIdentifier.
	DBClusterIdentifier *string
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
	// An optional pagination token provided by a previous DescribeDBClusters request.
	// If this parameter is specified, the response includes only records beyond the
	// marker, up to the value specified by MaxRecords.
	Marker *string
}

// Contains the result of a successful invocation of the DescribeDBClusters action.
type DescribeDBClustersOutput struct {
	// Contains a list of DB clusters for the user.
	DBClusters []*types.DBCluster
	// A pagination token that can be used in a later DescribeDBClusters request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeDBClustersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBClusters{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBClusters{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDBClusters(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBClusters",
	}
}
