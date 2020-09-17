// Code generated by smithy-go-codegen DO NOT EDIT.

package polly

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/polly/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of pronunciation lexicons stored in an AWS Region. For more
// information, see Managing Lexicons
// (https://docs.aws.amazon.com/polly/latest/dg/managing-lexicons.html).
func (c *Client) ListLexicons(ctx context.Context, params *ListLexiconsInput, optFns ...func(*Options)) (*ListLexiconsOutput, error) {
	stack := middleware.NewStack("ListLexicons", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListLexiconsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListLexicons(options.Region), middleware.Before)
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
			OperationName: "ListLexicons",
			Err:           err,
		}
	}
	out := result.(*ListLexiconsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLexiconsInput struct {
	// An opaque pagination token returned from previous ListLexicons operation. If
	// present, indicates where to continue the list of lexicons.
	NextToken *string
}

type ListLexiconsOutput struct {
	// A list of lexicon names and attributes.
	Lexicons []*types.LexiconDescription
	// The pagination token to use in the next request to continue the listing of
	// lexicons. NextToken is returned only if the response is truncated.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListLexiconsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListLexicons{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListLexicons{}, middleware.After)
}

func newServiceMetadataMiddleware_opListLexicons(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "polly",
		OperationName: "ListLexicons",
	}
}
