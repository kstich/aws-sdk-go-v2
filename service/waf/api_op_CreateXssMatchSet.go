// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
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
// global use. Creates an XssMatchSet (), which you use to allow, block, or count
// requests that contain cross-site scripting attacks in the specified part of web
// requests. AWS WAF searches for character sequences that are likely to be
// malicious strings. To create and configure an XssMatchSet, perform the following
// steps:
//
//     * Use GetChangeToken () to get the change token that you provide in
// the ChangeToken parameter of a CreateXssMatchSet request.
//
//     * Submit a
// CreateXssMatchSet request.
//
//     * Use GetChangeToken to get the change token
// that you provide in the ChangeToken parameter of an UpdateXssMatchSet ()
// request.
//
//     * Submit an UpdateXssMatchSet () request to specify the parts of
// web requests in which you want to allow, block, or count cross-site scripting
// attacks.
//
// For more information about how to use the AWS WAF API to allow or
// block HTTP requests, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/).
func (c *Client) CreateXssMatchSet(ctx context.Context, params *CreateXssMatchSetInput, optFns ...func(*Options)) (*CreateXssMatchSetOutput, error) {
	stack := middleware.NewStack("CreateXssMatchSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateXssMatchSetMiddlewares(stack)
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
	addOpCreateXssMatchSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateXssMatchSet(options.Region), middleware.Before)
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
			OperationName: "CreateXssMatchSet",
			Err:           err,
		}
	}
	out := result.(*CreateXssMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to create an XssMatchSet ().
type CreateXssMatchSetInput struct {
	// A friendly name or description for the XssMatchSet () that you're creating. You
	// can't change Name after you create the XssMatchSet.
	Name *string
	// The value returned by the most recent call to GetChangeToken ().
	ChangeToken *string
}

// The response to a CreateXssMatchSet request.
type CreateXssMatchSetOutput struct {
	// The ChangeToken that you used to submit the CreateXssMatchSet request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus ().
	ChangeToken *string
	// An XssMatchSet ().
	XssMatchSet *types.XssMatchSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateXssMatchSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateXssMatchSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateXssMatchSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateXssMatchSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "CreateXssMatchSet",
	}
}
