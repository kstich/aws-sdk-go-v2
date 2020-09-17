// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the account IDs that have access to the specified portfolio. A delegated
// admin can list the accounts that have access to the shared portfolio. Note that
// if a delegated admin is de-registered, they can no longer perform this
// operation.
func (c *Client) ListPortfolioAccess(ctx context.Context, params *ListPortfolioAccessInput, optFns ...func(*Options)) (*ListPortfolioAccessOutput, error) {
	stack := middleware.NewStack("ListPortfolioAccess", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListPortfolioAccessMiddlewares(stack)
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
	addOpListPortfolioAccessValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPortfolioAccess(options.Region), middleware.Before)
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
			OperationName: "ListPortfolioAccess",
			Err:           err,
		}
	}
	out := result.(*ListPortfolioAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPortfolioAccessInput struct {
	// The maximum number of items to return with this call.
	PageSize *int32
	// The portfolio identifier.
	PortfolioId *string
	// The language code.
	//
	//     * en - English (default)
	//
	//     * jp - Japanese
	//
	//     * zh
	// - Chinese
	AcceptLanguage *string
	// The ID of an organization node the portfolio is shared with. All children of
	// this node with an inherited portfolio share will be returned.
	OrganizationParentId *string
	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string
}

type ListPortfolioAccessOutput struct {
	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string
	// Information about the AWS accounts with access to the portfolio.
	AccountIds []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListPortfolioAccessMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListPortfolioAccess{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPortfolioAccess{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPortfolioAccess(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "ListPortfolioAccess",
	}
}
