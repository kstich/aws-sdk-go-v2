// Code generated by smithy-go-codegen DO NOT EDIT.

package codeguruprofiler

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeguruprofiler/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes a profiling group.
func (c *Client) DescribeProfilingGroup(ctx context.Context, params *DescribeProfilingGroupInput, optFns ...func(*Options)) (*DescribeProfilingGroupOutput, error) {
	stack := middleware.NewStack("DescribeProfilingGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeProfilingGroupMiddlewares(stack)
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
	addOpDescribeProfilingGroupValidationMiddleware(stack)

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
			OperationName: "DescribeProfilingGroup",
			Err:           err,
		}
	}
	out := result.(*DescribeProfilingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The structure representing the describeProfilingGroupRequest.
type DescribeProfilingGroupInput struct {
	// The profiling group name.
	ProfilingGroupName *string
}

// The structure representing the describeProfilingGroupResponse.
type DescribeProfilingGroupOutput struct {
	// Information about a profiling group.
	ProfilingGroup *types.ProfilingGroupDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeProfilingGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeProfilingGroup{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeProfilingGroup{}, middleware.After)
}
