// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the description of specific Amazon FSx for Lustre data repository tasks,
// if one or more TaskIds values are provided in the request, or if filters are
// used in the request. You can use filters to narrow the response to include just
// tasks for specific file systems, or tasks in a specific lifecycle state.
// Otherwise, it returns all data repository tasks owned by your AWS account in the
// AWS Region of the endpoint that you're calling.  <p>When retrieving all tasks,
// you can paginate the response by using the optional <code>MaxResults</code>
// parameter to limit the number of tasks returned in a response. If more tasks
// remain, Amazon FSx returns a <code>NextToken</code> value in the response. In
// this case, send a later request with the <code>NextToken</code> request
// parameter set to the value of <code>NextToken</code> from the last response.</p>
func (c *Client) DescribeDataRepositoryTasks(ctx context.Context, params *DescribeDataRepositoryTasksInput, optFns ...func(*Options)) (*DescribeDataRepositoryTasksOutput, error) {
	stack := middleware.NewStack("DescribeDataRepositoryTasks", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeDataRepositoryTasksMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDataRepositoryTasks(options.Region), middleware.Before)
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
			OperationName: "DescribeDataRepositoryTasks",
			Err:           err,
		}
	}
	out := result.(*DescribeDataRepositoryTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDataRepositoryTasksInput struct {
	// (Optional) Opaque pagination token returned from a previous operation (String).
	// If present, this token indicates from what point you can continue processing the
	// request, where the previous NextToken value left off.
	NextToken *string
	// (Optional) You can use filters to narrow the DescribeDataRepositoryTasks
	// response to include just tasks for specific file systems, or tasks in a specific
	// lifecycle state.
	Filters []*types.DataRepositoryTaskFilter
	// The maximum number of resources to return in the response. This value must be an
	// integer greater than zero.
	MaxResults *int32
	// (Optional) IDs of the tasks whose descriptions you want to retrieve (String).
	TaskIds []*string
}

type DescribeDataRepositoryTasksOutput struct {
	// (Optional) Opaque pagination token returned from a previous operation (String).
	// If present, this token indicates from what point you can continue processing the
	// request, where the previous NextToken value left off.
	NextToken *string
	// The collection of data repository task descriptions returned.
	DataRepositoryTasks []*types.DataRepositoryTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeDataRepositoryTasksMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDataRepositoryTasks{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDataRepositoryTasks{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDataRepositoryTasks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fsx",
		OperationName: "DescribeDataRepositoryTasks",
	}
}
