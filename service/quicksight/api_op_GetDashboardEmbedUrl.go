// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Generates a URL and authorization code that you can embed in your web server
// code. Before you use this command, make sure that you have configured the
// dashboards and permissions. Currently, you can use GetDashboardEmbedURL only
// from the server, not from the user's browser. The following rules apply to the
// combination of URL and authorization code:
//
//     * They must be used together.
//
//
// * They can be used one time only.
//
//     * They are valid for 5 minutes after you
// run this command.
//
//     * The resulting user session is valid for 10 hours.
//
// For
// more information, see Embedding Amazon QuickSight Dashboards
// (https://docs.aws.amazon.com/quicksight/latest/user/embedding-dashboards.html)
// in the Amazon QuickSight User Guide or Embedding Amazon QuickSight Dashboards
// (https://docs.aws.amazon.com/quicksight/latest/APIReference/qs-dev-embedded-dashboards.html)
// in the Amazon QuickSight API Reference.
func (c *Client) GetDashboardEmbedUrl(ctx context.Context, params *GetDashboardEmbedUrlInput, optFns ...func(*Options)) (*GetDashboardEmbedUrlOutput, error) {
	stack := middleware.NewStack("GetDashboardEmbedUrl", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetDashboardEmbedUrlMiddlewares(stack)
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
	addOpGetDashboardEmbedUrlValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDashboardEmbedUrl(options.Region), middleware.Before)
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
			OperationName: "GetDashboardEmbedUrl",
			Err:           err,
		}
	}
	out := result.(*GetDashboardEmbedUrlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDashboardEmbedUrlInput struct {
	// Remove the reset button on the embedded dashboard. The default is FALSE, which
	// enables the reset button.
	ResetDisabled *bool
	// How many minutes the session is valid. The session lifetime must be 15-600
	// minutes.
	SessionLifetimeInMinutes *int64
	// The ID for the AWS account that contains the dashboard that you're embedding.
	AwsAccountId *string
	// The authentication method that the user uses to sign in.
	IdentityType types.IdentityType
	// Remove the undo/redo button on the embedded dashboard. The default is FALSE,
	// which enables the undo/redo button.
	UndoRedoDisabled *bool
	// The Amazon QuickSight user's Amazon Resource Name (ARN), for use with QUICKSIGHT
	// identity type. You can use this for any Amazon QuickSight users in your account
	// (readers, authors, or admins) authenticated as one of the following:
	//
	//     *
	// Active Directory (AD) users or group members
	//
	//     * Invited nonfederated users
	//
	//
	// * IAM users and IAM role-based sessions authenticated through Federated Single
	// Sign-On using SAML, OpenID Connect, or IAM federation.
	UserArn *string
	// The ID for the dashboard, also added to the IAM policy.
	DashboardId *string
}

type GetDashboardEmbedUrlOutput struct {
	// A single-use URL that you can put into your server-side webpage to embed your
	// dashboard. This URL is valid for 5 minutes. The API provides the URL with an
	// auth_code value that enables one (and only one) sign-on to a user session that
	// is valid for 10 hours.
	EmbedUrl *string
	// The AWS request ID for this operation.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetDashboardEmbedUrlMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetDashboardEmbedUrl{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDashboardEmbedUrl{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDashboardEmbedUrl(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "GetDashboardEmbedUrl",
	}
}
