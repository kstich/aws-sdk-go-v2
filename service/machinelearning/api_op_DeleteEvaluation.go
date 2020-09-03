// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Assigns the DELETED status to an Evaluation, rendering it unusable.  <p>After
// invoking the <code>DeleteEvaluation</code> operation, you can use the
// <code>GetEvaluation</code> operation to verify that the status of the
// <code>Evaluation</code> changed to <code>DELETED</code>.</p> <p> <b>Caution:</b>
// The results of the <code>DeleteEvaluation</code> operation are irreversible.</p>
func (c *Client) DeleteEvaluation(ctx context.Context, params *DeleteEvaluationInput, optFns ...func(*Options)) (*DeleteEvaluationOutput, error) {
	stack := middleware.NewStack("DeleteEvaluation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteEvaluationMiddlewares(stack)
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
	addOpDeleteEvaluationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteEvaluation(options.Region), middleware.Before)

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
			OperationName: "DeleteEvaluation",
			Err:           err,
		}
	}
	out := result.(*DeleteEvaluationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteEvaluationInput struct {
	// A user-supplied ID that uniquely identifies the Evaluation to delete.
	EvaluationId *string
}

// Represents the output of a DeleteEvaluation operation. The output indicates that
// Amazon Machine Learning (Amazon ML) received the request. You can use the
// GetEvaluation operation and check the value of the Status parameter to see
// whether an Evaluation is marked as DELETED.
type DeleteEvaluationOutput struct {
	// A user-supplied ID that uniquely identifies the Evaluation. This value should be
	// identical to the value of the EvaluationId in the request.
	EvaluationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteEvaluationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteEvaluation{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteEvaluation{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteEvaluation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "machinelearning",
		OperationName: "DeleteEvaluation",
	}
}
