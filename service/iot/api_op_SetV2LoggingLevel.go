// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the logging level.
func (c *Client) SetV2LoggingLevel(ctx context.Context, params *SetV2LoggingLevelInput, optFns ...func(*Options)) (*SetV2LoggingLevelOutput, error) {
	stack := middleware.NewStack("SetV2LoggingLevel", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpSetV2LoggingLevelMiddlewares(stack)
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
	addOpSetV2LoggingLevelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetV2LoggingLevel(options.Region), middleware.Before)
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
			OperationName: "SetV2LoggingLevel",
			Err:           err,
		}
	}
	out := result.(*SetV2LoggingLevelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetV2LoggingLevelInput struct {
	// The log level.
	LogLevel types.LogLevel
	// The log target.
	LogTarget *types.LogTarget
}

type SetV2LoggingLevelOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpSetV2LoggingLevelMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpSetV2LoggingLevel{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpSetV2LoggingLevel{}, middleware.After)
}

func newServiceMetadataMiddleware_opSetV2LoggingLevel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "SetV2LoggingLevel",
	}
}
