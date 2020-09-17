// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the specified fields for the specified stack.
func (c *Client) UpdateStack(ctx context.Context, params *UpdateStackInput, optFns ...func(*Options)) (*UpdateStackOutput, error) {
	stack := middleware.NewStack("UpdateStack", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateStackMiddlewares(stack)
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
	addOpUpdateStackValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateStack(options.Region), middleware.Before)
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
			OperationName: "UpdateStack",
			Err:           err,
		}
	}
	out := result.(*UpdateStackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateStackInput struct {
	// The actions that are enabled or disabled for users during their streaming
	// sessions. By default, these actions are enabled.
	UserSettings []*types.UserSetting
	// The persistent application settings for users of a stack. When these settings
	// are enabled, changes that users make to applications and Windows settings are
	// automatically saved after each session and applied to the next session.
	ApplicationSettings *types.ApplicationSettings
	// The list of interface VPC endpoint (interface endpoint) objects. Users of the
	// stack can connect to AppStream 2.0 only through the specified endpoints.
	AccessEndpoints []*types.AccessEndpoint
	// The domains where AppStream 2.0 streaming sessions can be embedded in an iframe.
	// You must approve the domains that you want to host embedded AppStream 2.0
	// streaming sessions.
	EmbedHostDomains []*string
	// The stack attributes to delete.
	AttributesToDelete []types.StackAttribute
	// The description to display.
	Description *string
	// The stack name to display.
	DisplayName *string
	// The storage connectors to enable.
	StorageConnectors []*types.StorageConnector
	// The URL that users are redirected to after they choose the Send Feedback link.
	// If no URL is specified, no Send Feedback link is displayed.
	FeedbackURL *string
	// Deletes the storage connectors currently enabled for the stack.
	DeleteStorageConnectors *bool
	// The name of the stack.
	Name *string
	// The URL that users are redirected to after their streaming session ends.
	RedirectURL *string
}

type UpdateStackOutput struct {
	// Information about the stack.
	Stack *types.Stack

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateStackMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateStack{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateStack{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateStack(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appstream",
		OperationName: "UpdateStack",
	}
}
