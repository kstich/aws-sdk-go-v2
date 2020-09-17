// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
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
// global use. Permanently deletes an XssMatchSet (). You can't delete an
// XssMatchSet if it's still used in any Rules or if it still contains any
// XssMatchTuple () objects. If you just want to remove an XssMatchSet from a Rule,
// use UpdateRule (). To permanently delete an XssMatchSet from AWS WAF, perform
// the following steps:
//
//     * Update the XssMatchSet to remove filters, if any.
// For more information, see UpdateXssMatchSet ().
//
//     * Use GetChangeToken () to
// get the change token that you provide in the ChangeToken parameter of a
// DeleteXssMatchSet request.
//
//     * Submit a DeleteXssMatchSet request.
func (c *Client) DeleteXssMatchSet(ctx context.Context, params *DeleteXssMatchSetInput, optFns ...func(*Options)) (*DeleteXssMatchSetOutput, error) {
	stack := middleware.NewStack("DeleteXssMatchSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteXssMatchSetMiddlewares(stack)
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
	addOpDeleteXssMatchSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteXssMatchSet(options.Region), middleware.Before)
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
			OperationName: "DeleteXssMatchSet",
			Err:           err,
		}
	}
	out := result.(*DeleteXssMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to delete an XssMatchSet () from AWS WAF.
type DeleteXssMatchSetInput struct {
	// The value returned by the most recent call to GetChangeToken ().
	ChangeToken *string
	// The XssMatchSetId of the XssMatchSet () that you want to delete. XssMatchSetId
	// is returned by CreateXssMatchSet () and by ListXssMatchSets ().
	XssMatchSetId *string
}

// The response to a request to delete an XssMatchSet () from AWS WAF.
type DeleteXssMatchSetOutput struct {
	// The ChangeToken that you used to submit the DeleteXssMatchSet request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus ().
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteXssMatchSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteXssMatchSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteXssMatchSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteXssMatchSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "DeleteXssMatchSet",
	}
}
