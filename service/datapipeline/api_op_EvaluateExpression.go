// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Task runners call EvaluateExpression to evaluate a string in the context of the
// specified object. For example, a task runner can evaluate SQL queries stored in
// Amazon S3.
func (c *Client) EvaluateExpression(ctx context.Context, params *EvaluateExpressionInput, optFns ...func(*Options)) (*EvaluateExpressionOutput, error) {
	stack := middleware.NewStack("EvaluateExpression", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpEvaluateExpressionMiddlewares(stack)
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
	addOpEvaluateExpressionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opEvaluateExpression(options.Region), middleware.Before)

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
			OperationName: "EvaluateExpression",
			Err:           err,
		}
	}
	out := result.(*EvaluateExpressionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for EvaluateExpression.
type EvaluateExpressionInput struct {
	// The ID of the pipeline.
	PipelineId *string
	// The expression to evaluate.
	Expression *string
	// The ID of the object.
	ObjectId *string
}

// Contains the output of EvaluateExpression.
type EvaluateExpressionOutput struct {
	// The evaluated expression.
	EvaluatedExpression *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpEvaluateExpressionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpEvaluateExpression{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpEvaluateExpression{}, middleware.After)
}

func newServiceMetadataMiddleware_opEvaluateExpression(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datapipeline",
		OperationName: "EvaluateExpression",
	}
}
