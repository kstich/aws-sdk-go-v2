// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the platform application objects for the supported push notification
// services, such as APNS and GCM (Firebase Cloud Messaging). The results for
// ListPlatformApplications are paginated and return a limited list of
// applications, up to 100. If additional records are available after the first
// page results, then a NextToken string will be returned. To receive the next
// page, you call ListPlatformApplications using the NextToken string received from
// the previous call. When there are no more records to return, NextToken will be
// null. For more information, see Using Amazon SNS Mobile Push Notifications
// (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html). This action is
// throttled at 15 transactions per second (TPS).
func (c *Client) ListPlatformApplications(ctx context.Context, params *ListPlatformApplicationsInput, optFns ...func(*Options)) (*ListPlatformApplicationsOutput, error) {
	stack := middleware.NewStack("ListPlatformApplications", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpListPlatformApplicationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPlatformApplications(options.Region), middleware.Before)

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
			OperationName: "ListPlatformApplications",
			Err:           err,
		}
	}
	out := result.(*ListPlatformApplicationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for ListPlatformApplications action.
type ListPlatformApplicationsInput struct {
	// NextToken string is used when calling ListPlatformApplications action to
	// retrieve additional records that are available after the first page results.
	NextToken *string
}

// Response for ListPlatformApplications action.
type ListPlatformApplicationsOutput struct {
	// NextToken string is returned when calling ListPlatformApplications action if
	// additional records are available after the first page results.
	NextToken *string
	// Platform applications returned when calling ListPlatformApplications action.
	PlatformApplications []*types.PlatformApplication

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpListPlatformApplicationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpListPlatformApplications{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpListPlatformApplications{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPlatformApplications(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "ListPlatformApplications",
	}
}
