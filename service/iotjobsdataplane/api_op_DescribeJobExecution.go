// Code generated by smithy-go-codegen DO NOT EDIT.

package iotjobsdataplane

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotjobsdataplane/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets details of a job execution.
func (c *Client) DescribeJobExecution(ctx context.Context, params *DescribeJobExecutionInput, optFns ...func(*Options)) (*DescribeJobExecutionOutput, error) {
	stack := middleware.NewStack("DescribeJobExecution", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeJobExecutionMiddlewares(stack)
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
	addOpDescribeJobExecutionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeJobExecution(options.Region), middleware.Before)
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
			OperationName: "DescribeJobExecution",
			Err:           err,
		}
	}
	out := result.(*DescribeJobExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeJobExecutionInput struct {
	// The unique identifier assigned to this job when it was created.
	JobId *string
	// Optional. A number that identifies a particular job execution on a particular
	// device. If not specified, the latest job execution is returned.
	ExecutionNumber *int64
	// The thing name associated with the device the job execution is running on.
	ThingName *string
	// Optional. When set to true, the response contains the job document. The default
	// is false.
	IncludeJobDocument *bool
}

type DescribeJobExecutionOutput struct {
	// Contains data about a job execution.
	Execution *types.JobExecution

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeJobExecutionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeJobExecution{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeJobExecution{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeJobExecution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iot-jobs-data",
		OperationName: "DescribeJobExecution",
	}
}
