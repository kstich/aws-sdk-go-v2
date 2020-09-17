// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

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
// global use. Permanently deletes a ByteMatchSet (). You can't delete a
// ByteMatchSet if it's still used in any Rules or if it still includes any
// ByteMatchTuple () objects (any filters). If you just want to remove a
// ByteMatchSet from a Rule, use UpdateRule (). To permanently delete a
// ByteMatchSet, perform the following steps:
//
//     * Update the ByteMatchSet to
// remove filters, if any. For more information, see UpdateByteMatchSet ().
//
//     *
// Use GetChangeToken () to get the change token that you provide in the
// ChangeToken parameter of a DeleteByteMatchSet request.
//
//     * Submit a
// DeleteByteMatchSet request.
func (c *Client) DeleteByteMatchSet(ctx context.Context, params *DeleteByteMatchSetInput, optFns ...func(*Options)) (*DeleteByteMatchSetOutput, error) {
	stack := middleware.NewStack("DeleteByteMatchSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteByteMatchSetMiddlewares(stack)
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
	addOpDeleteByteMatchSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteByteMatchSet(options.Region), middleware.Before)
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
			OperationName: "DeleteByteMatchSet",
			Err:           err,
		}
	}
	out := result.(*DeleteByteMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteByteMatchSetInput struct {
	// The ByteMatchSetId of the ByteMatchSet () that you want to delete.
	// ByteMatchSetId is returned by CreateByteMatchSet () and by ListByteMatchSets ().
	ByteMatchSetId *string
	// The value returned by the most recent call to GetChangeToken ().
	ChangeToken *string
}

type DeleteByteMatchSetOutput struct {
	// The ChangeToken that you used to submit the DeleteByteMatchSet request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus ().
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteByteMatchSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteByteMatchSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteByteMatchSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteByteMatchSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "DeleteByteMatchSet",
	}
}
