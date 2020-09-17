// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Inserts or deletes RegexPatternString objects in a RegexPatternSet
// (). For each RegexPatternString object, you specify the following values:
//
//     *
// Whether to insert or delete the RegexPatternString.
//
//     * The regular
// expression pattern that you want to insert or delete. For more information, see
// RegexPatternSet ().
//
// For example, you can create a RegexPatternString such as
// B[a@]dB[o0]t. AWS WAF will match this RegexPatternString to:
//
//     * BadBot
//
//
// * BadB0t
//
//     * B@dBot
//
//     * B@dB0t
//
// To create and configure a RegexPatternSet,
// perform the following steps:
//
//     * Create a RegexPatternSet. For more
// information, see CreateRegexPatternSet ().
//
//     * Use GetChangeToken () to get
// the change token that you provide in the ChangeToken parameter of an
// UpdateRegexPatternSet request.
//
//     * Submit an UpdateRegexPatternSet request to
// specify the regular expression pattern that you want AWS WAF to watch for.
//
// For
// more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/).
func (c *Client) UpdateRegexPatternSet(ctx context.Context, params *UpdateRegexPatternSetInput, optFns ...func(*Options)) (*UpdateRegexPatternSetOutput, error) {
	stack := middleware.NewStack("UpdateRegexPatternSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateRegexPatternSetMiddlewares(stack)
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
	addOpUpdateRegexPatternSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRegexPatternSet(options.Region), middleware.Before)
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
			OperationName: "UpdateRegexPatternSet",
			Err:           err,
		}
	}
	out := result.(*UpdateRegexPatternSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRegexPatternSetInput struct {
	// An array of RegexPatternSetUpdate objects that you want to insert into or delete
	// from a RegexPatternSet ().
	Updates []*types.RegexPatternSetUpdate
	// The RegexPatternSetId of the RegexPatternSet () that you want to update.
	// RegexPatternSetId is returned by CreateRegexPatternSet () and by
	// ListRegexPatternSets ().
	RegexPatternSetId *string
	// The value returned by the most recent call to GetChangeToken ().
	ChangeToken *string
}

type UpdateRegexPatternSetOutput struct {
	// The ChangeToken that you used to submit the UpdateRegexPatternSet request. You
	// can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus ().
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateRegexPatternSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateRegexPatternSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateRegexPatternSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateRegexPatternSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "UpdateRegexPatternSet",
	}
}
