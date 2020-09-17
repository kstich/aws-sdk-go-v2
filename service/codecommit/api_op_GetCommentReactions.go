// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about reactions to a specified comment ID. Reactions from
// users who have been deleted will not be included in the count.
func (c *Client) GetCommentReactions(ctx context.Context, params *GetCommentReactionsInput, optFns ...func(*Options)) (*GetCommentReactionsOutput, error) {
	stack := middleware.NewStack("GetCommentReactions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetCommentReactionsMiddlewares(stack)
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
	addOpGetCommentReactionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCommentReactions(options.Region), middleware.Before)
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
			OperationName: "GetCommentReactions",
			Err:           err,
		}
	}
	out := result.(*GetCommentReactionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCommentReactionsInput struct {
	// The ID of the comment for which you want to get reactions information.
	CommentId *string
	// Optional. The Amazon Resource Name (ARN) of the user or identity for which you
	// want to get reaction information.
	ReactionUserArn *string
	// A non-zero, non-negative integer used to limit the number of returned results.
	// The default is the same as the allowed maximum, 1,000.
	MaxResults *int32
	// An enumeration token that, when provided in a request, returns the next batch of
	// the results.
	NextToken *string
}

type GetCommentReactionsOutput struct {
	// An array of reactions to the specified comment.
	ReactionsForComment []*types.ReactionForComment
	// An enumeration token that can be used in a request to return the next batch of
	// the results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetCommentReactionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetCommentReactions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCommentReactions{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetCommentReactions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "GetCommentReactions",
	}
}
