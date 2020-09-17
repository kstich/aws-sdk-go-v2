// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Permanently ends a session and closes the data connection between the Session
// Manager client and SSM Agent on the instance. A terminated session cannot be
// resumed.
func (c *Client) TerminateSession(ctx context.Context, params *TerminateSessionInput, optFns ...func(*Options)) (*TerminateSessionOutput, error) {
	stack := middleware.NewStack("TerminateSession", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpTerminateSessionMiddlewares(stack)
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
	addOpTerminateSessionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTerminateSession(options.Region), middleware.Before)
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
			OperationName: "TerminateSession",
			Err:           err,
		}
	}
	out := result.(*TerminateSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TerminateSessionInput struct {
	// The ID of the session to terminate.
	SessionId *string
}

type TerminateSessionOutput struct {
	// The ID of the session that has been terminated.
	SessionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpTerminateSessionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpTerminateSession{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpTerminateSession{}, middleware.After)
}

func newServiceMetadataMiddleware_opTerminateSession(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "TerminateSession",
	}
}
