// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a user account for the specified Amazon Connect instance.
func (c *Client) CreateUser(ctx context.Context, params *CreateUserInput, optFns ...func(*Options)) (*CreateUserOutput, error) {
	stack := middleware.NewStack("CreateUser", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateUserMiddlewares(stack)
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
	addOpCreateUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUser(options.Region), middleware.Before)

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
			OperationName: "CreateUser",
			Err:           err,
		}
	}
	out := result.(*CreateUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateUserInput struct {
	// One or more tags.
	Tags map[string]*string
	// The identifier of the security profile for the user.
	SecurityProfileIds []*string
	// The identifier of the Amazon Connect instance.
	InstanceId *string
	// The phone settings for the user.
	PhoneConfig *types.UserPhoneConfig
	// The identifier of the hierarchy group for the user.
	HierarchyGroupId *string
	// The identifier of the user account in the directory used for identity
	// management. If Amazon Connect cannot access the directory, you can specify this
	// identifier to authenticate users. If you include the identifier, we assume that
	// Amazon Connect cannot access the directory. Otherwise, the identity information
	// is used to authenticate users from your directory. This parameter is required if
	// you are using an existing directory for identity management in Amazon Connect
	// when Amazon Connect cannot access your directory to authenticate users. If you
	// are using SAML for identity management and include this parameter, an error is
	// returned.
	DirectoryUserId *string
	// The information about the identity of the user.
	IdentityInfo *types.UserIdentityInfo
	// The password for the user account. A password is required if you are using
	// Amazon Connect for identity management. Otherwise, it is an error to include a
	// password.
	Password *string
	// The identifier of the routing profile for the user.
	RoutingProfileId *string
	// The user name for the account. For instances not using SAML for identity
	// management, the user name can include up to 20 characters. If you are using SAML
	// for identity management, the user name can include up to 64 characters from
	// [a-zA-Z0-9_-.\@]+.
	Username *string
}

type CreateUserOutput struct {
	// The identifier of the user account.
	UserId *string
	// The Amazon Resource Name (ARN) of the user account.
	UserArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateUserMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateUser{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateUser{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "CreateUser",
	}
}
