// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Posts a comment on the comparison between two commits.
func (c *Client) PostCommentForComparedCommit(ctx context.Context, params *PostCommentForComparedCommitInput, optFns ...func(*Options)) (*PostCommentForComparedCommitOutput, error) {
	stack := middleware.NewStack("PostCommentForComparedCommit", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPostCommentForComparedCommitMiddlewares(stack)
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
	addIdempotencyToken_opPostCommentForComparedCommitMiddleware(stack, options)
	addOpPostCommentForComparedCommitValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPostCommentForComparedCommit(options.Region), middleware.Before)

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
			OperationName: "PostCommentForComparedCommit",
			Err:           err,
		}
	}
	out := result.(*PostCommentForComparedCommitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PostCommentForComparedCommitInput struct {
	// The content of the comment you want to make.
	Content *string
	// A unique, client-generated idempotency token that, when provided in a request,
	// ensures the request cannot be repeated with a changed parameter. If a request is
	// received with the same parameters and a token is included, the request returns
	// information about the initial request that used that token.
	ClientRequestToken *string
	// The name of the repository where you want to post a comment on the comparison
	// between commits.
	RepositoryName *string
	// To establish the directionality of the comparison, the full commit ID of the
	// before commit. Required for commenting on any commit unless that commit is the
	// initial commit.
	BeforeCommitId *string
	// The location of the comparison where you want to comment.
	Location *types.Location
	// To establish the directionality of the comparison, the full commit ID of the
	// after commit.
	AfterCommitId *string
}

type PostCommentForComparedCommitOutput struct {
	// The location of the comment in the comparison between the two commits.
	Location *types.Location
	// In the directionality you established, the blob ID of the after blob.
	AfterBlobId *string
	// In the directionality you established, the full commit ID of the before commit.
	BeforeCommitId *string
	// The name of the repository where you posted a comment on the comparison between
	// commits.
	RepositoryName *string
	// The content of the comment you posted.
	Comment *types.Comment
	// In the directionality you established, the full commit ID of the after commit.
	AfterCommitId *string
	// In the directionality you established, the blob ID of the before blob.
	BeforeBlobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPostCommentForComparedCommitMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPostCommentForComparedCommit{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPostCommentForComparedCommit{}, middleware.After)
}

type idempotencyToken_initializeOpPostCommentForComparedCommit struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpPostCommentForComparedCommit) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpPostCommentForComparedCommit) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*PostCommentForComparedCommitInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *PostCommentForComparedCommitInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opPostCommentForComparedCommitMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpPostCommentForComparedCommit{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opPostCommentForComparedCommit(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "PostCommentForComparedCommit",
	}
}
