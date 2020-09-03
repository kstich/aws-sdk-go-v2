// Code generated by smithy-go-codegen DO NOT EDIT.

package iot1clickprojects

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot1clickprojects/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns an object describing a project.
func (c *Client) DescribeProject(ctx context.Context, params *DescribeProjectInput, optFns ...func(*Options)) (*DescribeProjectOutput, error) {
	stack := middleware.NewStack("DescribeProject", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeProjectMiddlewares(stack)
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
	addOpDescribeProjectValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeProject(options.Region), middleware.Before)

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
			OperationName: "DescribeProject",
			Err:           err,
		}
	}
	out := result.(*DescribeProjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeProjectInput struct {
	// The name of the project to be described.
	ProjectName *string
}

type DescribeProjectOutput struct {
	// An object describing the project.
	Project *types.ProjectDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeProjectMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeProject{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeProject{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeProject(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iot1click",
		OperationName: "DescribeProject",
	}
}
