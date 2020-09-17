// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables the specified Contributor Insights rules. When rules are enabled, they
// immediately begin analyzing log data.
func (c *Client) EnableInsightRules(ctx context.Context, params *EnableInsightRulesInput, optFns ...func(*Options)) (*EnableInsightRulesOutput, error) {
	stack := middleware.NewStack("EnableInsightRules", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpEnableInsightRulesMiddlewares(stack)
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
	addOpEnableInsightRulesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableInsightRules(options.Region), middleware.Before)
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
			OperationName: "EnableInsightRules",
			Err:           err,
		}
	}
	out := result.(*EnableInsightRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableInsightRulesInput struct {
	// An array of the rule names to enable. If you need to find out the names of your
	// rules, use DescribeInsightRules
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DescribeInsightRules.html).
	RuleNames []*string
}

type EnableInsightRulesOutput struct {
	// An array listing the rules that could not be enabled. You cannot disable or
	// enable built-in rules.
	Failures []*types.PartialFailure

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpEnableInsightRulesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpEnableInsightRules{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpEnableInsightRules{}, middleware.After)
}

func newServiceMetadataMiddleware_opEnableInsightRules(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "EnableInsightRules",
	}
}
