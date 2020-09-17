// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a description of the specified resource in the specified stack. For
// deleted stacks, DescribeStackResource returns resource information for up to 90
// days after the stack has been deleted.
func (c *Client) DescribeStackResource(ctx context.Context, params *DescribeStackResourceInput, optFns ...func(*Options)) (*DescribeStackResourceOutput, error) {
	stack := middleware.NewStack("DescribeStackResource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeStackResourceMiddlewares(stack)
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
	addOpDescribeStackResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStackResource(options.Region), middleware.Before)
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
			OperationName: "DescribeStackResource",
			Err:           err,
		}
	}
	out := result.(*DescribeStackResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for DescribeStackResource () action.
type DescribeStackResourceInput struct {
	// The logical name of the resource as specified in the template. Default: There is
	// no default value.
	LogicalResourceId *string
	// The name or the unique stack ID that is associated with the stack, which are not
	// always interchangeable:
	//
	//     * Running stacks: You can specify either the
	// stack's name or its unique stack ID.
	//
	//     * Deleted stacks: You must specify the
	// unique stack ID.
	//
	// Default: There is no default value.
	StackName *string
}

// The output for a DescribeStackResource () action.
type DescribeStackResourceOutput struct {
	// A StackResourceDetail structure containing the description of the specified
	// resource in the specified stack.
	StackResourceDetail *types.StackResourceDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeStackResourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeStackResource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeStackResource{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeStackResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DescribeStackResource",
	}
}
