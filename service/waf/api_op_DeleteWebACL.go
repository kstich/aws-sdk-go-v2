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
// global use. Permanently deletes a WebACL (). You can't delete a WebACL if it
// still contains any Rules. To delete a WebACL, perform the following steps:
//
//
// * Update the WebACL to remove Rules, if any. For more information, see
// UpdateWebACL ().
//
//     * Use GetChangeToken () to get the change token that you
// provide in the ChangeToken parameter of a DeleteWebACL request.
//
//     * Submit a
// DeleteWebACL request.
func (c *Client) DeleteWebACL(ctx context.Context, params *DeleteWebACLInput, optFns ...func(*Options)) (*DeleteWebACLOutput, error) {
	stack := middleware.NewStack("DeleteWebACL", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteWebACLMiddlewares(stack)
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
	addOpDeleteWebACLValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteWebACL(options.Region), middleware.Before)
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
			OperationName: "DeleteWebACL",
			Err:           err,
		}
	}
	out := result.(*DeleteWebACLOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteWebACLInput struct {
	// The value returned by the most recent call to GetChangeToken ().
	ChangeToken *string
	// The WebACLId of the WebACL () that you want to delete. WebACLId is returned by
	// CreateWebACL () and by ListWebACLs ().
	WebACLId *string
}

type DeleteWebACLOutput struct {
	// The ChangeToken that you used to submit the DeleteWebACL request. You can also
	// use this value to query the status of the request. For more information, see
	// GetChangeTokenStatus ().
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteWebACLMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteWebACL{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteWebACL{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteWebACL(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "DeleteWebACL",
	}
}
