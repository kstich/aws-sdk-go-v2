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

// Adds termination SIP credentials for the specified Amazon Chime Voice Connector.
func (c *Client) PutVoiceConnectorTerminationCredentials(ctx context.Context, params *PutVoiceConnectorTerminationCredentialsInput, optFns ...func(*Options)) (*PutVoiceConnectorTerminationCredentialsOutput, error) {
	stack := middleware.NewStack("PutVoiceConnectorTerminationCredentials", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutVoiceConnectorTerminationCredentialsMiddlewares(stack)
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
	addOpPutVoiceConnectorTerminationCredentialsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutVoiceConnectorTerminationCredentials(options.Region), middleware.Before)
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
			OperationName: "PutVoiceConnectorTerminationCredentials",
			Err:           err,
		}
	}
	out := result.(*PutVoiceConnectorTerminationCredentialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutVoiceConnectorTerminationCredentialsInput struct {
	// The termination SIP credentials.
	Credentials []*types.Credential
	// The Amazon Chime Voice Connector ID.
	VoiceConnectorId *string
}

type PutVoiceConnectorTerminationCredentialsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutVoiceConnectorTerminationCredentialsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutVoiceConnectorTerminationCredentials{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutVoiceConnectorTerminationCredentials{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutVoiceConnectorTerminationCredentials(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "PutVoiceConnectorTerminationCredentials",
	}
}
