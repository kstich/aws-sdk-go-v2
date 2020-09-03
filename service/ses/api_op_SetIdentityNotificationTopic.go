// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets an Amazon Simple Notification Service (Amazon SNS) topic to use when
// delivering notifications. When you use this operation, you specify a verified
// identity, such as an email address or domain. When you send an email that uses
// the chosen identity in the Source field, Amazon SES sends notifications to the
// topic you specified. You can send bounce, complaint, or delivery notifications
// (or any combination of the three) to the Amazon SNS topic that you specify. You
// can execute this operation no more than once per second. For more information
// about feedback notification, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/notifications.html).
func (c *Client) SetIdentityNotificationTopic(ctx context.Context, params *SetIdentityNotificationTopicInput, optFns ...func(*Options)) (*SetIdentityNotificationTopicOutput, error) {
	stack := middleware.NewStack("SetIdentityNotificationTopic", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpSetIdentityNotificationTopicMiddlewares(stack)
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
	addOpSetIdentityNotificationTopicValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetIdentityNotificationTopic(options.Region), middleware.Before)

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
			OperationName: "SetIdentityNotificationTopic",
			Err:           err,
		}
	}
	out := result.(*SetIdentityNotificationTopicOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to specify the Amazon SNS topic to which Amazon SES will
// publish bounce, complaint, or delivery notifications for emails sent with that
// identity as the Source. For information about Amazon SES notifications, see the
// Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/notifications-via-sns.html).
type SetIdentityNotificationTopicInput struct {
	// The type of notifications that will be published to the specified Amazon SNS
	// topic.
	NotificationType types.NotificationType
	// The Amazon Resource Name (ARN) of the Amazon SNS topic. If the parameter is
	// omitted from the request or a null value is passed, SnsTopic is cleared and
	// publishing is disabled.
	SnsTopic *string
	// The identity (email address or domain) that you want to set the Amazon SNS topic
	// for. You can only specify a verified identity for this parameter. You can
	// specify an identity by using its name or by using its Amazon Resource Name
	// (ARN). The following examples are all valid identities: sender@example.com,
	// example.com, arn:aws:ses:us-east-1:123456789012:identity/example.com.
	Identity *string
}

// An empty element returned on a successful request.
type SetIdentityNotificationTopicOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpSetIdentityNotificationTopicMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpSetIdentityNotificationTopic{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpSetIdentityNotificationTopic{}, middleware.After)
}

func newServiceMetadataMiddleware_opSetIdentityNotificationTopic(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "SetIdentityNotificationTopic",
	}
}
