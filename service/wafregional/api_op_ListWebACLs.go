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
// global use. Returns an array of WebACLSummary () objects in the response.
func (c *Client) ListWebACLs(ctx context.Context, params *ListWebACLsInput, optFns ...func(*Options)) (*ListWebACLsOutput, error) {
	stack := middleware.NewStack("ListWebACLs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListWebACLsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListWebACLs(options.Region), middleware.Before)
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
			OperationName: "ListWebACLs",
			Err:           err,
		}
	}
	out := result.(*ListWebACLsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWebACLsInput struct {
	// Specifies the number of WebACL objects that you want AWS WAF to return for this
	// request. If you have more WebACL objects than the number that you specify for
	// Limit, the response includes a NextMarker value that you can use to get another
	// batch of WebACL objects.
	Limit *int32
	// If you specify a value for Limit and you have more WebACL objects than the
	// number that you specify for Limit, AWS WAF returns a NextMarker value in the
	// response that allows you to list another group of WebACL objects. For the second
	// and subsequent ListWebACLs requests, specify the value of NextMarker from the
	// previous response to get information about another batch of WebACL objects.
	NextMarker *string
}

type ListWebACLsOutput struct {
	// An array of WebACLSummary () objects.
	WebACLs []*types.WebACLSummary
	// If you have more WebACL objects than the number that you specified for Limit in
	// the request, the response includes a NextMarker value. To list more WebACL
	// objects, submit another ListWebACLs request, and specify the NextMarker value
	// from the response in the NextMarker value in the next request.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListWebACLsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListWebACLs{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListWebACLs{}, middleware.After)
}

func newServiceMetadataMiddleware_opListWebACLs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "ListWebACLs",
	}
}
