// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the user attributes and metadata for a user.
func (c *Client) GetUser(ctx context.Context, params *GetUserInput, optFns ...func(*Options)) (*GetUserOutput, error) {
	stack := middleware.NewStack("GetUser", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetUserMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpGetUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetUser(options.Region), middleware.Before)
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
			OperationName: "GetUser",
			Err:           err,
		}
	}
	out := result.(*GetUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to get information about the user.
type GetUserInput struct {
	// The access token returned by the server response to get information about the
	// user.
	AccessToken *string
}

// Represents the response from the server from the request to get information
// about the user.
type GetUserOutput struct {
	// The user name of the user you wish to retrieve from the get user request.
	Username *string
	// The MFA options that are enabled for the user. The possible values in this list
	// are SMS_MFA and SOFTWARE_TOKEN_MFA.
	UserMFASettingList []*string
	// An array of name-value pairs representing user attributes. For custom
	// attributes, you must prepend the custom: prefix to the attribute name.
	UserAttributes []*types.AttributeType
	// This response parameter is no longer supported. It provides information only
	// about SMS MFA configurations. It doesn't provide information about TOTP software
	// token MFA configurations. To look up information about either type of MFA
	// configuration, use the use the GetUserResponse$UserMFASettingList () response
	// instead.
	MFAOptions []*types.MFAOptionType
	// The user's preferred MFA setting.
	PreferredMfaSetting *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetUserMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetUser{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetUser{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "GetUser",
	}
}
