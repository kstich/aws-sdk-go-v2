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

// Gets information about the specified portfolio. A delegated admin is authorized
// to invoke this command.
func (c *Client) DescribePortfolio(ctx context.Context, params *DescribePortfolioInput, optFns ...func(*Options)) (*DescribePortfolioOutput, error) {
	stack := middleware.NewStack("DescribePortfolio", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribePortfolioMiddlewares(stack)
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
	addOpDescribePortfolioValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePortfolio(options.Region), middleware.Before)
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
			OperationName: "DescribePortfolio",
			Err:           err,
		}
	}
	out := result.(*DescribePortfolioOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePortfolioInput struct {
	// The language code.
	//
	//     * en - English (default)
	//
	//     * jp - Japanese
	//
	//     * zh
	// - Chinese
	AcceptLanguage *string
	// The portfolio identifier.
	Id *string
}

type DescribePortfolioOutput struct {
	// Information about the associated budgets.
	Budgets []*types.BudgetDetail
	// Information about the TagOptions associated with the portfolio.
	TagOptions []*types.TagOptionDetail
	// Information about the tags associated with the portfolio.
	Tags []*types.Tag
	// Information about the portfolio.
	PortfolioDetail *types.PortfolioDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribePortfolioMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePortfolio{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePortfolio{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribePortfolio(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "DescribePortfolio",
	}
}
