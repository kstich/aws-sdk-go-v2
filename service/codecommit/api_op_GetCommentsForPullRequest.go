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

// Returns comments made on a pull request. Reaction counts might include numbers
// from user identities who were deleted after the reaction was made. For a count
// of reactions from active identities, use GetCommentReactions.
func (c *Client) GetCommentsForPullRequest(ctx context.Context, params *GetCommentsForPullRequestInput, optFns ...func(*Options)) (*GetCommentsForPullRequestOutput, error) {
	stack := middleware.NewStack("GetCommentsForPullRequest", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetCommentsForPullRequestMiddlewares(stack)
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
	addOpGetCommentsForPullRequestValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCommentsForPullRequest(options.Region), middleware.Before)

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
			OperationName: "GetCommentsForPullRequest",
			Err:           err,
		}
	}
	out := result.(*GetCommentsForPullRequestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCommentsForPullRequestInput struct {
	// The full commit ID of the commit in the destination branch that was the tip of
	// the branch at the time the pull request was created.
	BeforeCommitId *string
	// The full commit ID of the commit in the source branch that was the tip of the
	// branch at the time the comment was made.
	AfterCommitId *string
	// An enumeration token that, when provided in a request, returns the next batch of
	// the results.
	NextToken *string
	// A non-zero, non-negative integer used to limit the number of returned results.
	// The default is 100 comments. You can return up to 500 comments with a single
	// request.
	MaxResults *int32
	// The system-generated ID of the pull request. To get this ID, use
	// ListPullRequests ().
	PullRequestId *string
	// The name of the repository that contains the pull request.
	RepositoryName *string
}

type GetCommentsForPullRequestOutput struct {
	// An array of comment objects on the pull request.
	CommentsForPullRequestData []*types.CommentsForPullRequest
	// An enumeration token that can be used in a request to return the next batch of
	// the results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetCommentsForPullRequestMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetCommentsForPullRequest{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCommentsForPullRequest{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetCommentsForPullRequest(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "GetCommentsForPullRequest",
	}
}
