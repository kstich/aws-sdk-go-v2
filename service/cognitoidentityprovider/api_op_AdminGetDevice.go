// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the device, as an administrator. Calling this action requires developer
// credentials.
func (c *Client) AdminGetDevice(ctx context.Context, params *AdminGetDeviceInput, optFns ...func(*Options)) (*AdminGetDeviceOutput, error) {
	stack := middleware.NewStack("AdminGetDevice", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAdminGetDeviceMiddlewares(stack)
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
	addOpAdminGetDeviceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAdminGetDevice(options.Region), middleware.Before)
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
			OperationName: "AdminGetDevice",
			Err:           err,
		}
	}
	out := result.(*AdminGetDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to get the device, as an administrator.
type AdminGetDeviceInput struct {
	// The user name.
	Username *string
	// The device key.
	DeviceKey *string
	// The user pool ID.
	UserPoolId *string
}

// Gets the device response, as an administrator.
type AdminGetDeviceOutput struct {
	// The device.
	Device *types.DeviceType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAdminGetDeviceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAdminGetDevice{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminGetDevice{}, middleware.After)
}

func newServiceMetadataMiddleware_opAdminGetDevice(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminGetDevice",
	}
}
