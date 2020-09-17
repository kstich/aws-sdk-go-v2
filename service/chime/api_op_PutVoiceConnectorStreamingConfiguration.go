// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds a streaming configuration for the specified Amazon Chime Voice Connector.
// The streaming configuration specifies whether media streaming is enabled for
// sending to Amazon Kinesis. It also sets the retention period, in hours, for the
// Amazon Kinesis data.
func (c *Client) PutVoiceConnectorStreamingConfiguration(ctx context.Context, params *PutVoiceConnectorStreamingConfigurationInput, optFns ...func(*Options)) (*PutVoiceConnectorStreamingConfigurationOutput, error) {
	stack := middleware.NewStack("PutVoiceConnectorStreamingConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutVoiceConnectorStreamingConfigurationMiddlewares(stack)
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
	addOpPutVoiceConnectorStreamingConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutVoiceConnectorStreamingConfiguration(options.Region), middleware.Before)
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
			OperationName: "PutVoiceConnectorStreamingConfiguration",
			Err:           err,
		}
	}
	out := result.(*PutVoiceConnectorStreamingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutVoiceConnectorStreamingConfigurationInput struct {
	// The Amazon Chime Voice Connector ID.
	VoiceConnectorId *string
	// The streaming configuration details to add.
	StreamingConfiguration *types.StreamingConfiguration
}

type PutVoiceConnectorStreamingConfigurationOutput struct {
	// The updated streaming configuration details.
	StreamingConfiguration *types.StreamingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutVoiceConnectorStreamingConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutVoiceConnectorStreamingConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutVoiceConnectorStreamingConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutVoiceConnectorStreamingConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "PutVoiceConnectorStreamingConfiguration",
	}
}
