// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the details of the specified model version.
func (c *Client) GetModelVersion(ctx context.Context, params *GetModelVersionInput, optFns ...func(*Options)) (*GetModelVersionOutput, error) {
	stack := middleware.NewStack("GetModelVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetModelVersionMiddlewares(stack)
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
	addOpGetModelVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetModelVersion(options.Region), middleware.Before)

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
			OperationName: "GetModelVersion",
			Err:           err,
		}
	}
	out := result.(*GetModelVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetModelVersionInput struct {
	// The model ID.
	ModelId *string
	// The model type.
	ModelType types.ModelTypeEnum
	// The model version number.
	ModelVersionNumber *string
}

type GetModelVersionOutput struct {
	// The model version ARN.
	Arn *string
	// The training data source.
	TrainingDataSource types.TrainingDataSourceEnum
	// The training data schema.
	TrainingDataSchema *types.TrainingDataSchema
	// The model version status.
	Status *string
	// The model ID.
	ModelId *string
	// The model version number.
	ModelVersionNumber *string
	// The event details.
	ExternalEventsDetail *types.ExternalEventsDetail
	// The model type.
	ModelType types.ModelTypeEnum

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetModelVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetModelVersion{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetModelVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetModelVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "GetModelVersion",
	}
}
