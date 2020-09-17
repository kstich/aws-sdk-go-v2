// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all principal ARNs associated with the specified portfolio.
func (c *Client) ListPrincipalsForPortfolio(ctx context.Context, params *ListPrincipalsForPortfolioInput, optFns ...func(*Options)) (*ListPrincipalsForPortfolioOutput, error) {
	stack := middleware.NewStack("ListPrincipalsForPortfolio", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListPrincipalsForPortfolioMiddlewares(stack)
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
	addOpListPrincipalsForPortfolioValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPrincipalsForPortfolio(options.Region), middleware.Before)
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
			OperationName: "ListPrincipalsForPortfolio",
			Err:           err,
		}
	}
	out := result.(*ListPrincipalsForPortfolioOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPrincipalsForPortfolioInput struct {
	// The maximum number of items to return with this call.
	PageSize *int32
	// The language code.
	//
	//     * en - English (default)
	//
	//     * jp - Japanese
	//
	//     * zh
	// - Chinese
	AcceptLanguage *string
	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string
	// The portfolio identifier.
	PortfolioId *string
}

type ListPrincipalsForPortfolioOutput struct {
	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string
	// The IAM principals (users or roles) associated with the portfolio.
	Principals []*types.Principal

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListPrincipalsForPortfolioMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListPrincipalsForPortfolio{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPrincipalsForPortfolio{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPrincipalsForPortfolio(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "ListPrincipalsForPortfolio",
	}
}
