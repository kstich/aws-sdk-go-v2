// Code generated by smithy-go-codegen DO NOT EDIT.

package codegurureviewer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codegurureviewer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the metadata associated with the code review along with its status.
func (c *Client) DescribeCodeReview(ctx context.Context, params *DescribeCodeReviewInput, optFns ...func(*Options)) (*DescribeCodeReviewOutput, error) {
	stack := middleware.NewStack("DescribeCodeReview", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeCodeReviewMiddlewares(stack)
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
	addOpDescribeCodeReviewValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCodeReview(options.Region), middleware.Before)
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
			OperationName: "DescribeCodeReview",
			Err:           err,
		}
	}
	out := result.(*DescribeCodeReviewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCodeReviewInput struct {
	// The Amazon Resource Name (ARN) of the CodeReview
	// (https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_CodeReview.html)
	// object.
	CodeReviewArn *string
}

type DescribeCodeReviewOutput struct {
	// Information about the code review.
	CodeReview *types.CodeReview

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeCodeReviewMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeCodeReview{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeCodeReview{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeCodeReview(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-reviewer",
		OperationName: "DescribeCodeReview",
	}
}
