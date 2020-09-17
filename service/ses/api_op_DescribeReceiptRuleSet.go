// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the details of the specified receipt rule set. For information about
// managing receipt rule sets, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-managing-receipt-rule-sets.html).
// You can execute this operation no more than once per second.
func (c *Client) DescribeReceiptRuleSet(ctx context.Context, params *DescribeReceiptRuleSetInput, optFns ...func(*Options)) (*DescribeReceiptRuleSetOutput, error) {
	stack := middleware.NewStack("DescribeReceiptRuleSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeReceiptRuleSetMiddlewares(stack)
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
	addOpDescribeReceiptRuleSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReceiptRuleSet(options.Region), middleware.Before)
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
			OperationName: "DescribeReceiptRuleSet",
			Err:           err,
		}
	}
	out := result.(*DescribeReceiptRuleSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to return the details of a receipt rule set. You use
// receipt rule sets to receive email with Amazon SES. For more information, see
// the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-concepts.html).
type DescribeReceiptRuleSetInput struct {
	// The name of the receipt rule set to describe.
	RuleSetName *string
}

// Represents the details of the specified receipt rule set.
type DescribeReceiptRuleSetOutput struct {
	// The metadata for the receipt rule set, which consists of the rule set name and
	// the timestamp of when the rule set was created.
	Metadata *types.ReceiptRuleSetMetadata
	// A list of the receipt rules that belong to the specified receipt rule set.
	Rules []*types.ReceiptRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeReceiptRuleSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeReceiptRuleSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeReceiptRuleSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeReceiptRuleSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "DescribeReceiptRuleSet",
	}
}
