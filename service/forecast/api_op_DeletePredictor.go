// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a predictor created using the CreatePredictor () operation. You can
// delete only predictor that have a status of ACTIVE or CREATE_FAILED. To get the
// status, use the DescribePredictor () operation.
func (c *Client) DeletePredictor(ctx context.Context, params *DeletePredictorInput, optFns ...func(*Options)) (*DeletePredictorOutput, error) {
	stack := middleware.NewStack("DeletePredictor", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeletePredictorMiddlewares(stack)
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
	addOpDeletePredictorValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePredictor(options.Region), middleware.Before)
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
			OperationName: "DeletePredictor",
			Err:           err,
		}
	}
	out := result.(*DeletePredictorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePredictorInput struct {
	// The Amazon Resource Name (ARN) of the predictor to delete.
	PredictorArn *string
}

type DeletePredictorOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeletePredictorMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeletePredictor{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeletePredictor{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeletePredictor(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "DeletePredictor",
	}
}
