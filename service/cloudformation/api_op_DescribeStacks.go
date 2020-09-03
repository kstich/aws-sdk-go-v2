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

// Returns the description for the specified stack; if no stack name was specified,
// then it returns the description for all the stacks created. If the stack does
// not exist, an AmazonCloudFormationException is returned.
func (c *Client) DescribeStacks(ctx context.Context, params *DescribeStacksInput, optFns ...func(*Options)) (*DescribeStacksOutput, error) {
	stack := middleware.NewStack("DescribeStacks", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeStacksMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStacks(options.Region), middleware.Before)

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
			OperationName: "DescribeStacks",
			Err:           err,
		}
	}
	out := result.(*DescribeStacksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for DescribeStacks () action.
type DescribeStacksInput struct {
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
	// A string that identifies the next page of stacks that you want to retrieve.
	NextToken *string
}

// The output for a DescribeStacks () action.
type DescribeStacksOutput struct {
	// If the output exceeds 1 MB in size, a string that identifies the next page of
	// stacks. If no additional page exists, this value is null.
	NextToken *string
	// A list of stack structures.
	Stacks []*types.Stack

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeStacksMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeStacks{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeStacks{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeStacks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DescribeStacks",
	}
}
