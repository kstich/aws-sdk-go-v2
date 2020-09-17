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

// Lists all exported output values in the account and Region in which you call
// this action. Use this action to see the exported output values that you can
// import into other stacks. To import values, use the Fn::ImportValue
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-importvalue.html)
// function. For more information, see  AWS CloudFormation Export Stack Output
// Values
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-exports.html).
func (c *Client) ListExports(ctx context.Context, params *ListExportsInput, optFns ...func(*Options)) (*ListExportsOutput, error) {
	stack := middleware.NewStack("ListExports", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpListExportsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListExports(options.Region), middleware.Before)
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
			OperationName: "ListExports",
			Err:           err,
		}
	}
	out := result.(*ListExportsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListExportsInput struct {
	// A string (provided by the ListExports () response output) that identifies the
	// next page of exported output values that you asked to retrieve.
	NextToken *string
}

type ListExportsOutput struct {
	// The output for the ListExports () action.
	Exports []*types.Export
	// If the output exceeds 100 exported output values, a string that identifies the
	// next page of exports. If there is no additional page, this value is null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpListExportsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpListExports{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpListExports{}, middleware.After)
}

func newServiceMetadataMiddleware_opListExports(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "ListExports",
	}
}
