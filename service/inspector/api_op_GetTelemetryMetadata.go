// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Information about the data that is collected for the specified assessment run.
func (c *Client) GetTelemetryMetadata(ctx context.Context, params *GetTelemetryMetadataInput, optFns ...func(*Options)) (*GetTelemetryMetadataOutput, error) {
	stack := middleware.NewStack("GetTelemetryMetadata", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetTelemetryMetadataMiddlewares(stack)
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
	addOpGetTelemetryMetadataValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTelemetryMetadata(options.Region), middleware.Before)

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
			OperationName: "GetTelemetryMetadata",
			Err:           err,
		}
	}
	out := result.(*GetTelemetryMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTelemetryMetadataInput struct {
	// The ARN that specifies the assessment run that has the telemetry data that you
	// want to obtain.
	AssessmentRunArn *string
}

type GetTelemetryMetadataOutput struct {
	// Telemetry details.
	TelemetryMetadata []*types.TelemetryMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetTelemetryMetadataMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetTelemetryMetadata{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetTelemetryMetadata{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetTelemetryMetadata(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "GetTelemetryMetadata",
	}
}
