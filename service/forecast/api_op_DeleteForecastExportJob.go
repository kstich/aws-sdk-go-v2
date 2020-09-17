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

// Deletes a forecast export job created using the CreateForecastExportJob ()
// operation. You can delete only export jobs that have a status of ACTIVE or
// CREATE_FAILED. To get the status, use the DescribeForecastExportJob ()
// operation.
func (c *Client) DeleteForecastExportJob(ctx context.Context, params *DeleteForecastExportJobInput, optFns ...func(*Options)) (*DeleteForecastExportJobOutput, error) {
	stack := middleware.NewStack("DeleteForecastExportJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteForecastExportJobMiddlewares(stack)
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
	addOpDeleteForecastExportJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteForecastExportJob(options.Region), middleware.Before)
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
			OperationName: "DeleteForecastExportJob",
			Err:           err,
		}
	}
	out := result.(*DeleteForecastExportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteForecastExportJobInput struct {
	// The Amazon Resource Name (ARN) of the forecast export job to delete.
	ForecastExportJobArn *string
}

type DeleteForecastExportJobOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteForecastExportJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteForecastExportJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteForecastExportJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteForecastExportJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "DeleteForecastExportJob",
	}
}
