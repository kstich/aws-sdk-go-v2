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
// Creates a RegexPatternSet (), which you reference in a
// RegexPatternSetReferenceStatement (), to have AWS WAF inspect a web request
// component for the specified patterns.
func (c *Client) CreateRegexPatternSet(ctx context.Context, params *CreateRegexPatternSetInput, optFns ...func(*Options)) (*CreateRegexPatternSetOutput, error) {
	stack := middleware.NewStack("CreateRegexPatternSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateRegexPatternSetMiddlewares(stack)
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
	addOpCreateRegexPatternSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRegexPatternSet(options.Region), middleware.Before)
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
			OperationName: "CreateRegexPatternSet",
			Err:           err,
		}
	}
	out := result.(*CreateRegexPatternSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRegexPatternSetInput struct {
	// The name of the set. You cannot change the name after you create the set.
	Name *string
	// An array of key:value pairs to associate with the resource.
	Tags []*types.Tag
	// Array of regular expression strings.
	RegularExpressionList []*types.Regex
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
	// A description of the set that helps with identification. You cannot change the
	// description of a set after you create it.
	Description *string
}

type CreateRegexPatternSetOutput struct {
	// High-level information about a RegexPatternSet (), returned by operations like
	// create and list. This provides information like the ID, that you can use to
	// retrieve and manage a RegexPatternSet, and the ARN, that you provide to the
	// RegexPatternSetReferenceStatement () to use the pattern set in a Rule ().
	Summary *types.RegexPatternSetSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateRegexPatternSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRegexPatternSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRegexPatternSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateRegexPatternSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "CreateRegexPatternSet",
	}
}
