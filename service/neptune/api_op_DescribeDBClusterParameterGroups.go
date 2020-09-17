// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of DBClusterParameterGroup descriptions. If a
// DBClusterParameterGroupName parameter is specified, the list will contain only
// the description of the specified DB cluster parameter group.
func (c *Client) DescribeDBClusterParameterGroups(ctx context.Context, params *DescribeDBClusterParameterGroupsInput, optFns ...func(*Options)) (*DescribeDBClusterParameterGroupsOutput, error) {
	stack := middleware.NewStack("DescribeDBClusterParameterGroups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeDBClusterParameterGroupsMiddlewares(stack)
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
	addOpDescribeDBClusterParameterGroupsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBClusterParameterGroups(options.Region), middleware.Before)
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
			OperationName: "DescribeDBClusterParameterGroups",
			Err:           err,
		}
	}
	out := result.(*DescribeDBClusterParameterGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDBClusterParameterGroupsInput struct {
	// The name of a specific DB cluster parameter group to return details for.
	// Constraints:
	//
	//     * If supplied, must match the name of an existing
	// DBClusterParameterGroup.
	DBClusterParameterGroupName *string
	// An optional pagination token provided by a previous
	// DescribeDBClusterParameterGroups request. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string
	// This parameter is not currently supported.
	Filters []*types.Filter
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
}

type DescribeDBClusterParameterGroupsOutput struct {
	// An optional pagination token provided by a previous
	// DescribeDBClusterParameterGroups request. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string
	// A list of DB cluster parameter groups.
	DBClusterParameterGroups []*types.DBClusterParameterGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeDBClusterParameterGroupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBClusterParameterGroups{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBClusterParameterGroups{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDBClusterParameterGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBClusterParameterGroups",
	}
}
