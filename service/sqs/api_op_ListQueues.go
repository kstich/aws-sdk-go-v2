// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of your queues. The maximum number of queues that can be returned
// is 1,000. If you specify a value for the optional QueueNamePrefix parameter,
// only queues with a name that begins with the specified value are returned.
// Cross-account permissions don't apply to this action. For more information, see
// Grant Cross-Account Permissions to a Role and a User Name
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name)
// in the Amazon Simple Queue Service Developer Guide.
func (c *Client) ListQueues(ctx context.Context, params *ListQueuesInput, optFns ...func(*Options)) (*ListQueuesOutput, error) {
	stack := middleware.NewStack("ListQueues", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpListQueuesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListQueues(options.Region), middleware.Before)
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
			OperationName: "ListQueues",
			Err:           err,
		}
	}
	out := result.(*ListQueuesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ListQueuesInput struct {
	// Pagination token to request the next set of results.
	NextToken *string
	// Maximum number of results to include in the response.
	MaxResults *int32
	// A string to use for filtering the list results. Only those queues whose name
	// begins with the specified string are returned. Queue URLs and names are
	// case-sensitive.
	QueueNamePrefix *string
}

// A list of your queues.
type ListQueuesOutput struct {
	// Pagination token to include in the next request.
	NextToken *string
	// A list of queue URLs, up to 1,000 entries, or the value of MaxResults that you
	// sent in the request.
	QueueUrls []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpListQueuesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpListQueues{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpListQueues{}, middleware.After)
}

func newServiceMetadataMiddleware_opListQueues(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sqs",
		OperationName: "ListQueues",
	}
}
