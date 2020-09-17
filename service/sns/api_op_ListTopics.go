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

// Returns a list of the requester's topics. Each call returns a limited list of
// topics, up to 100. If there are more topics, a NextToken is also returned. Use
// the NextToken parameter in a new ListTopics call to get further results. This
// action is throttled at 30 transactions per second (TPS).
func (c *Client) ListTopics(ctx context.Context, params *ListTopicsInput, optFns ...func(*Options)) (*ListTopicsOutput, error) {
	stack := middleware.NewStack("ListTopics", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpListTopicsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTopics(options.Region), middleware.Before)
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
			OperationName: "ListTopics",
			Err:           err,
		}
	}
	out := result.(*ListTopicsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTopicsInput struct {
	// Token returned by the previous ListTopics request.
	NextToken *string
}

// Response for ListTopics action.
type ListTopicsOutput struct {
	// A list of topic ARNs.
	Topics []*types.Topic
	// Token to pass along to the next ListTopics request. This element is returned if
	// there are additional topics to retrieve.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpListTopicsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpListTopics{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpListTopics{}, middleware.After)
}

func newServiceMetadataMiddleware_opListTopics(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "ListTopics",
	}
}
