// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers an Alexa-enabled device built by an Original Equipment Manufacturer
// (OEM) using Alexa Voice Service (AVS).
func (c *Client) RegisterAVSDevice(ctx context.Context, params *RegisterAVSDeviceInput, optFns ...func(*Options)) (*RegisterAVSDeviceOutput, error) {
	stack := middleware.NewStack("RegisterAVSDevice", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRegisterAVSDeviceMiddlewares(stack)
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
	addOpRegisterAVSDeviceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterAVSDevice(options.Region), middleware.Before)

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
			OperationName: "RegisterAVSDevice",
			Err:           err,
		}
	}
	out := result.(*RegisterAVSDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterAVSDeviceInput struct {
	// The device type ID for your AVS device generated by Amazon when the OEM creates
	// a new product on Amazon's Developer Console.
	AmazonId *string
	// The ARN of the room with which to associate your AVS device.
	RoomArn *string
	// The product ID used to identify your AVS device during authorization.
	ProductId *string
	// The code that is obtained after your AVS device has made a POST request to LWA
	// as a part of the Device Authorization Request component of the OAuth code-based
	// linking specification.
	UserCode *string
	// The client ID of the OEM used for code-based linking authorization on an AVS
	// device.
	ClientId *string
	// The key generated by the OEM that uniquely identifies a specified instance of
	// your AVS device.
	DeviceSerialNumber *string
}

type RegisterAVSDeviceOutput struct {
	// The ARN of the device.
	DeviceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRegisterAVSDeviceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterAVSDevice{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterAVSDevice{}, middleware.After)
}

func newServiceMetadataMiddleware_opRegisterAVSDevice(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "RegisterAVSDevice",
	}
}
