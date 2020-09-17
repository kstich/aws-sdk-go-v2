// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the name, ARN, rules, definition, and effective dates of a Cost Category
// that's defined in the account. You have the option to use EffectiveOn to return
// a Cost Category that is active on a specific date. If there is no EffectiveOn
// specified, you’ll see a Cost Category that is effective on the current date. If
// Cost Category is still effective, EffectiveEnd is omitted in the response.
func (c *Client) DescribeCostCategoryDefinition(ctx context.Context, params *DescribeCostCategoryDefinitionInput, optFns ...func(*Options)) (*DescribeCostCategoryDefinitionOutput, error) {
	stack := middleware.NewStack("DescribeCostCategoryDefinition", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeCostCategoryDefinitionMiddlewares(stack)
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
	addOpDescribeCostCategoryDefinitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCostCategoryDefinition(options.Region), middleware.Before)
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
			OperationName: "DescribeCostCategoryDefinition",
			Err:           err,
		}
	}
	out := result.(*DescribeCostCategoryDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCostCategoryDefinitionInput struct {
	// The unique identifier for your Cost Category.
	CostCategoryArn *string
	// The date when the Cost Category was effective.
	EffectiveOn *string
}

type DescribeCostCategoryDefinitionOutput struct {
	// The structure of Cost Categories. This includes detailed metadata and the set of
	// rules for the CostCategory object.
	CostCategory *types.CostCategory

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeCostCategoryDefinitionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCostCategoryDefinition{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCostCategoryDefinition{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeCostCategoryDefinition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "DescribeCostCategoryDefinition",
	}
}
