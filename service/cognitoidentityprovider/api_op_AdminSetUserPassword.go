// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the specified user's password in a user pool as an administrator. Works on
// any user. The password can be temporary or permanent. If it is temporary, the
// user status will be placed into the FORCE_CHANGE_PASSWORD state. When the user
// next tries to sign in, the InitiateAuth/AdminInitiateAuth response will contain
// the NEW_PASSWORD_REQUIRED challenge. If the user does not sign in before it
// expires, the user will not be able to sign in and their password will need to be
// reset by an administrator. Once the user has set a new password, or the password
// is permanent, the user status will be set to Confirmed.
func (c *Client) AdminSetUserPassword(ctx context.Context, params *AdminSetUserPasswordInput, optFns ...func(*Options)) (*AdminSetUserPasswordOutput, error) {
	stack := middleware.NewStack("AdminSetUserPassword", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAdminSetUserPasswordMiddlewares(stack)
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
	addOpAdminSetUserPasswordValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAdminSetUserPassword(options.Region), middleware.Before)
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
			OperationName: "AdminSetUserPassword",
			Err:           err,
		}
	}
	out := result.(*AdminSetUserPasswordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AdminSetUserPasswordInput struct {
	// The user pool ID for the user pool where you want to set the user's password.
	UserPoolId *string
	// The password for the user.
	Password *string
	// True if the password is permanent, False if it is temporary.
	Permanent *bool
	// The user name of the user whose password you wish to set.
	Username *string
}

type AdminSetUserPasswordOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAdminSetUserPasswordMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAdminSetUserPassword{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminSetUserPassword{}, middleware.After)
}

func newServiceMetadataMiddleware_opAdminSetUserPassword(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminSetUserPassword",
	}
}
