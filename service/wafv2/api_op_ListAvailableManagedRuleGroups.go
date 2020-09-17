// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from the
// prior release, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
// Retrieves an array of managed rule groups that are available for you to use.
// This list includes all AWS Managed Rules rule groups and the AWS Marketplace
// managed rule groups that you're subscribed to.
func (c *Client) ListAvailableManagedRuleGroups(ctx context.Context, params *ListAvailableManagedRuleGroupsInput, optFns ...func(*Options)) (*ListAvailableManagedRuleGroupsOutput, error) {
	stack := middleware.NewStack("ListAvailableManagedRuleGroups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListAvailableManagedRuleGroupsMiddlewares(stack)
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
	addOpListAvailableManagedRuleGroupsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAvailableManagedRuleGroups(options.Region), middleware.Before)
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
			OperationName: "ListAvailableManagedRuleGroups",
			Err:           err,
		}
	}
	out := result.(*ListAvailableManagedRuleGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAvailableManagedRuleGroupsInput struct {
	// Specifies whether this is for an AWS CloudFront distribution or for a regional
	// application. A regional application can be an Application Load Balancer (ALB) or
	// an API Gateway stage. To work with CloudFront, you must also specify the Region
	// US East (N. Virginia) as follows:
	//
	//     * CLI - Specify the Region when you use
	// the CloudFront scope: --scope=CLOUDFRONT --region=us-east-1.
	//
	//     * API and SDKs
	// - For all calls, use the Region endpoint us-east-1.
	Scope types.Scope
	// The maximum number of objects that you want AWS WAF to return for this request.
	// If more objects are available, in the response, AWS WAF provides a NextMarker
	// value that you can use in a subsequent call to get the next batch of objects.
	Limit *int32
	// When you request a list of objects with a Limit setting, if the number of
	// objects that are still available for retrieval exceeds the limit, AWS WAF
	// returns a NextMarker value in the response. To retrieve the next batch of
	// objects, provide the marker from the prior call in your next request.
	NextMarker *string
}

type ListAvailableManagedRuleGroupsOutput struct {
	//
	ManagedRuleGroups []*types.ManagedRuleGroupSummary
	// When you request a list of objects with a Limit setting, if the number of
	// objects that are still available for retrieval exceeds the limit, AWS WAF
	// returns a NextMarker value in the response. To retrieve the next batch of
	// objects, provide the marker from the prior call in your next request.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListAvailableManagedRuleGroupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListAvailableManagedRuleGroups{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAvailableManagedRuleGroups{}, middleware.After)
}

func newServiceMetadataMiddleware_opListAvailableManagedRuleGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "ListAvailableManagedRuleGroups",
	}
}
