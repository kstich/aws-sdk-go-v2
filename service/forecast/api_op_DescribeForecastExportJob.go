// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes a forecast export job created using the CreateForecastExportJob ()
// operation. In addition to listing the properties provided by the user in the
// CreateForecastExportJob request, this operation lists the following
// properties:
//
//     * CreationTime
//
//     * LastModificationTime
//
//     * Status
//
//     *
// Message - If an error occurred, information about the error.
func (c *Client) DescribeForecastExportJob(ctx context.Context, params *DescribeForecastExportJobInput, optFns ...func(*Options)) (*DescribeForecastExportJobOutput, error) {
	stack := middleware.NewStack("DescribeForecastExportJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeForecastExportJobMiddlewares(stack)
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
	addOpDescribeForecastExportJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeForecastExportJob(options.Region), middleware.Before)

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
			OperationName: "DescribeForecastExportJob",
			Err:           err,
		}
	}
	out := result.(*DescribeForecastExportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeForecastExportJobInput struct {
	// The Amazon Resource Name (ARN) of the forecast export job.
	ForecastExportJobArn *string
}

type DescribeForecastExportJobOutput struct {
	// The Amazon Resource Name (ARN) of the exported forecast.
	ForecastArn *string
	// When the forecast export job was created.
	CreationTime *time.Time
	// The name of the forecast export job.
	ForecastExportJobName *string
	// If an error occurred, an informational message about the error.
	Message *string
	// The path to the Amazon Simple Storage Service (Amazon S3) bucket where the
	// forecast is exported.
	Destination *types.DataDestination
	// When the last successful export job finished.
	LastModificationTime *time.Time
	// The ARN of the forecast export job.
	ForecastExportJobArn *string
	// The status of the forecast export job. States include:
	//
	//     * ACTIVE
	//
	//     *
	// CREATE_PENDING, CREATE_IN_PROGRESS, CREATE_FAILED
	//
	//     * DELETE_PENDING,
	// DELETE_IN_PROGRESS, DELETE_FAILED
	//
	// The Status of the forecast export job must be
	// ACTIVE before you can access the forecast in your S3 bucket.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeForecastExportJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeForecastExportJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeForecastExportJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeForecastExportJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "DescribeForecastExportJob",
	}
}
