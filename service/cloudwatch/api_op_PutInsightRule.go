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

// Creates a Contributor Insights rule. Rules evaluate log events in a CloudWatch
// Logs log group, enabling you to find contributor data for the log events in that
// log group. For more information, see Using Contributor Insights to Analyze
// High-Cardinality Data
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/ContributorInsights.html).
// If you create a rule, delete it, and then re-create it with the same name,
// historical data from the first time the rule was created might not be available.
func (c *Client) PutInsightRule(ctx context.Context, params *PutInsightRuleInput, optFns ...func(*Options)) (*PutInsightRuleOutput, error) {
	stack := middleware.NewStack("PutInsightRule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpPutInsightRuleMiddlewares(stack)
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
	addOpPutInsightRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutInsightRule(options.Region), middleware.Before)
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
			OperationName: "PutInsightRule",
			Err:           err,
		}
	}
	out := result.(*PutInsightRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutInsightRuleInput struct {
	// The state of the rule. Valid values are ENABLED and DISABLED.
	RuleState *string
	// A unique name for the rule.
	RuleName *string
	// A list of key-value pairs to associate with the Contributor Insights rule. You
	// can associate as many as 50 tags with a rule. Tags can help you organize and
	// categorize your resources. You can also use them to scope user permissions, by
	// granting a user permission to access or change only the resources that have
	// certain tag values. To be able to associate tags with a rule, you must have the
	// cloudwatch:TagResource permission in addition to the cloudwatch:PutInsightRule
	// permission. If you are using this operation to update an existing Contributor
	// Insights rule, any tags you specify in this parameter are ignored. To change the
	// tags of an existing rule, use TagResource
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_TagResource.html).
	Tags []*types.Tag
	// The definition of the rule, as a JSON object. For details on the valid syntax,
	// see Contributor Insights Rule Syntax
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/ContributorInsights-RuleSyntax.html).
	RuleDefinition *string
}

type PutInsightRuleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpPutInsightRuleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpPutInsightRule{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpPutInsightRule{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutInsightRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "PutInsightRule",
	}
}
