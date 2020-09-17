// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a temporary URL to start an AppStream 2.0 streaming session for the
// specified user. A streaming URL enables application streaming to be tested
// without user setup.
func (c *Client) CreateStreamingURL(ctx context.Context, params *CreateStreamingURLInput, optFns ...func(*Options)) (*CreateStreamingURLOutput, error) {
	stack := middleware.NewStack("CreateStreamingURL", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateStreamingURLMiddlewares(stack)
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
	addOpCreateStreamingURLValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStreamingURL(options.Region), middleware.Before)
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
			OperationName: "CreateStreamingURL",
			Err:           err,
		}
	}
	out := result.(*CreateStreamingURLOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateStreamingURLInput struct {
	// The name of the fleet.
	FleetName *string
	// The time that the streaming URL will be valid, in seconds. Specify a value
	// between 1 and 604800 seconds. The default is 60 seconds.
	Validity *int64
	// The session context. For more information, see Session Context
	// (https://docs.aws.amazon.com/appstream2/latest/developerguide/managing-stacks-fleets.html#managing-stacks-fleets-parameters)
	// in the Amazon AppStream 2.0 Administration Guide.
	SessionContext *string
	// The identifier of the user.
	UserId *string
	// The name of the application to launch after the session starts. This is the name
	// that you specified as Name in the Image Assistant.
	ApplicationId *string
	// The name of the stack.
	StackName *string
}

type CreateStreamingURLOutput struct {
	// The URL to start the AppStream 2.0 streaming session.
	StreamingURL *string
	// The elapsed time, in seconds after the Unix epoch, when this URL expires.
	Expires *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateStreamingURLMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateStreamingURL{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateStreamingURL{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateStreamingURL(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appstream",
		OperationName: "CreateStreamingURL",
	}
}
