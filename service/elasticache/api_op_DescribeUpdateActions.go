// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns details of the update actions
func (c *Client) DescribeUpdateActions(ctx context.Context, params *DescribeUpdateActionsInput, optFns ...func(*Options)) (*DescribeUpdateActionsOutput, error) {
	stack := middleware.NewStack("DescribeUpdateActions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeUpdateActionsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeUpdateActions(options.Region), middleware.Before)
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
			OperationName: "DescribeUpdateActions",
			Err:           err,
		}
	}
	out := result.(*DescribeUpdateActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeUpdateActionsInput struct {
	// Dictates whether to include node level update status in the response
	ShowNodeLevelUpdateStatus *bool
	// The Elasticache engine to which the update applies. Either Redis or Memcached
	Engine *string
	// The status of the update action.
	UpdateActionStatus []types.UpdateActionStatus
	// The cache cluster IDs
	CacheClusterIds []*string
	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string
	// The status of the service update
	ServiceUpdateStatus []types.ServiceUpdateStatus
	// The maximum number of records to include in the response
	MaxRecords *int32
	// The replication group IDs
	ReplicationGroupIds []*string
	// The unique ID of the service update
	ServiceUpdateName *string
	// The range of time specified to search for service updates that are in available
	// status
	ServiceUpdateTimeRange *types.TimeRangeFilter
}

type DescribeUpdateActionsOutput struct {
	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string
	// Returns a list of update actions
	UpdateActions []*types.UpdateAction

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeUpdateActionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeUpdateActions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeUpdateActions{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeUpdateActions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DescribeUpdateActions",
	}
}
