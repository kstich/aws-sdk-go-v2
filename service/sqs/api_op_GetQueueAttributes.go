// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets attributes for the specified queue. To determine whether a queue is FIFO
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html),
// you can check whether QueueName ends with the .fifo suffix.
func (c *Client) GetQueueAttributes(ctx context.Context, params *GetQueueAttributesInput, optFns ...func(*Options)) (*GetQueueAttributesOutput, error) {
	stack := middleware.NewStack("GetQueueAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpGetQueueAttributesMiddlewares(stack)
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
	addOpGetQueueAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetQueueAttributes(options.Region), middleware.Before)
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
			OperationName: "GetQueueAttributes",
			Err:           err,
		}
	}
	out := result.(*GetQueueAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type GetQueueAttributesInput struct {
	// The URL of the Amazon SQS queue whose attribute information is retrieved. Queue
	// URLs and names are case-sensitive.
	QueueUrl *string
	// A list of attributes for which to retrieve information. In the future, new
	// attributes might be added. If you write code that calls this action, we
	// recommend that you structure your code so that it can handle new attributes
	// gracefully. The following attributes are supported: The
	// ApproximateNumberOfMessagesDelayed, ApproximateNumberOfMessagesNotVisible, and
	// ApproximateNumberOfMessagesVisible metrics may not achieve consistency until at
	// least 1 minute after the producers stop sending messages. This period is
	// required for the queue metadata to reach eventual consistency.
	//
	//     * All –
	// Returns all values.
	//
	//     * ApproximateNumberOfMessages – Returns the approximate
	// number of messages available for retrieval from the queue.
	//
	//     *
	// ApproximateNumberOfMessagesDelayed – Returns the approximate number of messages
	// in the queue that are delayed and not available for reading immediately. This
	// can happen when the queue is configured as a delay queue or when a message has
	// been sent with a delay parameter.
	//
	//     * ApproximateNumberOfMessagesNotVisible –
	// Returns the approximate number of messages that are in flight. Messages are
	// considered to be in flight if they have been sent to a client but have not yet
	// been deleted or have not yet reached the end of their visibility window.
	//
	//     *
	// CreatedTimestamp – Returns the time when the queue was created in seconds (epoch
	// time (http://en.wikipedia.org/wiki/Unix_time)).
	//
	//     * DelaySeconds – Returns
	// the default delay on the queue in seconds.
	//
	//     * LastModifiedTimestamp –
	// Returns the time when the queue was last changed in seconds (epoch time
	// (http://en.wikipedia.org/wiki/Unix_time)).
	//
	//     * MaximumMessageSize – Returns
	// the limit of how many bytes a message can contain before Amazon SQS rejects
	// it.
	//
	//     * MessageRetentionPeriod – Returns the length of time, in seconds, for
	// which Amazon SQS retains a message.
	//
	//     * Policy – Returns the policy of the
	// queue.
	//
	//     * QueueArn – Returns the Amazon resource name (ARN) of the queue.
	//
	//
	// * ReceiveMessageWaitTimeSeconds – Returns the length of time, in seconds, for
	// which the ReceiveMessage action waits for a message to arrive.
	//
	//     *
	// RedrivePolicy – The string that includes the parameters for the dead-letter
	// queue functionality of the source queue as a JSON object. For more information
	// about the redrive policy and dead-letter queues, see Using Amazon SQS
	// Dead-Letter Queues
	// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html)
	// in the Amazon Simple Queue Service Developer Guide.
	//
	//         *
	// deadLetterTargetArn – The Amazon Resource Name (ARN) of the dead-letter queue to
	// which Amazon SQS moves messages after the value of maxReceiveCount is
	// exceeded.
	//
	//         * maxReceiveCount – The number of times a message is
	// delivered to the source queue before being moved to the dead-letter queue. When
	// the ReceiveCount for a message exceeds the maxReceiveCount for a queue, Amazon
	// SQS moves the message to the dead-letter-queue.
	//
	//     * VisibilityTimeout –
	// Returns the visibility timeout for the queue. For more information about the
	// visibility timeout, see Visibility Timeout
	// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html)
	// in the Amazon Simple Queue Service Developer Guide.
	//
	//     <p>The following
	// attributes apply only to <a
	// href="https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html">server-side-encryption</a>:</p>
	// <ul> <li> <p> <code>KmsMasterKeyId</code> – Returns the ID of an AWS-managed
	// customer master key (CMK) for Amazon SQS or a custom CMK. For more information,
	// see <a
	// href="https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-sse-key-terms">Key
	// Terms</a>. </p> </li> <li> <p> <code>KmsDataKeyReusePeriodSeconds</code> –
	// Returns the length of time, in seconds, for which Amazon SQS can reuse a data
	// key to encrypt or decrypt messages before calling AWS KMS again. For more
	// information, see <a
	// href="https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-how-does-the-data-key-reuse-period-work">How
	// Does the Data Key Reuse Period Work?</a>. </p> </li> </ul> <p>The following
	// attributes apply only to <a
	// href="https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html">FIFO
	// (first-in-first-out) queues</a>:</p> <ul> <li> <p> <code>FifoQueue</code> –
	// Returns whether the queue is FIFO. For more information, see <a
	// href="https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-understanding-logic">FIFO
	// Queue Logic</a> in the <i>Amazon Simple Queue Service Developer Guide</i>.</p>
	// <note> <p>To determine whether a queue is <a
	// href="https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html">FIFO</a>,
	// you can check whether <code>QueueName</code> ends with the <code>.fifo</code>
	// suffix.</p> </note> </li> <li> <p> <code>ContentBasedDeduplication</code> –
	// Returns whether content-based deduplication is enabled for the queue. For more
	// information, see <a
	// href="https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-exactly-once-processing">Exactly-Once
	// Processing</a> in the <i>Amazon Simple Queue Service Developer Guide</i>. </p>
	// </li> </ul>
	AttributeNames []types.QueueAttributeName
}

// A list of returned queue attributes.
type GetQueueAttributesOutput struct {
	// A map of attributes to their respective values.
	Attributes map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpGetQueueAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpGetQueueAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpGetQueueAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetQueueAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sqs",
		OperationName: "GetQueueAttributes",
	}
}
